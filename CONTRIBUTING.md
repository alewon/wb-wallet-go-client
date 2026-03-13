# Contributing

## Development

1. Keep the project intentionally simple and explicit.
2. Do not add hidden validation or automatic retries unless they are clearly documented.
3. Prefer separate request/response structs over shared abstractions.
4. Keep the public API predictable and transparent.

## Local checks

```bash
gofmt -w client.go types.go
go test ./...
```

## Pull requests

- describe which WB Pay methods or structures were changed
- mention the exact documentation pages used
- avoid unrelated refactors in the same change
