package goplaces

// SearchRequest defines a text search with optional filters.
type SearchRequest struct {
	Query        string        `json:"query"`
	Filters      *Filters      `json:"filters,omitempty"`
	LocationBias *LocationBias `json:"location_bias,omitempty"`
	Limit        int           `json:"limit,omitempty"`
	PageToken    string        `json:"page_token,omitempty"`
}

// Filters are optional search refinements.
type Filters struct {
	Keyword     string   `json:"keyword,omitempty"`
	Types       []string `json:"types,omitempty"`
	OpenNow     *bool    `json:"open_now,omitempty"`
	MinRating   *float64 `json:"min_rating,omitempty"`
	PriceLevels []int    `json:"price_levels,omitempty"`
}

// LocationBias limits search results to a circular area.
type LocationBias struct {
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	RadiusM float64 `json:"radius_m"`
}

// LatLng holds geographic coordinates.
type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// SearchResponse contains a list of places and optional pagination token.
type SearchResponse struct {
	Results       []PlaceSummary `json:"results"`
	NextPageToken string         `json:"next_page_token,omitempty"`
}

// PlaceSummary is a compact view of a place.
type PlaceSummary struct {
	PlaceID    string   `json:"place_id"`
	Name       string   `json:"name,omitempty"`
	Address    string   `json:"address,omitempty"`
	Location   *LatLng  `json:"location,omitempty"`
	Rating     *float64 `json:"rating,omitempty"`
	PriceLevel *int     `json:"price_level,omitempty"`
	Types      []string `json:"types,omitempty"`
	OpenNow    *bool    `json:"open_now,omitempty"`
}

// PlaceDetails is a detailed view of a place.
type PlaceDetails struct {
	PlaceID    string   `json:"place_id"`
	Name       string   `json:"name,omitempty"`
	Address    string   `json:"address,omitempty"`
	Location   *LatLng  `json:"location,omitempty"`
	Rating     *float64 `json:"rating,omitempty"`
	PriceLevel *int     `json:"price_level,omitempty"`
	Types      []string `json:"types,omitempty"`
	Phone      string   `json:"phone,omitempty"`
	Website    string   `json:"website,omitempty"`
	Hours      []string `json:"hours,omitempty"`
	OpenNow    *bool    `json:"open_now,omitempty"`
}

// LocationResolveRequest resolves a text location into place candidates.
type LocationResolveRequest struct {
	LocationText string `json:"location_text"`
	Limit        int    `json:"limit,omitempty"`
}

// LocationResolveResponse contains resolved locations.
type LocationResolveResponse struct {
	Results []ResolvedLocation `json:"results"`
}

// ResolvedLocation is a place candidate for a location string.
type ResolvedLocation struct {
	PlaceID  string   `json:"place_id"`
	Name     string   `json:"name,omitempty"`
	Address  string   `json:"address,omitempty"`
	Location *LatLng  `json:"location,omitempty"`
	Types    []string `json:"types,omitempty"`
}
