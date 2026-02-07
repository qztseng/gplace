# SKILL.md: gplace Automation for AI Agents

This file provides instructions and examples for AI agents to automate tasks using the `gplace` CLI.

## Objective
Use `gplace` to find, resolve, and extract detailed metadata about places for research, planning, or data collection.

## Command Reference

### 1. Search for Places
Find places using a text query with optional filters.
```bash
gplace search "query" [--limit N] [--min-rating R] [--price-level P] [--open-now] [--json]
```

### 2. Get Place Details
Fetch comprehensive metadata including AI summaries, price ranges, and amenities.
```bash
gplace details PLACE_ID [--reviews] [--json]
```

### 3. Route-based Search
Search for places along a driving or walking route.
```bash
gplace route "query" --from "start" --to "destination" [--mode DRIVE|WALK|BICYCLE] [--json]
```

## Workflow Examples

### Workflow: Find Top-Rated Quiet Coffee Shops
1. **Search**: Find coffee shops in a specific area.
   ```bash
   gplace search "quiet coffee shops in Seattle" --min-rating 4.5 --limit 10 --json
   ```
2. **Filter**: Identify candidate Place IDs from the JSON output.
3. **Inspect**: Get details for the best candidates to check for "serves coffee" and "review summaries".
   ```bash
   gplace details <PLACE_ID> --json
   ```

### Workflow: Planning a Commute Stop
Find a highly-rated vegetarian restaurant along a commute from Ballard to Downtown Seattle.
```bash
gplace route "vegetarian restaurant" --from "Ballard, Seattle" --to "Downtown Seattle" --mode DRIVE --limit 5 --json
```

## Tips for Agents
- Always use `--json` for predictable parsing.
- Use `details <ID> --reviews` if you need to perform sentiment analysis or extract specific user feedback.
- Numeric `price_level` is returned in JSON; use this for budget-based filtering.
