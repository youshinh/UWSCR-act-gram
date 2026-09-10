package capture

// CaptureStatus records non-fatal capture/inspection errors for an evidence event.
type CaptureStatus struct {
	OK     bool     `json:"ok"`
	Errors []string `json:"errors,omitempty"`
}

// Rect describes screen coordinates in a vendor-neutral JSON shape.
type Rect struct {
	Left   int `json:"left"`
	Top    int `json:"top"`
	Right  int `json:"right"`
	Bottom int `json:"bottom"`
}

// WindowInfo describes the active window observed for an event.
type WindowInfo struct {
	Title string `json:"title"`
	Rect  *Rect  `json:"rect,omitempty"`
}

// UIAElementInfo describes the UI Automation element observed at a point.
type UIAElementInfo struct {
	ControlType  string `json:"control_type,omitempty"`
	AutomationID string `json:"automation_id,omitempty"`
	Value        string `json:"value,omitempty"`
}

// ImageEvidence stores paths for screenshots associated with an event.
type ImageEvidence struct {
	BeforePath string `json:"before_path,omitempty"`
	AfterPath  string `json:"after_path,omitempty"`
	MarkedPath string `json:"marked_path,omitempty"`
}

// EvidenceEvent is the durable JSONL evidence schema emitted during recording.
type EvidenceEvent struct {
	ID            string          `json:"id"`
	Timestamp     int64           `json:"timestamp"`
	Type          string          `json:"type"`
	X             int             `json:"x,omitempty"`
	Y             int             `json:"y,omitempty"`
	RelX          int             `json:"rel_x,omitempty"`
	RelY          int             `json:"rel_y,omitempty"`
	Key           string          `json:"key,omitempty"`
	Window        WindowInfo      `json:"window"`
	UIAElement    *UIAElementInfo `json:"uia_element,omitempty"`
	Images        *ImageEvidence  `json:"images,omitempty"`
	CaptureStatus CaptureStatus   `json:"capture_status"`
}

// EvidenceSession describes a recording session and its evidence outputs.
type EvidenceSession struct {
	SchemaVersion string `json:"schema_version"`
	StartedAt     int64  `json:"started_at"`
	EventsPath    string `json:"events_path"`
	CapturesDir   string `json:"captures_dir"`
	TemplatesDir  string `json:"templates_dir"`
}
