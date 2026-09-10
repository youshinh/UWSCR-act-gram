package manual

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
)

type ManualStep struct {
	StepID        int    `json:"step_id"`
	StepNumber    int    `json:"step_number,omitempty"`
	Title         string `json:"title"`
	Instruction   string `json:"instruction"`
	Description   string `json:"description,omitempty"`
	WindowTitle   string `json:"window_title,omitempty"`
	TargetElement string `json:"target_element,omitempty"`
	ActionType    string `json:"action_type,omitempty"` // "click", "input", "shortcut", "wait"
	InputValue    string `json:"input_value,omitempty"`
	ClickX        int    `json:"click_x"`
	ClickY        int    `json:"click_y"`
	RelX          int    `json:"rel_x,omitempty"`
	RelY          int    `json:"rel_y,omitempty"`
	UWSCode       string `json:"uws_code"`
	ImagePath     string `json:"image_path"`
	AudioScript   string `json:"audio_script,omitempty"`
}

type ManualSession struct {
	mu          sync.Mutex
	CurrentStep int
	Steps       []ManualStep
	IsRunning   bool
}

func NewManualSession() *ManualSession {
	return &ManualSession{
		CurrentStep: 0,
		Steps:       make([]ManualStep, 0),
		IsRunning:   false,
	}
}

// SetSteps はステップ一覧を設定・更新します
func (ms *ManualSession) SetSteps(steps []ManualStep) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.Steps = steps
	ms.renumberStepsLocked()
}

// GetSteps は現在のステップ一覧を返します
func (ms *ManualSession) GetSteps() []ManualStep {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	copied := make([]ManualStep, len(ms.Steps))
	copy(copied, ms.Steps)
	return copied
}

// UpdateStep は指定インデックスのステップを更新します
func (ms *ManualSession) UpdateStep(index int, step ManualStep) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if index < 0 || index >= len(ms.Steps) {
		return fmt.Errorf("インデックス範囲外です (index: %d)", index)
	}
	step.StepID = index + 1
	step.StepNumber = index + 1
	if step.Description == "" {
		step.Description = step.Instruction
	}
	if step.Instruction == "" {
		step.Instruction = step.Description
	}
	ms.Steps[index] = step
	return nil
}

// AddStep は末尾または指定位置に新しいステップを追加します
func (ms *ManualSession) AddStep(step ManualStep, afterIndex int) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if step.Description == "" {
		step.Description = step.Instruction
	}
	if step.Instruction == "" {
		step.Instruction = step.Description
	}

	if afterIndex < 0 || afterIndex >= len(ms.Steps) {
		ms.Steps = append(ms.Steps, step)
	} else {
		ms.Steps = append(ms.Steps[:afterIndex+1], append([]ManualStep{step}, ms.Steps[afterIndex+1:]...)...)
	}
	ms.renumberStepsLocked()
}

// DeleteStep は指定インデックスのステップを削除します
func (ms *ManualSession) DeleteStep(index int) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if index < 0 || index >= len(ms.Steps) {
		return fmt.Errorf("インデックス範囲外です (index: %d)", index)
	}
	ms.Steps = append(ms.Steps[:index], ms.Steps[index+1:]...)
	ms.renumberStepsLocked()
	if ms.CurrentStep >= len(ms.Steps) && len(ms.Steps) > 0 {
		ms.CurrentStep = len(ms.Steps) - 1
	}
	return nil
}

// MoveStep はステップの順序を移動します
func (ms *ManualSession) MoveStep(fromIndex, toIndex int) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if fromIndex < 0 || fromIndex >= len(ms.Steps) || toIndex < 0 || toIndex >= len(ms.Steps) {
		return fmt.Errorf("無効な移動インデックスです")
	}
	item := ms.Steps[fromIndex]
	ms.Steps = append(ms.Steps[:fromIndex], ms.Steps[fromIndex+1:]...)
	ms.Steps = append(ms.Steps[:toIndex], append([]ManualStep{item}, ms.Steps[toIndex:]...)...)
	ms.renumberStepsLocked()
	return nil
}

// renumberStepsLocked はステップ番号を1始まりで再採番します (mu保有前提)
func (ms *ManualSession) renumberStepsLocked() {
	for i := range ms.Steps {
		ms.Steps[i].StepID = i + 1
		ms.Steps[i].StepNumber = i + 1
		if ms.Steps[i].Description == "" {
			ms.Steps[i].Description = ms.Steps[i].Instruction
		}
		if ms.Steps[i].Instruction == "" {
			ms.Steps[i].Instruction = ms.Steps[i].Description
		}
	}
}

// SaveToFile はステップ一覧を scenario.json として保存します
func (ms *ManualSession) SaveToFile(filePath string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	data, err := json.MarshalIndent(ms.Steps, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

// LoadFromFile は scenario.json からステップ一覧を復元します
func (ms *ManualSession) LoadFromFile(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var steps []ManualStep
	if err := json.Unmarshal(data, &steps); err != nil {
		return err
	}

	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.Steps = steps
	ms.renumberStepsLocked()
	return nil
}

// BuildCombinedScript は全ステップの UWSCR コードを1本の連続実行可能なスクリプトに結合します
func (ms *ManualSession) BuildCombinedScript() string {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	var sb strings.Builder
	sb.WriteString("// ========================================================\n")
	sb.WriteString("// UWSCR 統合自動実行ワークフロー\n")
	sb.WriteString(fmt.Sprintf("// 生成ステップ数: %d\n", len(ms.Steps)))
	sb.WriteString("// ========================================================\n\n")

	for i, step := range ms.Steps {
		sb.WriteString(fmt.Sprintf("// --- STEP %d: %s ---\n", i+1, step.Title))
		if step.Instruction != "" {
			sb.WriteString(fmt.Sprintf("// 指示: %s\n", step.Instruction))
		}
		if step.WindowTitle != "" {
			sb.WriteString(fmt.Sprintf("// 対象ウィンドウ: %s\n", step.WindowTitle))
		}

		code := strings.TrimSpace(step.UWSCode)
		if code != "" {
			sb.WriteString(code)
			sb.WriteString("\n")
		} else {
			sb.WriteString("// (コード定義なし)\n")
		}
		sb.WriteString("SLEEP(0.5)\n\n")
	}

	return sb.String()
}

// ExecuteStep は指定されたステップを同期実行します
func (ms *ManualSession) ExecuteStep(stepIdx int, runSync func(code string) (string, bool, error)) (*ManualStep, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if stepIdx < 0 || stepIdx >= len(ms.Steps) {
		return nil, fmt.Errorf("ステップ番号が範囲外です: %d", stepIdx)
	}

	ms.CurrentStep = stepIdx
	step := ms.Steps[stepIdx]

	if runSync != nil && step.UWSCode != "" {
		_, _, err := runSync(step.UWSCode)
		if err != nil {
			return nil, err
		}
	}

	return &step, nil
}
