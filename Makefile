
# Auto-discover all modules from go.work file using jq
MODULES := $(shell go work edit -json | jq -r '.Use[].DiskPath')

.PHONY: lint

lint:
	@for module in $(MODULES); do \
		echo "Linting $$module..."; \
		(cd $$module && go vet ./...) || exit 1; \
	done
	@which staticcheck > /dev/null 2>&1 || (echo "staticcheck not installed. Run: go install honnef.co/go/tools/cmd/staticcheck@latest" && exit 1)
	@for module in $(MODULES); do \
		echo "Running staticcheck on $$module..."; \
		(cd $$module && staticcheck ./...) || exit 1; \
	done
