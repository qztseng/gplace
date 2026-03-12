# SKILL.md: gplace Automation for AI Agents

This file provides instructions and examples for AI agents to automate tasks using the `gplace` CLI.

## Objective
Use `gplace` to find, resolve, and extract detailed metadata about places for research, planning, or data collection.

## Command Reference

### 1. Search for Places
Find places using a text query with optional filters.
```bash
gplace search "query" [--type TYPE] [--limit N] [--min-rating R] [--price-level P] [--open-now] [--local] [--json]
```

### 2. Get Place Details
Fetch comprehensive metadata including AI summaries, price ranges, and amenities.
```bash
gplace details PLACE_ID [--reviews] [--language BCP-47] [--local] [--json]
```

### 3. Route-based Search
Search for places along a driving or walking route.
```bash
gplace route "query" --from "start" --to "destination" [--mode DRIVE|WALK|BICYCLE] [--json]
```

## Review Extraction & Localization Workflow

To provide the most relevant and high-quality reviews, you SHOULD follow this two-step process:

1.  **Step 1: Identify the Place & Location**
    Perform a `gplace search` or `gplace nearby` to find the target establishment. Note the country/region from the address components or the returned address string.
2.  **Step 2: Fetch Localized Reviews**
    Call `gplace details <PLACE_ID> --reviews --language <CODE>` using the local language of the place (e.g., `ja` for Japan, `fr` for France).
    *   **Why?** Google's "Most Relevant" algorithm is language-dependent. Fetching reviews in the local language provides deeper, more authentic insights.
    *   **Shortcut**: If you are unsure of the language code, use the `--local` flag to have the tool auto-detect the primary local language.

## Place Types: (Food & Drink types must follow the types listed below:
Asian: japanese_restaurant, sushi_restaurant, ramen_restaurant, chinese_restaurant, korean_restaurant, thai_restaurant, indian_restaurant, vietnamese_restaurant, dim_sum_restaurant, noodle_shop, afghani_restaurant, bangladeshi_restaurant, burmese_restaurant, cambodian_restaurant, filipino_restaurant, indonesian_restaurant, malaysian_restaurant, pakistani_restaurant, sri_lankan_restaurant, taiwanese_restaurant, tibetan_restaurant

Western/International: american_restaurant, italian_restaurant, french_restaurant, mexican_restaurant, spanish_restaurant, greek_restaurant, mediterranean_restaurant, steak_house, pizza_restaurant, hamburger_restaurant, african_restaurant, argentinian_restaurant, australian_restaurant, austrian_restaurant, belgian_restaurant, brazilian_restaurant, british_restaurant, caribbean_restaurant, chilean_restaurant, colombian_restaurant, cuban_restaurant, czech_restaurant, danish_restaurant, dutch_restaurant, european_restaurant, german_restaurant, hungarian_restaurant, irish_restaurant, israeli_restaurant, lebanese_restaurant, moroccan_restaurant, peruvian_restaurant, polish_restaurant, portuguese_restaurant, romanian_restaurant, russian_restaurant, scandinavian_restaurant, swiss_restaurant, turkish_restaurant, ukrainian_restaurant

Cafes & Desserts: cafe, coffee_shop, bakery, dessert_restaurant, ice_cream_shop, tea_house, juice_shop, donut_shop, pastry_shop, acai_shop, bagel_shop, cake_shop, chocolate_shop, confectionery

Bar & Nightlife: bar, pub, wine_bar, beer_garden, brewery, cocktail_bar, izakaya_restaurant, brewpub, hookah_bar, lounge_bar, sports_bar

Specialty/Style: bistro, brunch_restaurant, buffet_restaurant, diner, fine_dining_restaurant, food_court, gastropub, halal_restaurant, vegan_restaurant, vegetarian_restaurant, seafood_restaurant, barbecue_restaurant, steak_house


To improve search accuracy, you MUST map user requests to the most specific official Google Place Type using the `--type` flag.

## Output Format Requirements (Mandatory)
When presenting a place to the user, you MUST synthesize a comprehensive overview using all available data:

1.  **Credibility & Pricing**:
    *   Display Rating score and total count (e.g., `4.5 ⭐ (1,240 reviews)`).
    *   Display Price Level as symbols (e.g., `$$`).
    *   Explicitly state the `price_range` if available in the JSON.
2.  **The "Review Summary" synthesis**:
    *   You MUST use the `reviewSummary` field (if available) as the primary source for the establishment's vibe.
    *   You MUST analyze the individual `reviews` (usually 5 are returned) to find specific details not in the summary.
3.  **Recommended Dishes & Specialties**:
    *   For any food or drink establishment, you MUST extract a list of recommended dishes, signature drinks/foods, or specialties mentioned across both the `reviewSummary` and the individual `reviews`.
4.  **Bullet-Point Summary**:
    *   Provide 3-4 bullet points covering the pros, cons, and unique features (e.g., "Great for groups," "Hidden gem," "Reservations required").
5.  **Actionable Info**:
    *   Include Address, current "Open/Closed" status, phone number, and the Google Maps URI.

**Full Data Fetch**: If initial search output lacks details, you MUST call `gplace details <PLACE_ID> --reviews --local` before finalizing your response to ensure you have the `review_summary` and individual reviews for synthesis.
