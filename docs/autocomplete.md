# Autocomplete

Autocomplete returns place + query suggestions for partial text.

## CLI

```bash
gplace autocomplete "cof" \
  --session-token "gplace-demo" \
  --limit 5 \
  --language en \
  --region US
```

Optional location bias:

```bash
gplace autocomplete "pizza" --lat 40.7411 --lng -73.9897 --radius-m 1500
```

## Library

```go
response, err := client.Autocomplete(ctx, gplace.AutocompleteRequest{
    Input:        "cof",
    SessionToken: "gplace-demo",
    Limit:        5,
    Language:     "en",
    Region:       "US",
})
```

## Notes

- Use a session token for billing consistency across autocomplete + details.
- Limit is applied client-side after the API response.
