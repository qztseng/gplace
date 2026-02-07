# GEMINI.md - Project Context (gplace)

## Project Overview
`gplace` is a modern Go client library and CLI tool for the **Google Places API (New)**. It is a specialized version of the original `goplaces` project, optimized for CLI use with enhanced metadata support.

### Key Features
- **Enhanced Details**: Fetches `userRatingCount`, `priceRange`, and detailed "serves" flags (Beer, Coffee, etc.).
- **Summaries**: Supports Editorial, AI-generated (Generative), and Review summaries.
- **Price Representation**: Renders numeric price levels as repeated dollar signs (e.g., `$$`).
- **No Photos**: Photo fetching has been removed to keep the CLI tool focused on textual metadata and scripting efficiency.
- **AI Integration**: Designed to be used by AI agents. Agents can be instructed to create `SKILL.md` files to automate tasks using `gplace`.

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
