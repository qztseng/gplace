# 📍 gplace — Modern Places for Go

Modern Go client + CLI for the Google Places API (New). Fast for humans, tidy for scripts.

> **Note**: This is a fork and renamed version of the original `goplaces` project, specifically enhanced with detailed place attributes and optimized for CLI usage.

## Highlights

- **Text Search**: Search with filters (keyword, type, open now, min rating, price levels).
- **Autocomplete**: Suggestions for places + queries (session tokens supported).
- **Nearby Search**: Find places around a location restriction.
- **Route Search**: Search for places along a driving path (Routes API).
- **Detailed Place Info**: 
  - **Ratings**: includes `user_rating_count` (e.g., `4.9 (1515)`).
  - **Price Info**: Numeric price levels (rendered as `$$`) and specific `price_range` data.
  - **Summaries**: Editorial summaries, AI-generated overviews, and review summaries.
  - **Amenities**: Detailed "serves" flags (Beer, Breakfast, Brunch, Cocktails, Coffee, Dessert, Dinner, Lunch, Vegetarian, Wine).
  - **Status & Maps**: Business status and direct Google Maps URIs.
- **AI Ready**: You can ask an AI bot (like Gemini or Claude) to implement a specific `SKILL.md` to automate workflows using this `gplace` command.
- **Tidy Output**: Colorized human output + clean `--json` mode for scripting. (No photos included to keep CLI lightweight).

## Install / Run

- Go: `go install github.com/qztseng/gplace/cmd/gplace@latest`
- Source: `make gplace`

## Config

```bash
export GOOGLE_PLACES_API_KEY="..."
```

## CLI

```text
gplace [--api-key=KEY] [--json] [--no-color] <command>

Commands:
  autocomplete  Autocomplete places and queries.
  nearby        Search nearby places by location.
  search        Search places by text query.
  route         Search places along a route.
  details       Fetch place details by place ID.
  resolve       Resolve a location string to candidate places.
```

### Examples

**Search with detailed ratings:**
```bash
gplace search "coffee" --min-rating 4 --limit 5
```

**Fetch comprehensive details (including AI summaries and amenities):**
```bash
gplace details ChIJ2x4b1GmrQjQRukd8iZLakA8 --reviews
```

**JSON Output for scripting:**
```bash
gplace details ChIJ2x4b1GmrQjQRukd8iZLakA8 --json
```

## Library Usage

```go
client := gplace.NewClient(gplace.Options{
    APIKey: os.Getenv("GOOGLE_PLACES_API_KEY"),
})

details, err := client.DetailsWithOptions(ctx, gplace.DetailsRequest{
    PlaceID:        "ChIJ2x4b1GmrQjQRukd8iZLakA8",
    IncludeReviews: true,
})
```

## AI & Skills
This tool is designed to be "AI-friendly". If you are using an AI agent to manage your workspace, you can provide it with the `gplace` binary and ask:
> "Create a `SKILL.md` that uses the `gplace` command to find highly-rated vegetarian restaurants along my commute."

## Testing
```bash
make lint test
```