package gplace

import (
	"strconv"
	"strings"
)

func mapPriceRange(payload *priceRangePayload) *PriceRange {
	if payload == nil {
		return nil
	}
	return &PriceRange{
		StartPrice: mapMoney(payload.StartPrice),
		EndPrice:   mapMoney(payload.EndPrice),
	}
}

func mapMoney(payload *moneyPayload) *Money {
	if payload == nil {
		return nil
	}
	units, _ := strconv.ParseInt(payload.Units, 10, 64)
	return &Money{
		CurrencyCode: payload.CurrencyCode,
		Units:        units,
		Nanos:        payload.Nanos,
	}
}

func mapText(payload *localizedTextPayload) string {
	if payload == nil {
		return ""
	}
	return payload.Text
}

func mapGenerativeSummary(payload *generativeSummaryPayload) string {
	if payload == nil {
		return ""
	}
	return mapText(payload.Overview)
}

func mapReviewSummary(payload *reviewSummaryPayload) string {
	if payload == nil {
		return ""
	}
	return mapText(payload.Overview)
}

func mapReviews(reviews []reviewPayload) []Review {
	if len(reviews) == 0 {
		return nil
	}
	mapped := make([]Review, 0, len(reviews))
	for _, review := range reviews {
		mapped = append(mapped, Review{
			Name:                           review.Name,
			RelativePublishTimeDescription: review.RelativePublishTimeDescription,
			Text:                           mapLocalizedText(review.Text),
			OriginalText:                   mapLocalizedText(review.OriginalText),
			Rating:                         review.Rating,
			Author:                         mapAuthorAttribution(review.AuthorAttribution),
			PublishTime:                    review.PublishTime,
			FlagContentURI:                 review.FlagContentURI,
			GoogleMapsURI:                  review.GoogleMapsURI,
			VisitDate:                      mapVisitDate(review.VisitDate),
		})
	}
	return mapped
}

func mapLocalizedText(text *localizedTextPayload) *LocalizedText {
	if text == nil {
		return nil
	}
	// Avoid emitting empty text structs downstream.
	if strings.TrimSpace(text.Text) == "" && strings.TrimSpace(text.LanguageCode) == "" {
		return nil
	}
	return &LocalizedText{
		Text:         text.Text,
		LanguageCode: text.LanguageCode,
	}
}

func mapAuthorAttribution(author *authorAttributionPayload) *AuthorAttribution {
	if author == nil {
		return nil
	}
	// Drop empty attribution blocks to keep JSON clean.
	if strings.TrimSpace(author.DisplayName) == "" && strings.TrimSpace(author.URI) == "" && strings.TrimSpace(author.PhotoURI) == "" {
		return nil
	}
	return &AuthorAttribution{
		DisplayName: author.DisplayName,
		URI:         author.URI,
		PhotoURI:    author.PhotoURI,
	}
}

func mapAuthorAttributions(authors []authorAttributionPayload) []AuthorAttribution {
	if len(authors) == 0 {
		return nil
	}
	mapped := make([]AuthorAttribution, 0, len(authors))
	for _, author := range authors {
		mapped = append(mapped, AuthorAttribution(author))
	}
	return mapped
}

func mapVisitDate(date *visitDatePayload) *ReviewVisitDate {
	if date == nil {
		return nil
	}
	// Treat zeroed dates as missing.
	if date.Year == 0 && date.Month == 0 && date.Day == 0 {
		return nil
	}
	return &ReviewVisitDate{
		Year:  date.Year,
		Month: date.Month,
		Day:   date.Day,
	}
}

func mapLatLng(loc *location) *LatLng {
	if loc == nil {
		return nil
	}
	return &LatLng{Lat: loc.Latitude, Lng: loc.Longitude}
}

func displayName(name *displayNamePayload) string {
	if name == nil {
		return ""
	}
	return name.Text
}

func openNow(hours *openingHours) *bool {
	if hours == nil {
		return nil
	}
	return hours.OpenNow
}

func weekdayDescriptions(hours *openingHours) []string {
	if hours == nil {
		return nil
	}
	return hours.WeekdayDescriptions
}

func mapPriceLevel(value string) *int {
	if value == "" {
		return nil
	}
	if mapped, ok := enumToPriceLevel[value]; ok {
		return &mapped
	}
	return nil
}
