
# Auto-discover ALL modules from go.work file (for lint, build, etc.)
MODULES := $(shell go work edit -json | jq -r '.Use[].DiskPath')

# Hard-coded list of PUBLIC packages that need documentation
# Add new public packages here as they are created
PUBLIC_MODULES := ./sprites

# Default port for documentation server
DOC_PORT ?= 8080

.PHONY: lint docs docs-check docs-server

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

# Check that all public symbols have documentation comments
docs-check:
	@echo "Checking documentation for public modules: $(PUBLIC_MODULES)"
	@for module in $(PUBLIC_MODULES); do \
		echo "  Checking $$module..."; \
		(cd $$module && go doc -all . > /dev/null 2>&1) || (echo "  Warning: Could not generate docs for $$module"; exit 1); \
	done
	@echo "Documentation check complete."

# Generate and display documentation summary for public modules
docs:
	@echo "Generating documentation for public modules: $(PUBLIC_MODULES)"
	@echo ""
	@for module in $(PUBLIC_MODULES); do \
		echo "=== $$module ==="; \
		(cd $$module && go doc -short .) || true; \
		echo ""; \
	done

# Start a local documentation server using pkgsite
# Install pkgsite with: go install golang.org/x/pkgsite/cmd/pkgsite@latest
# Note: pkgsite serves all packages in the tree, but the browser view will
# organize them appropriately
docs-server:
	@which pkgsite > /dev/null 2>&1 || (echo "pkgsite not installed. Run: go install golang.org/x/pkgsite/cmd/pkgsite@latest" && exit 1)
	@echo "Starting documentation server on http://localhost:$(DOC_PORT)"
	@echo "Public modules: $(PUBLIC_MODULES)"
	@echo "Press Ctrl+C to stop"
	@pkgsite -http=localhost:$(DOC_PORT) .

# Show help for documentation targets
docs-help:
	@echo "Documentation targets:"
	@echo "  make docs        - Generate and display docs for public modules only"
	@echo "  make docs-check  - Verify docs exist for public modules only"
	@echo "  make docs-server - Start local docs server (serves all, but public matter)"
	@echo "  make docs-help   - Show this help message"
	@echo ""
	@echo "Public modules (documented): $(PUBLIC_MODULES)"
	@echo "All modules (linted): $(MODULES)"
	@echo ""
	@echo "To view docs for a specific public package:"
	@echo "  go doc ./sprites           # Show package doc"
	@echo "  go doc ./sprites.Sprite      # Show Sprite type doc"
	@echo "  go doc ./sprites.NewSprite   # Show NewSprite function doc"
