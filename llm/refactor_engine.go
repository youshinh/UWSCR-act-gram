package llm

import (
	"fmt"
)

type ExecutionEvent struct {
	StepID    string `json:"step_id"`
	Command   string `json:"command"`
	StartTime int64  `json:"start_time"` // ミリ秒（T1）
	EndTime   int64  `json:"end_time"`   // ミリ秒（T2）
	IdleTime  int64  `json:"idle_time"`  // ミリ秒（T3）
}

type RefactorEngine struct{}

func NewRefactorEngine() *RefactorEngine {
	return &RefactorEngine{}
}

func (re *RefactorEngine) ProposeOptimization(originalCode string, events []ExecutionEvent, ragContext string, provider LLMProvider, model string) (string, error) {
	var eventsStr string
	for _, e := range events {
		eventsStr += fmt.Sprintf("- StepID: %s, Command: %s, Start: %d ms, End: %d ms, Idle: %d ms\n", e.StepID, e.Command, e.StartTime, e.EndTime, e.IdleTime)
	}

	prompt := fmt.Sprintf(`
あなたはUWSCRおよびWindowsシステムプログラミングの天才アーキテクトです。
提示された「元のコード」と、実行時に計測された「ミリ秒単位のファクトログ」を厳密に分析し、極限まで無駄を削ぎ落とした「非同期・並行処理化またはミリ秒ループ監視に最適化されたUWSCRコード」を提案してください。

元のコード:
"""
%s
"""

実測ファクトデータ（ミリ秒タイムライン）:
%s

%s

【ミッション】
1. 因果関係（データ依存性）の厳密な分析
2. クリティカルパスの再構築（固定SLEEPの排除、ミリ秒ループ監視への書き換え、非同期事前タスク起動）
3. 返却は、必ず以下の形式の純粋なJSONオブジェクトのみにしてください（マークダウンのコードブロック 'json ... ' や説明文、前書きなどは一切含めないでください）：

{
  "estimated_time_saved_ms": 削れると推定されるミリ秒数,
  "bottlenecks": ["指摘1", "指摘2"],
  "refactored_code": "リファクタリング後の完全なUWSCRコード"
}
`, originalCode, eventsStr, ragContext)

	resp := provider.GenerateText(prompt, "", model)
	if resp.Error != nil {
		return "", resp.Error
	}
	return resp.Text, nil
}
