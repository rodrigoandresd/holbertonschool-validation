CHECK_FAILED=0
build: check ## Build the website
	hugo -d dist/
	if [ $(CHECK_FAILED) -eq 1 ]; then \
		echo "FAIL"; exit 1; \

clean: ## Clean the build directory
	rm -r dist/*

post: ## Create a new blog post
	hugo new posts/$(POST_NAME).md
## Update the title of the new post
	sed -i "s/title: .*/title: $(POST_TITLE)/" ./content/posts/$(POST_NAME).md

check: ## Check markdown files for any syntax mistake
	markdownlint ./post/*.md && markdown-link-check ./post/*.md || true

validate: ## W3C validator
	make build
	./w3c_validator.py ./dist/index.html || true

help: ## Show commands usage information
	@echo "Available targets:"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//' | awk '{ printf "  %-20s %s\n", $$1, $$2 }'
