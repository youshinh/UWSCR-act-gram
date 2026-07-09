# Evidence Layer Design

## Purpose

The Evidence Layer records Windows operations as durable, analyzable evidence.

The system must avoid relying only on mouse coordinates or screenshots. Each operation should capture enough context to later generate UWSCR scripts, manuals, verification rules, and workflow assets.

## Event Storage

Events should be written incrementally as JSON Lines.

Recommended file:

```text
sessions/recording_YYYYMMDD_HHMMSS/events.jsonl
```

JSONL is preferred because it preserves partial data even if the application crashes during recording.

## Required Event Categories

### Input Events

- mouse_down
- mouse_up
- click
- double_click
- right_click
- wheel
- drag_start
- drag_move
- drag_end
- key_down
- key_up
- text_input
- hotkey
- clipboard_change

### Window Evidence

- hwnd
- title
- class_name
- process_id
- process_name
- exe_path
- window_rect
- client_rect
- monitor_id
- dpi_scale
- is_foreground
- is_visible
- is_enabled

### UI Automation Evidence

- name
- automation_id
- class_name
- control_type
- localized_control_type
- runtime_id
- bounding_rect
- is_enabled
- is_offscreen
- value
- legacy_value
- supported_patterns
- clickable_point
- parent_chain
- children_summary

### Screenshot Evidence

For click-like operations, capture at least:

- screenshot_before
- screenshot_after
- active_window_capture
- target_element_crop
- template_image

The before/after pair is essential because it enables state-change analysis.

### UWSCR-Assisted Evidence

UWSCR can be used not only for execution but also for observation.

Candidate observations:

- GETID result
- STATUS(id, ST_X)
- STATUS(id, ST_Y)
- STATUS(id, ST_WIDTH)
- STATUS(id, ST_HEIGHT)
- CHKIMG-style image matching result
- CLKITEM feasibility

### OCR Evidence

OCR should be optional and event-triggered.

Use OCR when:

- UIA information is weak
- text exists only in bitmap regions
- error or confirmation dialogs appear
- a verifier needs textual evidence
- AI judgment should receive compressed text instead of a full image

### Application-Specific Evidence

Where possible, collect richer context:

| Application | Source | Evidence |
|---|---|---|
| Excel | COM | workbook, sheet, cell, selection |
| Word | COM | document, selection |
| Outlook | COM | mail item, subject, recipients |
| Explorer | Shell API | current folder, selected items |
| Browser | UIA / Accessibility | URL, title, focused element |
| Business app | UIA / Win32 / OCR | UI tree, selected row, visible controls |

## Evidence Event Schema

```json
{
  "schema_version": "1.0",
  "event_id": "evt_20260710_000001",
  "session_id": "rec_20260710_153000",
  "timestamp": 1783674600123,
  "type": "click",
  "input": {
    "x": 1200,
    "y": 680,
    "button": "left",
    "click_count": 1,
    "key": "",
    "text": ""
  },
  "window": {
    "hwnd": "0x001A0920",
    "title": "Order Management",
    "class_name": "WindowsForms10.Window",
    "process_id": 12345,
    "process_name": "OrderSystem.exe",
    "exe_path": "C:\\Apps\\OrderSystem\\OrderSystem.exe",
    "rect": { "x": 100, "y": 80, "w": 1400, "h": 900 },
    "dpi_scale": 1.25
  },
  "uia": {
    "name": "Search",
    "automation_id": "btnSearch",
    "class_name": "Button",
    "control_type": "Button",
    "runtime_id": "42.100.5.8",
    "rect": { "x": 1120, "y": 640, "w": 120, "h": 40 },
    "is_enabled": true,
    "is_offscreen": false,
    "value": "",
    "supported_patterns": ["InvokePattern"],
    "parent_chain": [
      { "name": "Search Conditions", "control_type": "Group" },
      { "name": "Order Management", "control_type": "Window" }
    ]
  },
  "images": {
    "before": "captures/evt_000001_before.png",
    "after": "captures/evt_000001_after.png",
    "target_crop": "captures/evt_000001_target.png",
    "template": "templates/search_button.png"
  },
  "ocr": {
    "target_text": "Search",
    "window_text_summary": "Order No Customer Search Clear"
  },
  "capture_status": {
    "uia_success": true,
    "screenshot_success": true,
    "ocr_success": false,
    "errors": []
  }
}
```

## Recorder Requirements

- write `session.json` when recording starts
- write each event to `events.jsonl` immediately
- capture before/after screenshots for click-like events
- keep capture errors inside the event instead of silently dropping evidence
- drain all queued events before stopping
- never lose the last events on F8 or UI-triggered stop
- keep raw events and normalized actions separate
