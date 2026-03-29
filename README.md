**Billbot** is cost calculator for **Claude Code**, **OpenAI Codex**, and **Cursor**. a small command-line tool that reads session files from your computer and prints **rough dollar estimates** for API-style token usage. Nothing is uploaded: it only opens local JSONL logs and Cursor’s SQLite stores, multiplies tokens by rates in a YAML file (built-in or yours), and shows tables or JSON.

Use it when you want a quick sense of spend after a long agent session, or a roll-up across everything the tool can find on disk not when you need an exact invoice.

> Totals are **estimates** from public list-style rates. Real bills depend on plans, seats, discounts, and how the vendor rounds.

<p align="center">
  <img src="assets/demo.gif" alt="Screen recording: billbot --help, then analyzing a sample Claude session with colors and prompts" width="920">
</p>

---

## What you get

- **Stays local** — no network calls, no API keys for analysis.
- **Three sources** — Anthropic (Claude Code) JSONL, Codex JSONL, Cursor `state.vscdb` / `store.db` (and Cursor JSONL).
- **Two ways to run it** — either “everything I can see” (default when you pass a directory), or “this one session” (a path to a file, or a directory with `--latest-only`).
- **Easy to script** — `--format json` for stable output.
- **Your pricing file** — optional `--pricing-file` when models or rates change.

## Setup

You need [Go](https://go.dev/) **1.23** (see `go.mod`).

```bash
go install github.com/geekmonkey/billbot/cmd/billbot@latest
```

Or clone and build:

```bash
go build -o billbot ./cmd/billbot
./billbot --help
./billbot --provider cc
```

## How to run it

`billbot`, `billbot analyze`, and `billbot project` all boil down to “analyze this path.” If you omit the path, it uses the current directory (`.`).

**1. Pass a directory, without `--latest-only` (the usual default)**  
billbot does **not** walk your repo. It looks in the normal places each tool stores logs (see below), groups results by **Anthropic / OpenAI / Cursor**, lists each session with subtotals, then prints **grand totals**. If a vendor has no files, that section is skipped. Broken files are mentioned on **stderr** and skipped.

**2. Pass a file, or a directory with `--latest-only`**  
You get one session: **per turn** (your question and usage for that turn) and **per completion** (each API call), plus JSON if you ask for it. For a directory, billbot picks the newest matching log (Claude tries to match your project path first; with `auto` it may fall back to another vendor—watch **stderr** for hints).

## Where it looks (by tool)

**Claude Code (`--provider anthropic`, or `cc` / `claude` / `claude-code`)**  
Session `*.jsonl` under `~/.claude/projects/…` and `~/.config/claude/projects/…` (including `sessions/` subfolders), plus global `~/.claude/sessions` and `~/.config/claude/sessions`. It ignores `history.jsonl` and, by default, `agent-*.jsonl` (use `--include-agents` to include those).

**Codex (`--provider openai` or `codex`)**  
All `*.jsonl` under `$CODEX_HOME/sessions` (recursive). If `CODEX_HOME` is unset, that’s **`~/.codex`**.

**Cursor (`--provider cursor`)**  
`state.vscdb` and `store.db` under workspace storage:

| OS | Folder |
|----|--------|
| macOS | `~/Library/Application Support/Cursor/User/workspaceStorage` |
| Windows | `%APPDATA%\Cursor\User\workspaceStorage` |
| Linux | `~/.config/Cursor/User/workspaceStorage` |

**Single file**  
Point at a `.jsonl` or Cursor DB path. With `--provider auto`, the path is used to guess the vendor; a plain `something.jsonl` with no hints is treated as Anthropic.

## Flags (the ones people change often)

| Flag | What it does |
|------|----------------|
| `--format table` or `json` | Human tables vs machine-readable output. |
| `--provider auto` … | Limit scans: `anthropic` (or `cc`, …), `openai` (or `codex`), `cursor`. |
| `--latest-only` | With a **directory**, analyze only one resolved session instead of everything. |
| `--pricing-file` | Use your YAML instead of the embedded defaults. |
| `--show-prompts` | In **table** mode, show prompt text in the tables (off by default). |
| `--prompt-width` | Max width for prompt columns (default `72`, clamped 8–32000). |
| `--color auto\|always\|never` | ANSI colors for tables; **`auto`** uses color on a TTY unless `NO_COLOR` is set. |
| `--no-color` | Same as `--color never`. |
| `--include-agents` | Include Claude `agent-*.jsonl` when scanning Claude paths. |

Colors apply only to **`--format table`**, not JSON.

## JSON shape (for scripts)

- **One session:** `meta`, `turns[]` (each with `answers[]`), and a flat `answers[]`.
- **Full aggregate:** `providers[]` (each with `sessions[]`, `totals`, …), `grand_totals`, and `meta.disclaimer`. Session rows can include `first_prompt` when present.

## Pricing quirks

Rates live in **`internal/pricing/default.yaml`** (USD per **million** tokens, including optional cache lines). If the log names a model that isn’t in the table (and isn’t aliased), that line shows **$0** and aggregates count it in `unknown_model_answers`.

For **Anthropic**, short model ids also get a **prefix fallback** (e.g. `claude-sonnet-4…` maps to a known Sonnet row). **OpenAI** and **Cursor** need a matching or aliased name; **Cursor** can map internal names via `cursor.model_aliases` to OpenAI model keys in the same file.

## Example commands

```bash
# All sessions billbot can find, split by vendor
./billbot

# Only Claude logs
./billbot --provider anthropic .

# One session tied to a repo (newest match)
./billbot --latest-only ~/src/my-app

# Explicit file
./billbot analyze ~/.claude/projects/-Users-you-myapp/session-abc.jsonl

# Tables with prompts and forced colors (nice in recordings / some terminals)
./billbot analyze testdata/anthropic/min.jsonl --show-prompts --color always

# Pipe-friendly
./billbot analyze session.jsonl --format json
```

## Custom pricing

Copy **`internal/pricing/default.yaml`**, adjust `models` (and `model_aliases` if you rename ids). For Cursor-specific strings, use `cursor.model_aliases` pointing at OpenAI model keys.

```bash
./billbot --pricing-file ~/config/billbot-pricing.yaml
```

## Regenerating the demo GIF

The animation is recorded with [VHS](https://github.com/charmbracelet/vhs). Install **`vhs`**, **`ffmpeg`**, and **`ttyd`** (Homebrew: `brew install vhs ffmpeg ttyd`), then from the **repository root**:

```bash
vhs demo/demo.tape
```

That overwrites **`assets/demo.gif`**.

## Development

```bash
go test ./...
```
