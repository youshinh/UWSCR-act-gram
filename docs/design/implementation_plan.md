# Implementation Plan for Codex

## Purpose

This file is the handoff plan for Codex or another coding agent.

Read these design documents first:

- `docs/design/architecture.md`
- `docs/design/evidence_layer.md`
- `docs/design/actiondsl_schema.md`
- `docs/design/runtime_policy.md`
- `docs/design/jobnet_scheduler.md`

The project direction is:

```text
Evidence-driven, local-first, low-token Windows business automation builder
```

Do not optimize only for AI script generation. Build the foundation for durable operational assets:

- evidence logs
- ActionDSL
- UWSCR scripts
- manuals
- verification reports
- JobFlow / JobNet scheduler definitions
- Mermaid exports

## Development Rules

1. Prefer small, reviewable commits.
2. Keep existing app behavior working.
3. Do not remove current UI features unless explicitly replaced.
4. Keep runtime no-token by default.
5. Use deterministic code generation where possible.
6. Use LLM only for authoring, repair, and optional exception analysis.
7. Add tests for pure Go packages where feasible.
8. Keep data schemas vendor-neutral JSON.

## Phase 1: Evidence Recorder MVP

Goal: make recorder output reliable evidence instead of fragile event logs.

### Tasks

- Add `capture/evidence.go`.
- Define `EvidenceEvent`, `WindowInfo`, `UIAElementInfo`, `ImageEvidence`, `CaptureStatus` structs.
- Change recording output to include:
  - `session.json`
  - `events.jsonl`
  - `captures/`
  - `templates/`
- Write each event to JSONL immediately.
- Capture click before/after screenshots.
- Preserve capture errors in event JSON.
- Drain all queued events before stop.
- Keep current `log.json` generation temporarily for backward compatibility.

### Acceptance Criteria

- A recording session produces `events.jsonl`.
- Click events contain window info and image paths when capture succeeds.
- Capture failure does not drop the event.
- StopRecording does not lose final events.
- Existing manual generation path is not broken.

## Phase 2: UIA and Window Evidence Improvements

Goal: improve target identification.

### Tasks

- Add `capture/win32.go` for window metadata.
- Expand active window capture:
  - HWND
  - title
  - class name
  - process id
  - process name if possible
  - rect
- Replace or wrap current `inspectElementAtPoint` to return richer `UIAElementInfo`.
- Add parent chain where feasible.
- Add clear error reporting when UIA fails.

### Acceptance Criteria

- Evidence event contains stable window metadata.
- UIA failure is visible in `capture_status.errors`.
- Click target can be reasoned from name, automation id, control type, and bounding rect.

## Phase 3: ActionDSL Package

Goal: introduce canonical intermediate representation.

### Tasks

- Add package `actiondsl`.
- Define Go structs matching `docs/design/actiondsl_schema.md`.
- Add JSON marshal/unmarshal helpers.
- Add basic validator:
  - unique step ids
  - valid action names
  - valid risk values
  - high risk requires confirmation
  - fallback strategy validation
- Add normalizer skeleton:
  - Evidence events -> basic ActionDSL steps
  - click -> click step
  - key/text sequence -> input_text or hotkey step

### Acceptance Criteria

- `go test ./actiondsl` passes.
- A sample evidence JSONL can be converted into a basic ActionDSL JSON file.
- High-risk action without confirmation fails validation.

## Phase 4: UWSCR Generator MVP

Goal: generate deterministic UWSCR from ActionDSL.

### Tasks

- Add package `generator` or extend existing generation area.
- Implement ActionDSL -> UWSCR template generator.
- Support actions:
  - click
  - input_text
  - press_key
  - hotkey
  - wait_window
  - verify_file
  - human confirmation
- Generation priority:
  1. `clkitem` / UIA name
  2. image template fallback
  3. window-relative coordinate fallback
- Add static checker:
  - block `CTRL_WIN`
  - block `STATUS_X` / `STATUS_Y`
  - check obvious `mouseorg(id)` / `mouseorg(0)` pairing

### Acceptance Criteria

- Sample ActionDSL generates a runnable `.uws` file.
- Generated click steps prefer `clkitem` where possible.
- High-risk step inserts confirmation prompt.
- Static checker catches known invalid UWSCR tokens.

## Phase 5: Manual Generator from ActionDSL

Goal: generate manuals from stable ActionDSL, not raw LLM output.

### Tasks

- Generate Markdown manual from ActionDSL.
- Include:
  - purpose
  - preconditions
  - inputs
  - steps
  - screenshots when available
  - verification points
  - risk and approval points
  - troubleshooting
- Keep optional LLM polishing separate from deterministic manual generation.

### Acceptance Criteria

- A workflow produces `manual.md` without LLM.
- HTML generation can reuse existing manual functionality.
- Manual contains evidence image links when available.

## Phase 6: Runtime Policy Enforcement

Goal: keep daily operation low-cost and safe.

### Tasks

- Add runtime policy structs.
- Add default policies:
  - strict_offline
  - hybrid_local
  - cloud_assist
  - authoring
- Log AI calls in runtime report.
- Require approval before cloud runtime call by default.

### Acceptance Criteria

- Runtime mode can be serialized and loaded.
- Production default disallows cloud runtime calls.
- Runtime report records token/cost estimate fields even when zero.

## Phase 7: JobFlow / JobNet MVP

Goal: support higher-level business automation.

### Tasks

- Add package `jobnet` or `scheduler`.
- Define structs:
  - JobNet
  - JobFlow
  - Job
  - Trigger
  - RetryPolicy
  - ResourceLock
- Validate JobFlow DAG:
  - no missing dependency
  - no cycle
  - unique job ids
- Implement Mermaid export.
- Implement sequential DAG executor skeleton.
- Support job types:
  - uwscr
  - actiondsl
  - verify_file
  - human_task placeholder

### Acceptance Criteria

- `go test ./jobnet` or `go test ./scheduler` passes.
- Sample JobFlow exports valid Mermaid markdown.
- Sample JobFlow validates dependencies.
- Executor can run or dry-run jobs in dependency order.

## Phase 8: GUI Integration

Goal: surface the new architecture in the Wails UI.

### Tasks

- Add record session viewer for evidence files.
- Add ActionDSL preview/export.
- Add UWSCR generation button from ActionDSL.
- Add Mermaid export button for JobFlow/JobNet.
- Add runtime mode display.
- Add confirmation UI for high-risk actions.

### Acceptance Criteria

- User can record operation and inspect evidence.
- User can generate ActionDSL and UWSCR from a session.
- User can export Mermaid from a JobFlow/JobNet.

## Suggested First Codex Task

Start with Phase 1 only.

Recommended prompt:

```text
Implement Phase 1 from docs/design/implementation_plan.md.
Add evidence structs, JSONL event writing, session.json, before/after click screenshots, capture error preservation, and safe stop/drain behavior. Keep current log.json compatibility. Run gofmt and relevant tests/build checks.
```

## Non-Goals for Initial Work

Do not implement these first:

- full scheduler UI
- distributed execution
- cloud LLM optimizer
- complex OCR pipeline
- browser DOM integration
- Excel COM integration

These are future extensions after Evidence and ActionDSL are stable.
