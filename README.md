# 📍 gplace — Modern Google Places for AI & Humans

`gplace` is a high-performance Go client and CLI tool built exclusively for the **Google Places API (New)**. It is a specialized, feature-rich version of the original `goplaces` project, optimized for AI agent automation and deep metadata extraction.

---

## ⚠️ Important: Billing & SKUs
This tool is designed to fetch **comprehensive metadata**. By default, it requests fields that fall into the higher pricing tiers of the Google Places API (New). You are billed for the highest SKU applicable to your request.

| SKU Tier | Included Data |
| :--- | :--- |
| **Essentials** | ID, Address, Location (Lat/Lng), Types. |
| **Pro** | **Ratings, User Rating Count**, Price Level, AI Summaries. |
| **Enterprise** | **Reviews**, Opening Hours, Amenities (e.g., `serves_beer`, `serves_coffee`). |

**Note**: Using the `--reviews` flag or fetching full details will trigger **Enterprise-tier** billing. Use responsibly and monitor your Google Cloud Console.

---

## Highlights

- **Google Places API (New)**: Built on the latest v1 API with field-masking for speed and cost-control.
- **Auto-Localization**: Use the `--local` flag to automatically detect a place's country and fetch metadata (names, summaries, reviews) in the native local language.
- **AI Synthesis Ready**: Designed to be used as a "Skill" for AI agents (Gemini, Claude, GPT). It provides structured JSON output and a dedicated `SKILL.md` workflow.
- **Deep Metadata**: 
  - **Atmosphere**: Detailed "serves" flags (Breakfast, Brunch, Beer, Wine, etc.).
  - **Summaries**: Fetches `review_summary` (AI-generated vibe) and individual reviews.
  - **Pricing**: Numeric `price_level` (rendered as `$$`) and detailed `price_range` with currency.
- **Tidy & Fast**: Colorized human-readable output and compact JSON for scripting. *Photos are excluded to keep the tool focused on data and speed.*

---

## Installation

### From Source
```bash
make gplace
# The binary is created as ./gplace
```

### Global Install
```bash
go install github.com/qztseng/gplace/cmd/gplace@latest
```

---

## Configuration
Set your Google Cloud API Key in your environment:
```bash
export GOOGLE_PLACES_API_KEY="your_api_key_here"
```

---

## Usage

### 1. AI-Driven Localization (`--local`)
Fetch the **most relevant local reviews** and localized names without knowing the language code:
```bash
gplace details ChIJYdTD1o2LGGAR_8lyKP44pBM --local --reviews
```

### 2. Global Search
Search with filters like minimum rating and price level:
```bash
gplace search "specialty coffee" --min-rating 4.5 --price-level 2 --limit 5
```

### 3. Route-based Discovery
Find places (e.g., gas stations, cafes) along a driving route:
```bash
gplace route "gas station" --from "Tokyo" --to "Osaka" --json
```

---

## AI Agent Integration (SKILL.md)
This project includes a [SKILL.md](./SKILL.md) file. This is a special set of instructions for AI agents. If you are using a tool like **Gemini CLI**, you can point the agent to this repository, and it will automatically know how to:
1.  Map user requests to official Google Place Types.
2.  Detect local languages for authentic review analysis.
3.  Synthesize recommended dishes and "vibes" from combined AI summaries and user reviews.

---

## Development & Testing
```bash
make test  # Runs unit tests
make lint  # Runs golangci-lint
```

## License
MIT. Portions based on the original `goplaces` project.
