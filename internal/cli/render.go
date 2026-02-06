package cli

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/qztseng/gplace"
)

func renderSearch(color Color, response gplace.SearchResponse) string {
	var out bytes.Buffer
	count := len(response.Results)
	if count == 0 {
		return emptyResultsMessage
	}
	out.WriteString(color.Bold(fmt.Sprintf("Results (%d)", count)))
	out.WriteString("\n")

	for i, place := range response.Results {
		out.WriteString(fmt.Sprintf("%d. %s\n", i+1, formatTitle(color, place.Name, place.Address)))
		writePlaceSummary(&out, color, place)
		if i < count-1 {
			out.WriteString("\n")
		}
	}

	if strings.TrimSpace(response.NextPageToken) != "" {
		out.WriteString("\n")
		out.WriteString(color.Dim("Next page token:"))
		out.WriteString(" ")
		out.WriteString(response.NextPageToken)
	}

	return out.String()
}

func renderAutocomplete(color Color, response gplace.AutocompleteResponse) string {
	var out bytes.Buffer
	count := len(response.Suggestions)
	if count == 0 {
		return emptyResultsMessage
	}
	out.WriteString(color.Bold(fmt.Sprintf("Suggestions (%d)", count)))
	out.WriteString("\n")

	for i, suggestion := range response.Suggestions {
		title := formatTitle(color, autocompleteTitle(suggestion), autocompleteSubtitle(suggestion))
		out.WriteString(fmt.Sprintf("%d. %s\n", i+1, title))
		writeAutocompleteSuggestion(&out, color, suggestion)
		if i < count-1 {
			out.WriteString("\n")
		}
	}
	return out.String()
}

func renderNearby(color Color, response gplace.NearbySearchResponse) string {
	var out bytes.Buffer
	count := len(response.Results)
	if count == 0 {
		return emptyResultsMessage
	}
	out.WriteString(color.Bold(fmt.Sprintf("Nearby (%d)", count)))
	out.WriteString("\n")

	for i, place := range response.Results {
		out.WriteString(fmt.Sprintf("%d. %s\n", i+1, formatTitle(color, place.Name, place.Address)))
		writePlaceSummary(&out, color, place)
		if i < count-1 {
			out.WriteString("\n")
		}
	}

	if strings.TrimSpace(response.NextPageToken) != "" {
		out.WriteString("\n")
		out.WriteString(color.Dim("Next page token:"))
		out.WriteString(" ")
		out.WriteString(response.NextPageToken)
	}

	return out.String()
}

func renderDetails(color Color, place gplace.PlaceDetails) string {
	var out bytes.Buffer
	out.WriteString(color.Bold(formatTitle(color, place.Name, place.Address)))
	out.WriteString("\n")
	writePlaceDetails(&out, color, place)
	return out.String()
}

func renderResolve(color Color, response gplace.LocationResolveResponse) string {
	var out bytes.Buffer
	count := len(response.Results)
	if count == 0 {
		return emptyResultsMessage
	}
	out.WriteString(color.Bold(fmt.Sprintf("Resolved (%d)", count)))
	out.WriteString("\n")

	for i, place := range response.Results {
		out.WriteString(fmt.Sprintf("%d. %s\n", i+1, formatTitle(color, place.Name, place.Address)))
		writeResolvedLocation(&out, color, place)
		if i < count-1 {
			out.WriteString("\n")
		}
	}
	return out.String()
}

func renderRoute(color Color, response gplace.RouteResponse) string {
	var out bytes.Buffer
	count := len(response.Waypoints)
	if count == 0 {
		return emptyResultsMessage
	}
	out.WriteString(color.Bold(fmt.Sprintf("Route waypoints (%d)", count)))
	out.WriteString("\n")

	for i, waypoint := range response.Waypoints {
		out.WriteString(color.Bold(fmt.Sprintf("Waypoint %d", i+1)))
		out.WriteString(" ")
		out.WriteString(color.Dim(fmt.Sprintf("(%.6f, %.6f)", waypoint.Location.Lat, waypoint.Location.Lng)))
		out.WriteString("\n")

		if len(waypoint.Results) == 0 {
			out.WriteString(emptyResultsMessage)
			out.WriteString("\n")
		} else {
			for j, place := range waypoint.Results {
				out.WriteString(fmt.Sprintf("%d. %s\n", j+1, formatTitle(color, place.Name, place.Address)))
				writePlaceSummary(&out, color, place)
				if j < len(waypoint.Results)-1 {
					out.WriteString("\n")
				}
			}
		}

		if i < count-1 {
			out.WriteString("\n")
		}
	}

	return out.String()
}

func formatTitle(color Color, name string, address string) string {
	display := strings.TrimSpace(name)
	if display == "" {
		display = "(no name)"
	}
	if address == "" {
		return color.Cyan(display)
	}
	return color.Cyan(display) + " — " + address
}

