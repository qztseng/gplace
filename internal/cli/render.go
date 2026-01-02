package cli

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/steipete/goplaces"
)

func renderSearch(color Color, response goplaces.SearchResponse) string {
	var out bytes.Buffer
	count := len(response.Results)
	if count == 0 {
		return "No results."
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

func renderDetails(color Color, place goplaces.PlaceDetails) string {
	var out bytes.Buffer
	out.WriteString(color.Bold(formatTitle(color, place.Name, place.Address)))
	out.WriteString("\n")
	writePlaceDetails(&out, color, place)
	return out.String()
}

func renderResolve(color Color, response goplaces.LocationResolveResponse) string {
	var out bytes.Buffer
	count := len(response.Results)
	if count == 0 {
		return "No results."
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

func writePlaceSummary(out *bytes.Buffer, color Color, place goplaces.PlaceSummary) {
	writeLine(out, color, "ID", place.PlaceID)
	writeLocation(out, color, place.Location)
	writeRating(out, color, place.Rating, place.PriceLevel)
	writeTypes(out, color, place.Types)
	writeOpenNow(out, color, place.OpenNow)
}

func writePlaceDetails(out *bytes.Buffer, color Color, place goplaces.PlaceDetails) {
	writeLine(out, color, "ID", place.PlaceID)
	writeLocation(out, color, place.Location)
	writeRating(out, color, place.Rating, place.PriceLevel)
	writeTypes(out, color, place.Types)
	writeOpenNow(out, color, place.OpenNow)
	writeLine(out, color, "Phone", place.Phone)
	writeLine(out, color, "Website", place.Website)
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

func writeResolvedLocation(out *bytes.Buffer, color Color, place goplaces.ResolvedLocation) {
	writeLine(out, color, "ID", place.PlaceID)
	writeLocation(out, color, place.Location)
	writeTypes(out, color, place.Types)
}

func writeLocation(out *bytes.Buffer, color Color, loc *goplaces.LatLng) {
	if loc == nil {
		return
	}
	writeLine(out, color, "Location", fmt.Sprintf("%.6f, %.6f", loc.Lat, loc.Lng))
}

func writeRating(out *bytes.Buffer, color Color, rating *float64, priceLevel *int) {
	if rating == nil && priceLevel == nil {
		return
	}
	parts := make([]string, 0, 2)
	if rating != nil {
		parts = append(parts, fmt.Sprintf("%.1f", *rating))
	}
	if priceLevel != nil {
		parts = append(parts, fmt.Sprintf("$%d", *priceLevel))
	}
	writeLine(out, color, "Rating", strings.Join(parts, " · "))
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
