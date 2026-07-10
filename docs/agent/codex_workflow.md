# Codex Workflow

This document defines how Codex should work on this repository.

## First Step for Every Task

1. Check repository state.
2. Read `README.md`.
3. Read `docs/design/implementation_plan.md`.
4. Read related files under `docs/design/`.
5. Inspect the target source files before editing.

If `origin` is not configured in the checkout, continue with the current checkout and report that `git fetch origin main` could not be executed.

## Required Design Documents

- `docs/design/implementation_plan.md`
- `docs/design/architecture.md`
- `docs/design/evidence_layer.md`
- `docs/design/actiondsl_schema.md`
- `docs/design/runtime_policy.md`
- `docs/design/jobnet_scheduler.md`

## Implementation Rules

- Keep changes small and reviewable.
- Preserve existing UI and recorder behavior unless explicitly replacing it.
- Do not block low-level Windows hook callbacks.
- Put long-running capture/enrichment work in worker goroutines.
- Keep evidence output append-only where possible.
- Preserve legacy output formats until a migration is explicitly implemented.
- Prefer deterministic generation over runtime AI calls.
- Default runtime behavior must be no-token / local-first.

## Testing Rules

Run the narrowest relevant checks first.

Recommended order:

```bash
gofmt -w <changed-go-files>
go test ./capture
go test ./...
```

When testing on Linux, Windows-specific packages may fail because of Win32 APIs such as `syscall.NewLazyDLL`. In that case:

- report the failure clearly
- do not treat it as proof the implementation is broken
- run package-level tests that can be cross-compiled or isolated

## Commit Rules

Commit after each coherent unit of work.

Use clear messages, for example:

- `feat: add evidence recorder artifacts`
- `fix: drain recorder event queue on stop`
- `refactor: separate evidence events from legacy log events`
- `docs: update implementation plan`

## Completion Report Format

At the end of each task, report:

```text
Summary
- What changed
- Why it changed
- Important design decisions

Testing
- Commands run
- Passed checks
- Failed checks and exact reason

Files
- Changed files

Next
- Recommended next task
```

## Current Priority

Continue from `docs/design/implementation_plan.md` Phase 1 unless explicitly instructed otherwise.