const emptyResultsMessage = "No results."

func autocompleteTitle(suggestion gplace.AutocompleteSuggestion) string {
	if strings.TrimSpace(suggestion.MainText) != "" {
		return suggestion.MainText
	}
	return suggestion.Text
}

func autocompleteSubtitle(suggestion gplace.AutocompleteSuggestion) string {
	if strings.TrimSpace(suggestion.SecondaryText) != "" {
		return suggestion.SecondaryText
	}
	if strings.TrimSpace(suggestion.Text) == "" || strings.TrimSpace(suggestion.MainText) == "" {
		return ""
	}
	return suggestion.Text
}

func writePlaceSummary(out *bytes.Buffer, color Color, place gplace.PlaceSummary) {
	writeLine(out, color, "ID", place.PlaceID)
	writeLocation(out, color, place.Location)
	writeRating(out, color, place.Rating, place.UserRatingCount, place.PriceLevel, nil)
	writeTypes(out, color, place.Types)
	writeOpenNow(out, color, place.OpenNow)
}

func writeAutocompleteSuggestion(out *bytes.Buffer, color Color, suggestion gplace.AutocompleteSuggestion) {
	writeLine(out, color, "Kind", suggestion.Kind)
	writeLine(out, color, "ID", suggestion.PlaceID)
	writeLine(out, color, "Place", suggestion.Place)
	writeTypes(out, color, suggestion.Types)
	if suggestion.DistanceMeters != nil {
		writeLine(out, color, "Distance", fmt.Sprintf("%dm", *suggestion.DistanceMeters))
	}
}

func writePlaceDetails(out *bytes.Buffer, color Color, place gplace.PlaceDetails) {
	writeLine(out, color, "ID", place.PlaceID)
	writeLocation(out, color, place.Location)
	writeRating(out, color, place.Rating, place.UserRatingCount, place.PriceLevel, place.PriceRange)
	writeLine(out, color, "Status", place.BusinessStatus)
	writeTypes(out, color, place.Types)

	if place.PrimaryTypeDisplayName != "" {
		writeLine(out, color, "Primary Type", place.PrimaryTypeDisplayName)
	}

	writeLine(out, color, "Phone", place.Phone)
	writeLine(out, color, "Website", place.Website)
	writeLine(out, color, "Maps", place.GoogleMapsURI)

	if place.EditorialSummary != "" {
		out.WriteString(color.Dim("Summary:"))
		out.WriteString(" ")
		out.WriteString(place.EditorialSummary)
		out.WriteString("\n")
	}

	if place.GenerativeSummary != "" {
		out.WriteString(color.Dim("AI Overview:"))
		out.WriteString(" ")
		out.WriteString(place.GenerativeSummary)
		out.WriteString("\n")
	}

	if place.ReviewSummary != "" {
		out.WriteString(color.Dim("Review Summary:"))
		out.WriteString(" ")
		out.WriteString(place.ReviewSummary)
		out.WriteString("\n")
	}

	writeAmenities(out, color, place)
	writeOpenNow(out, color, place.OpenNow)
	writeReviews(out, color, place.Reviews)

	if len(place.Hours) > 0 {
		out.WriteString(color.Dim("Hours:"))
		out.WriteString("\n")
		for _, entry := range place.Hours {
			out.WriteString("  - ")
			out.WriteString(entry)
			out.WriteString("\n")
		}
	}
}

func writeAmenities(out *bytes.Buffer, color Color, place gplace.PlaceDetails) {
	var amenities []string
	add := func(val *bool, label string) {
		if val != nil && *val {
			amenities = append(amenities, label)
		}
	}

	add(place.ServesBeer, "Beer")
	add(place.ServesBreakfast, "Breakfast")
	add(place.ServesBrunch, "Brunch")
	add(place.ServesCocktails, "Cocktails")
	add(place.ServesCoffee, "Coffee")
	add(place.ServesDessert, "Dessert")
	add(place.ServesDinner, "Dinner")
	add(place.ServesLunch, "Lunch")
	add(place.ServesVegetarianFood, "Vegetarian")
	add(place.ServesWine, "Wine")

	if len(amenities) > 0 {
		writeLine(out, color, "Serves", strings.Join(amenities, ", "))
	}
}

func writeResolvedLocation(out *bytes.Buffer, color Color, place gplace.ResolvedLocation) {
	writeLine(out, color, "ID", place.PlaceID)
	writeLocation(out, color, place.Location)
	writeTypes(out, color, place.Types)
}

