# GEMINI.md - Project Context (gplace)

## Project Overview
`gplace` is a modern Go client library and CLI tool for the **Google Places API (New)**. It is a specialized version of the original `goplaces` project, optimized for AI agent automation and enhanced metadata extraction.

### ⚠️ Billing Warning
By default, `gplace` requests fields from the **Pro** and **Enterprise** SKUs of the Places API (New) to provide rich metadata. Requests for reviews, amenities, and AI summaries trigger **Enterprise-tier** pricing.

### Key Features
- **Enhanced Details**: Fetches `userRatingCount`, `priceRange`, and detailed "serves" flags (Beer, Coffee, etc.).
- **Auto-Localization**: Supports the `--local` flag for automatic local language detection (two-pass lookup).
- **Summaries**: Supports AI-generated Review summaries (`review_summary`) and individual reviews.
- **AI Integration**: Specifically designed for AI agents (Gemini, Claude). Agents should use `SKILL.md` to automate complex research workflows.
- **Photos**: Removed to keep the CLI focused on data and speed.
- **Field Masks**: Uses selective masks to fetch only necessary data.

### Core Technologies
- **Language**: Go (1.24.0+)
- **CLI Framework**: Kong
- **API**: Google Places API (New) & Google Routes API

---

## Building and Running

### Commands
| Task | Command |
| :--- | :--- |
| **Build CLI** | `make gplace` (creates `./gplace` binary) |
| **Run Tests** | `make test` |
| **Lint Code** | `make lint` |

### Environment Variables
- `GOOGLE_PLACES_API_KEY`: Required API key.

---

## Development Conventions

### Project Structure
- `cmd/gplace/`: CLI entry point.
- `internal/cli/`: CLI logic and rendering.
- `root/`: Core library logic (`client.go`, `types.go`, etc.).

### Design Choices
- **Field Masks**: Uses selective field masks to fetch only requested data and minimize costs.
- **Mapping**: Decouples Google API payloads (`payloads.go`) from internal types (`types.go`) via `mapping.go`.
- **Lightweight**: Photos are excluded by design to focus on data extraction.
