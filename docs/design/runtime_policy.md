# Runtime Policy Design

## Purpose

Runtime policy defines how actgram executes generated automation with minimal token usage, clear safety boundaries, and vendor-neutral model selection.

The default operational target is no-token runtime.

AI should be used primarily for authoring, repair, manual generation, and exceptional judgment. Daily repetitive execution should rely on UWSCR, UIA, Win32, image matching, OCR, and deterministic rules.

## Runtime Modes

### Strict Offline Mode

- No external communication.
- No cloud LLM calls.
- No runtime token usage.
- Allowed mechanisms:
  - UWSCR
  - Win32
  - UI Automation
  - local image matching
  - local OCR if available
  - rule-based verification

### Hybrid Local Mode

- Allows local LLM only.
- Intended for low-risk local classification and short text reasoning.
- No cloud LLM calls.
- Suitable for:
  - OCR result cleanup
  - local error classification
  - simple yes/no judgment

### Cloud Assist Mode

- Allows cloud LLM only when policy permits.
- Cloud calls require explicit user approval by default.
- Suitable for:
  - complex exception analysis
  - script repair proposal
  - ambiguous visual judgment

### Authoring Mode

- Allows LLM usage for generation tasks.
- Suitable for:
  - initial UWSCR generation
  - ActionDSL normalization assistance
  - manual generation
  - workflow explanation
  - failure reflection

## Token Policy Schema

```json
{
  "runtime_default": "no_token",
  "allow_local_llm": true,
  "allow_cloud_llm": false,
  "cloud_llm_requires_user_approval": true,
  "max_runtime_ai_calls": 1,
  "max_runtime_ai_cost_jpy": 0,
  "record_ai_usage": true
}
```

## Decision Order

Runtime execution must use this order unless a workflow overrides it explicitly.

1. UWSCR deterministic operation
2. UIA / Win32 state query
3. image template matching
4. OCR + rules
5. local lightweight model
6. cloud model with approval
7. human intervention

## AI Usage Categories

### Preferred AI Usage

- initial script generation
- manual generation
- ActionDSL draft generation
- failure analysis
- repair proposal
- workflow improvement proposal

### Restricted AI Usage

- runtime screen judgment
- OCR semantic interpretation
- exception recovery

### Avoided AI Usage

- every-step action selection
- repeated daily deterministic operations
- high-risk action execution without approval

## Approval Gates

Approval is required for:

- high-risk business operation
- cloud LLM runtime call when not pre-approved
- workflow modification that affects production jobs
- destructive file operation
- external send operation

High-risk examples:

- delete
- send
- register
- approve
- confirm
- payment
- publish
- mail send

## Runtime Audit

Each run should produce a runtime report.

Recommended file:

```text
runs/run_YYYYMMDD_HHMMSS/runtime_report.json
```

Required fields:

- workflow_id
- jobnet_id if available
- trigger type
- start time
- end time
- status
- executed steps
- skipped steps
- failed steps
- approvals requested
- approvals accepted or rejected
- AI calls
- token estimate
- cost estimate
- screenshots on failure
- verification result

## Vendor Neutrality

Runtime policy must not depend on a single LLM vendor.

Provider interface should expose:

- GenerateText
- GenerateVisionText
- GenerateStructuredJSON
- EstimateCost
- SupportsLocal
- SupportsVision
- SupportsJSONSchema

The durable assets must be vendor-neutral:

- ActionDSL JSON
- UWSCR scripts
- Markdown manual
- HTML manual
- Mermaid diagram
- evidence logs
- runtime reports

## Failure Handling

Default failure order:

1. capture screenshot and evidence
2. run deterministic retry if configured
3. verify state again
4. use local classifier if allowed
5. ask user if unresolved
6. use cloud LLM only if approved
7. abort safely

## Production Default

For production workflows, the default should be:

```text
runtime_default = no_token
allow_local_llm = true
allow_cloud_llm = false
cloud_llm_requires_user_approval = true
```

This keeps daily operation fast, cheap, explainable, and safe.
