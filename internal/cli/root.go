package cli

import (
	"time"
)

type Root struct {
	Global  GlobalOptions `embed:""`
	Search  SearchCmd     `cmd:"" help:"Search places by text query."`
	Details DetailsCmd    `cmd:"" help:"Fetch place details by place ID."`
	Resolve ResolveCmd    `cmd:"" help:"Resolve a location string to candidate places."`
}

type GlobalOptions struct {
	APIKey  string        `help:"Google Places API key." env:"GOOGLE_PLACES_API_KEY"`
	BaseURL string        `help:"Places API base URL." env:"GOOGLE_PLACES_BASE_URL" default:"https://places.googleapis.com/v1"`
	Timeout time.Duration `help:"HTTP timeout." default:"10s"`
	JSON    bool          `help:"Output JSON."`
	NoColor bool          `help:"Disable color output."`
	Verbose bool          `help:"Verbose logging."`
	Version VersionFlag   `name:"version" help:"Print version and exit."`
}

type SearchCmd struct {
	Query      string   `arg:"" name:"query" help:"Search text."`
	Limit      int      `help:"Max results (1-20)." default:"10"`
	PageToken  string   `help:"Page token for pagination."`
	Keyword    string   `help:"Keyword to append to the query."`
	Type       []string `help:"Place type filter (includedType). Repeatable."`
	OpenNow    *bool    `help:"Return only currently open places."`
	MinRating  *float64 `help:"Minimum rating (0-5)."`
	PriceLevel []int    `help:"Price levels 0-4. Repeatable."`
	Lat        *float64 `help:"Latitude for location bias."`
	Lng        *float64 `help:"Longitude for location bias."`
	RadiusM    *float64 `help:"Radius in meters for location bias."`
}

type DetailsCmd struct {
	PlaceID string `arg:"" name:"place_id" help:"Place ID."`
}

type ResolveCmd struct {
	LocationText string `arg:"" name:"location" help:"Location text to resolve."`
	Limit        int    `help:"Max results (1-10)." default:"5"`
}
