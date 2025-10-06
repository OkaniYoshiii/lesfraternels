GOOSE := go run github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: goose-up
goose-up:
	$(GOOSE) up