func writeReviews(out *bytes.Buffer, color Color, reviews []gplace.Review) {
	if len(reviews) == 0 {
		return
	}
	out.WriteString(color.Dim("Reviews:"))
	out.WriteString("\n")

	// Keep CLI output compact by default.
	const maxReviews = 3
	count := len(reviews)
	limit := count
	if count > maxReviews {
		limit = maxReviews
	}

	for i := 0; i < limit; i++ {
		review := reviews[i]
		line := reviewLine(review)
		if line == "" {
			continue
		}
		out.WriteString("  - ")
		out.WriteString(line)
		out.WriteString("\n")
	}

	if count > maxReviews {
		out.WriteString(color.Dim(fmt.Sprintf("  ... %d more", count-maxReviews)))
		out.WriteString("\n")
	}
}

func writeLocation(out *bytes.Buffer, color Color, loc *gplace.LatLng) {
	if loc == nil {
		return
	}
	writeLine(out, color, "Location", fmt.Sprintf("%.6f, %.6f", loc.Lat, loc.Lng))
}

func writeRating(out *bytes.Buffer, color Color, rating *float64, count *int, priceLevel *int, priceRange *gplace.PriceRange) {
	if rating == nil && count == nil && priceLevel == nil && priceRange == nil {
		return
	}
	parts := make([]string, 0, 4)
	if rating != nil {
		r := fmt.Sprintf("%.1f", *rating)
		if count != nil {
			r += fmt.Sprintf(" (%d)", *count)
		}
		parts = append(parts, r)
	} else if count != nil {
		parts = append(parts, fmt.Sprintf("%d ratings", *count))
	}

	if priceRange != nil {
		parts = append(parts, formatPriceRange(priceRange))
	} else if priceLevel != nil {
		parts = append(parts, strings.Repeat("$", *priceLevel))
	}
	writeLine(out, color, "Rating", strings.Join(parts, " · "))
}

func formatPriceRange(pr *gplace.PriceRange) string {
	if pr == nil {
		return ""
	}
	start := formatMoney(pr.StartPrice)
	end := formatMoney(pr.EndPrice)
	if start == "" && end == "" {
		return ""
	}
	if start == end {
		return start
	}
	return fmt.Sprintf("%s–%s", start, end)
}

func formatMoney(m *gplace.Money) string {
	if m == nil {
		return ""
	}
	// Simple formatting: ignore nanos if they are 0 for brevity.
	// We use the currency code as a suffix/prefix.
	val := float64(m.Units) + float64(m.Nanos)/math.Pow(10, 9)

	// Format based on common currencies or just Code + Value
	symbol := m.CurrencyCode
	switch m.CurrencyCode {
	case "USD":
		symbol = "$"
	case "EUR":
		symbol = "€"
	case "GBP":
		symbol = "£"
	}

	if m.Nanos == 0 {
		return fmt.Sprintf("%s%.0f", symbol, val)
	}
	return fmt.Sprintf("%s%.2f", symbol, val)
}

func writeTypes(out *bytes.Buffer, color Color, types []string) {
	if len(types) == 0 {
		return
	}
	unique := uniqueStrings(types)
	writeLine(out, color, "Types", strings.Join(unique, ", "))
}

func writeOpenNow(out *bytes.Buffer, color Color, openNow *bool) {
	if openNow == nil {
		return
	}
	value := "no"
	if *openNow {
		value = "yes"
	}
	writeLine(out, color, "Open now", value)
}

func writeLine(out *bytes.Buffer, color Color, label string, value string) {
	if strings.TrimSpace(value) == "" {
		return
	}
	out.WriteString(color.Dim(label + ":"))
	out.WriteString(" ")
	out.WriteString(value)
	out.WriteString("\n")
}

func reviewLine(review gplace.Review) string {
	parts := make([]string, 0, 3)
	if review.Rating != nil {
		parts = append(parts, fmt.Sprintf("%.1f stars", *review.Rating))
	}
	if review.Author != nil && strings.TrimSpace(review.Author.DisplayName) != "" {
		parts = append(parts, "by "+review.Author.DisplayName)
	}
	if strings.TrimSpace(review.RelativePublishTimeDescription) != "" {
		parts = append(parts, "("+review.RelativePublishTimeDescription+")")
	}
	text := reviewText(review)
	if text != "" {
		parts = append(parts, text)
	}
	return strings.Join(parts, " ")
}

func reviewText(review gplace.Review) string {
	text := ""
	if review.Text != nil {
		text = review.Text.Text
	}
	// Fall back to original text when translation is empty.
	if strings.TrimSpace(text) == "" && review.OriginalText != nil {
		text = review.OriginalText.Text
	}
	return truncateText(strings.TrimSpace(text), 200)
}

func truncateText(value string, maxLen int) string {
	if maxLen <= 0 || value == "" {
		return value
	}
	if len(value) <= maxLen {
		return value
	}
	// Byte-based truncation is OK here because we only display previews.
	return strings.TrimSpace(value[:maxLen]) + "..."
}

func uniqueStrings(values []string) []string {
	seen := make(map[string]struct{}, len(values))
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}
