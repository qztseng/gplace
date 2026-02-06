package gplace

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	detailsFieldMaskBase   = "id,displayName,formattedAddress,location,rating,userRatingCount,priceLevel,priceRange,types,primaryType,primaryTypeDisplayName,businessStatus,googleMapsUri,editorialSummary,generativeSummary,reviewSummary,regularOpeningHours,currentOpeningHours,nationalPhoneNumber,websiteUri,servesBeer,servesBreakfast,servesBrunch,servesCocktails,servesCoffee,servesDessert,servesDinner,servesLunch,servesVegetarianFood,servesWine"
	detailsFieldMaskReview = "reviews"
)

// Details fetches details for a specific place ID.
func (c *Client) Details(ctx context.Context, placeID string) (PlaceDetails, error) {
	return c.DetailsWithOptions(ctx, DetailsRequest{PlaceID: placeID})
}

// DetailsWithOptions fetches place details with locale hints.
func (c *Client) DetailsWithOptions(ctx context.Context, req DetailsRequest) (PlaceDetails, error) {
	placeID := strings.TrimSpace(req.PlaceID)
	if placeID == "" {
		return PlaceDetails{}, ValidationError{Field: "place_id", Message: "required"}
	}

	endpoint, err := c.buildURL("/places/"+placeID, map[string]string{
		"languageCode": strings.TrimSpace(req.Language),
		"regionCode":   strings.TrimSpace(req.Region),
	})
	if err != nil {
		return PlaceDetails{}, err
	}

	payload, err := c.doRequest(ctx, http.MethodGet, endpoint, nil, detailsFieldMaskForRequest(req))
	if err != nil {
		return PlaceDetails{}, err
	}

	var place placeItem
	if err := json.Unmarshal(payload, &place); err != nil {
		return PlaceDetails{}, fmt.Errorf("gplace: decode place details: %w", err)
	}

	return mapPlaceDetails(place), nil
}

func detailsFieldMaskForRequest(req DetailsRequest) string {
	fields := []string{detailsFieldMaskBase}
	if req.IncludeReviews {
		// Reviews are heavy; opt-in to include them.
		fields = append(fields, detailsFieldMaskReview)
	}
	return strings.Join(fields, ",")
}

func mapPlaceDetails(place placeItem) PlaceDetails {
	return PlaceDetails{
		PlaceID:                place.ID,
		Name:                   displayName(place.DisplayName),
		Address:                place.FormattedAddress,
		Location:               mapLatLng(place.Location),
		Rating:                 place.Rating,
		UserRatingCount:        place.UserRatingCount,
		PriceLevel:             mapPriceLevel(place.PriceLevel),
		PriceRange:             mapPriceRange(place.PriceRange),
		BusinessStatus:         place.BusinessStatus,
		GoogleMapsURI:          place.GoogleMapsURI,
		PrimaryType:            place.PrimaryType,
		PrimaryTypeDisplayName: mapText(place.PrimaryTypeDisplayName),
		EditorialSummary:       mapText(place.EditorialSummary),
		GenerativeSummary:      mapGenerativeSummary(place.GenerativeSummary),
		ReviewSummary:          mapReviewSummary(place.ReviewSummary),
		Types:                  place.Types,
		Phone:                  place.NationalPhoneNumber,
		Website:                place.WebsiteURI,
		Hours:                  weekdayDescriptions(place.RegularOpeningHours),
		OpenNow:                openNow(place.CurrentOpeningHours),
		Reviews:                mapReviews(place.Reviews),
		ServesBeer:             place.ServesBeer,
		ServesBreakfast:        place.ServesBreakfast,
		ServesBrunch:           place.ServesBrunch,
		ServesCocktails:        place.ServesCocktails,
		ServesCoffee:           place.ServesCoffee,
		ServesDessert:          place.ServesDessert,
		ServesDinner:           place.ServesDinner,
		ServesLunch:            place.ServesLunch,
		ServesVegetarianFood:   place.ServesVegetarianFood,
		ServesWine:             place.ServesWine,
	}
}
