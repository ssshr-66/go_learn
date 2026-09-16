<!-- TRELLIS:START -->
# Trellis Instructions

These instructions are for AI assistants working in this project.

This project is managed by Trellis. The working knowledge you need lives under `.trellis/`:

- `.trellis/workflow.md` — development phases, when to create tasks, skill routing
- `.trellis/spec/` — package- and layer-scoped coding guidelines (read before writing code in a given layer)
- `.trellis/workspace/` — per-developer journals and session traces
- `.trellis/tasks/` — active and archived tasks (PRDs, research, jsonl context)

If a Trellis command is available on your platform (e.g. `/trellis:finish-work`, `/trellis:continue`), prefer it over manual steps. Not every platform exposes every command.

If you're using Codex or another agent-capable tool, additional project-scoped helpers may live in:
- `.agents/skills/` — reusable Trellis skills
- `.codex/agents/` — optional custom subagents

Managed by Trellis. Edits outside this block are preserved; edits inside may be overwritten by a future `trellis update`.

<!-- TRELLIS:END -->

## Project-Specific Learning Instructions

This repository is a guided Go engineering learning workspace, not a production service.

- Keep the primary engineering-learning language as Go. Algorithm practice may remain in Java.
- Teach through small runnable demos. Do not begin by asking the learner to read a large open-source repository.
- For each demo, explain one concept, show a small example when needed, ask the learner to write a similar piece, then review the learner's code and add one incremental requirement.
- Do not provide a complete solution before the learner has attempted the current checkpoint. AI may provide syntax hints, debugging guidance, tests, and review.
- Record the weekly plan and acceptance criteria in the active weekly task. Record daily goals, code output, tests, blockers, and next steps in that task's `daily/` notes.
- Before moving to the next concept, require a runnable result and a short explanation of the key design choices.
- Read `.trellis/spec/go/index.md` before writing or reviewing Go demo code.
