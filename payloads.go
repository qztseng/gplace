package gplace

type searchResponse struct {
	Places        []placeItem `json:"places"`
	NextPageToken string      `json:"nextPageToken"`
}

type placeItem struct {
	ID                     string                    `json:"id"`
	DisplayName            *displayNamePayload       `json:"displayName,omitempty"`
	FormattedAddress       string                    `json:"formattedAddress,omitempty"`
	Location               *location                 `json:"location,omitempty"`
	Rating                 *float64                  `json:"rating,omitempty"`
	UserRatingCount        *int                      `json:"userRatingCount,omitempty"`
	PriceLevel             string                    `json:"priceLevel,omitempty"`
	PriceRange             *priceRangePayload        `json:"priceRange,omitempty"`
	Types                  []string                  `json:"types,omitempty"`
	PrimaryType            string                    `json:"primaryType,omitempty"`
	PrimaryTypeDisplayName *localizedTextPayload     `json:"primaryTypeDisplayName,omitempty"`
	BusinessStatus         string                    `json:"businessStatus,omitempty"`
	GoogleMapsURI          string                    `json:"googleMapsUri,omitempty"`
	EditorialSummary       *localizedTextPayload     `json:"editorialSummary,omitempty"`
	GenerativeSummary      *generativeSummaryPayload `json:"generativeSummary,omitempty"`
	ReviewSummary          *reviewSummaryPayload     `json:"reviewSummary,omitempty"`
	CurrentOpeningHours    *openingHours             `json:"currentOpeningHours,omitempty"`
	RegularOpeningHours    *openingHours             `json:"regularOpeningHours,omitempty"`
	NationalPhoneNumber    string                    `json:"nationalPhoneNumber,omitempty"`
	WebsiteURI             string                    `json:"websiteUri,omitempty"`
	Reviews                []reviewPayload           `json:"reviews,omitempty"`
	ServesBeer             *bool                     `json:"servesBeer,omitempty"`
	ServesBreakfast        *bool                     `json:"servesBreakfast,omitempty"`
	ServesBrunch           *bool                     `json:"servesBrunch,omitempty"`
	ServesCocktails        *bool                     `json:"servesCocktails,omitempty"`
	ServesCoffee           *bool                     `json:"servesCoffee,omitempty"`
	ServesDessert          *bool                     `json:"servesDessert,omitempty"`
	ServesDinner           *bool                     `json:"servesDinner,omitempty"`
	ServesLunch            *bool                     `json:"servesLunch,omitempty"`
	ServesVegetarianFood   *bool                     `json:"servesVegetarianFood,omitempty"`
	ServesWine             *bool                     `json:"servesWine,omitempty"`
}

type generativeSummaryPayload struct {
	Overview *localizedTextPayload `json:"overview,omitempty"`
}

type reviewSummaryPayload struct {
	Overview *localizedTextPayload `json:"overview,omitempty"`
}

type priceRangePayload struct {
	StartPrice *moneyPayload `json:"startPrice,omitempty"`
	EndPrice   *moneyPayload `json:"endPrice,omitempty"`
}

type moneyPayload struct {
	CurrencyCode string `json:"currencyCode"`
	Units        string `json:"units"` // int64 as string in JSON
	Nanos        int32  `json:"nanos"`
}

type displayNamePayload struct {
	Text string `json:"text"`
}

type location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type openingHours struct {
	OpenNow             *bool    `json:"openNow,omitempty"`
	WeekdayDescriptions []string `json:"weekdayDescriptions,omitempty"`
}

type reviewPayload struct {
	Name                           string                    `json:"name,omitempty"`
	RelativePublishTimeDescription string                    `json:"relativePublishTimeDescription,omitempty"`
	Text                           *localizedTextPayload     `json:"text,omitempty"`
	OriginalText                   *localizedTextPayload     `json:"originalText,omitempty"`
	Rating                         *float64                  `json:"rating,omitempty"`
	AuthorAttribution              *authorAttributionPayload `json:"authorAttribution,omitempty"`
	PublishTime                    string                    `json:"publishTime,omitempty"`
	FlagContentURI                 string                    `json:"flagContentUri,omitempty"`
	GoogleMapsURI                  string                    `json:"googleMapsUri,omitempty"`
	VisitDate                      *visitDatePayload         `json:"visitDate,omitempty"`
}

type localizedTextPayload struct {
	Text         string `json:"text,omitempty"`
	LanguageCode string `json:"languageCode,omitempty"`
}

type authorAttributionPayload struct {
	DisplayName string `json:"displayName,omitempty"`
	URI         string `json:"uri,omitempty"`
	PhotoURI    string `json:"photoUri,omitempty"`
}

type visitDatePayload struct {
	Year  int `json:"year,omitempty"`
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
}
