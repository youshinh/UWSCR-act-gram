# ActionDSL Schema Design

## Purpose

ActionDSL is the canonical intermediate representation between raw evidence and generated artifacts.

Raw recorder logs are too unstable to generate UWSCR scripts directly. ActionDSL normalizes those logs into business-level operations that can be converted into UWSCR scripts, manuals, verification rules, and job flows.

## Design Goals

- Keep daily runtime mostly token-free.
- Avoid vendor lock-in by storing deterministic JSON assets.
- Separate evidence, intent, execution strategy, fallback, risk, and verification.
- Support UWSCR generation, manual generation, and workflow scheduling from the same source.
- Allow human approval gates for risky actions.

## Top-Level Workflow Schema

```json
{
  "schema_version": "1.0",
  "workflow_id": "wf_order_export",
  "name": "Export order list",
  "description": "Search orders and export the result as CSV.",
  "risk_level": "medium",
  "token_policy": {
    "runtime_default": "no_token",
    "allow_local_llm": true,
    "allow_cloud_llm": false,
    "cloud_llm_requires_user_approval": true,
    "max_runtime_ai_calls": 1
  },
  "inputs": [
    { "name": "order_no", "type": "string", "required": true }
  ],
  "outputs": [
    { "name": "output_csv_path", "type": "file" }
  ],
  "preconditions": [
    { "type": "window_exists", "title_contains": "Order" }
  ],
  "steps": [],
  "postconditions": [
    { "type": "file_exists", "path": "${output_csv_path}" }
  ]
}
```

## Step Schema

```json
{
  "step_id": "step_001",
  "intent": "Click the Search button",
  "action": "click",
  "target": {
    "strategy": "uia",
    "window_title_contains": "Order",
    "automation_id": "btnSearch",
    "name": "Search",
    "control_type": "Button"
  },
  "value": null,
  "fallback": [
    { "strategy": "clkitem", "name": "Search" },
    { "strategy": "image_template", "template": "templates/search_button.png" },
    { "strategy": "relative_click", "x": 1020, "y": 160 }
  ],
  "verify": {
    "type": "state_changed",
    "condition": "search_result_updated"
  },
  "risk": "low",
  "requires_confirmation": false,
  "evidence_refs": ["evt_20260710_000001"]
}
```

## Supported Actions

### Primitive UI Actions

- click
- double_click
- right_click
- input_text
- press_key
- hotkey
- select_combo
- check
- uncheck
- scroll
- drag
- wait
- wait_window
- wait_element
- copy
- paste

### Business Actions

- search
- open_record
- save
- export
- import
- print
- send
- delete
- register
- approve
- confirm
- cancel

### Verification Actions

- verify_window
- verify_text
- verify_element
- verify_image
- verify_file
- verify_clipboard
- verify_ocr
- verify_ai_eval

### Branching Actions

- if_element_exists
- if_text_exists
- if_image_exists
- if_file_exists
- if_ai_judgement

### Error Handling Actions

- on_error_retry
- on_error_screenshot
- on_error_ask_user
- on_error_ai_eval
- on_error_abort

## Target Strategy Priority

UWSCR generation should prefer stable selectors over coordinates.

Priority:

1. UIA / AutomationId
2. UIA / Name + ControlType
3. UWSCR `clkitem`
4. Image template matching
5. Window-relative coordinate
6. Absolute coordinate as last resort

## Risk Levels

### low

- search
- display
- copy
- navigation
- temporary input

### medium

- save
- overwrite
- export
- print
- file move

### high

- delete
- send
- register
- approve
- confirm
- payment
- publish
- mail send

High-risk steps must set `requires_confirmation` to true unless explicitly disabled by a governance policy.

## Confirmation Schema

```json
{
  "risk": "high",
  "requires_confirmation": true,
  "confirmation": {
    "message": "Execute order confirmation.",
    "summary_fields": ["supplier_name", "order_no", "amount"],
    "approval_mode": "manual_click"
  }
}
```

## Verification Schema

```json
{
  "type": "uia_value_equals",
  "target_automation_id": "txtOrderNo",
  "expected": "${order_no}",
  "level": "strict",
  "on_failure": "retry_or_abort"
}
```

Verification levels:

- strict: exact value, file, or UIA state match
- normal: expected text, element, or state change detected
- soft: OCR or AI-assisted judgement

## Artifact Relationship

```text
Evidence events
  ↓
ActionDSL workflow
  ↓
UWSCR script
Manual
Verification plan
JobFlow
JobNet
Mermaid diagram
```

ActionDSL must remain the source of truth for generated operational assets.
