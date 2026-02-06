# Photos

Photos support includes metadata on Place Details and a helper to fetch a photo URL.

## Details (metadata)

```bash
gplace details "PLACE_ID" --photos
```

## Photo URL

```bash
gplace photo "places/PLACE_ID/photos/PHOTO_ID" --max-width 1200
```

## Library

```go
details, err := client.DetailsWithOptions(ctx, gplace.DetailsRequest{
    PlaceID:       "PLACE_ID",
    IncludePhotos: true,
})

photo, err := client.PhotoMedia(ctx, gplace.PhotoMediaRequest{
    Name:       details.Photos[0].Name,
    MaxWidthPx: 1200,
})
```

## Notes

- Photo media always returns a URL (skip redirect) for easy downloading.
- Use `max-width`/`max-height` to control the asset size.
