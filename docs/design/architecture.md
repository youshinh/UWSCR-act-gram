# actgram Architecture Design

## Purpose

actgram is designed as an evidence-driven, local-first RPA builder for Windows operations.

The goal is not merely to generate UWSCR scripts with AI. The goal is to convert human Windows operations into reusable business assets:

- evidence logs
- ActionDSL workflows
- UWSCR scripts
- operation manuals
- verification reports
- schedulable job networks

## Core Positioning

Computer Use style systems observe the screen and decide the next action at runtime. actgram takes a different position.

actgram uses AI mainly during authoring, repair, exception analysis, and documentation. Daily operation should run mostly without token usage by executing deterministic UWSCR scripts with UIA, Win32, image matching, OCR, and rule-based verification.

## Layered Architecture

```text
User Operation
  ↓
Evidence Layer
  ↓
ActionDSL Layer
  ↓
Generation Layer
  ↓
Workflow Layer
  ↓
Scheduler Layer
  ↓
Execution Layer
  ↓
Verification / Governance Layer
  ↓
Business Asset Layer
```

## Layer Responsibilities

### Evidence Layer

Records what happened on Windows as durable evidence, not only as raw mouse and keyboard events.

Sources:

- low-level mouse and keyboard hooks
- Win32 window metadata
- UI Automation element metadata
- UWSCR-assisted observation
- screenshots before and after operations
- target element crops
- OCR where needed
- clipboard and application context

### ActionDSL Layer

Converts raw evidence into normalized business operations.

Examples:

- click button
- input text
- wait for window
- export file
- send mail
- verify result
- request human approval

ActionDSL is the primary durable intermediate representation. UWSCR, manuals, and workflow definitions should be generated from ActionDSL, not directly from raw logs.

### Generation Layer

Generates deterministic artifacts from ActionDSL.

Generated outputs:

- `.uws` script
- Markdown manual
- HTML manual
- verification plan
- Mermaid diagram

LLM may assist generation, but the canonical input and output must remain vendor-neutral.

### Workflow Layer

Combines multiple scripts or ActionDSL units into larger flows.

Workflow is not limited to a single macro. It expresses business procedures composed of reusable units.

### Scheduler Layer

Runs workflows by schedule, event, manual trigger, API trigger, or dependency completion.

The scheduler must support retry, resource lock, queueing, approval gates, and monitoring.

### Execution Layer

Executes UWSCR scripts and fallback actions.

Preferred runtime order:

1. deterministic UWSCR
2. UIA / Win32
3. image matching
4. OCR and rules
5. local lightweight model
6. cloud model only with policy approval

### Verification / Governance Layer

Confirms that the expected state was reached.

It also handles:

- risk classification
- human approval
- token policy enforcement
- audit logging
- failure capture
- reflection and improvement proposals

### Business Asset Layer

Stores reusable operational knowledge.

A business asset may contain:

- evidence
- ActionDSL
- UWSCR
- manual
- workflow definition
- scheduler definition
- verification report
- historical success and failure data

## Runtime Principle

Daily runtime should be no-token by default.

AI is allowed primarily in these cases:

- initial script generation
- manual generation
- ambiguous exception handling
- failure analysis
- repair proposal
- workflow improvement proposal

## Design Principles

1. Evidence first
2. ActionDSL first
3. No-token runtime by default
4. Human-in-the-loop for risky actions
5. Vendor-neutral artifacts
6. Deterministic execution before AI judgment
7. Verifiable automation
8. Workflow and scheduler readiness
9. Business asset accumulation
