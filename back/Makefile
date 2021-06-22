.PHONY: swag
swag:
	rm -r docs || true
	../bin/swag i --parseDependency --parseInternal -g ./cmd/main.go
