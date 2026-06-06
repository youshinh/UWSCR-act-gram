package knowledge

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type KnowledgeChunk struct {
	Source  string
	Content string
}

type RAGManager struct {
	baseDir string
	chunks  []KnowledgeChunk
}

func NewRAGManager(baseDir string) *RAGManager {
	return &RAGManager{
		baseDir: baseDir,
		chunks:  make([]KnowledgeChunk, 0),
	}
}

// LoadKnowledgeFiles は knowledge フォルダ内の txt/md/csv ファイルをインメモリにロードします
func (r *RAGManager) LoadKnowledgeFiles() error {
	if err := os.MkdirAll(r.baseDir, 0755); err != nil {
		return err
	}

	r.chunks = nil
	err := filepath.Walk(r.baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && (strings.HasSuffix(info.Name(), ".txt") || strings.HasSuffix(info.Name(), ".md") || strings.HasSuffix(info.Name(), ".csv")) {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			content := string(data)
			paragraphs := strings.Split(content, "\n\n")
			for _, p := range paragraphs {
				pTrimmed := strings.TrimSpace(p)
				if len(pTrimmed) > 10 {
					r.chunks = append(r.chunks, KnowledgeChunk{
						Source:  info.Name(),
						Content: pTrimmed,
					})
				}
			}
		}
		return nil
	})
	return err
}

// SearchRelevantContext は質問文やコード文脈から関係のある知識を検索・結合して返却します
func (r *RAGManager) SearchRelevantContext(prompt string, maxResults int) string {
	var matched []string
	count := 0
	words := strings.Fields(strings.ToLower(prompt))

	for _, chunk := range r.chunks {
		matchScore := 0
		chunkLower := strings.ToLower(chunk.Content)
		for _, word := range words {
			if len(word) > 2 && strings.Contains(chunkLower, word) {
				matchScore++
			}
		}

		if matchScore > 0 || strings.Contains(chunkLower, strings.ToLower(prompt)) {
			matched = append(matched, fmt.Sprintf("[%s]: %s", chunk.Source, chunk.Content))
			count++
			if count >= maxResults {
				break
			}
		}
	}

	if len(matched) == 0 {
		return ""
	}
	return "=== 関連する現場・業務ナレッジコンテキスト ===\n" + strings.Join(matched, "\n---\n") + "\n============================================\n"
}
