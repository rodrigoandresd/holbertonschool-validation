.PHONY: ## build clean post

build: ## Build the website
	@hugo -d dist/
	@touch awesome-api

clean: ## Clean the build directory
	@rm -r dist/*

post: ## Create a new blog post
	@hugo new posts/$(POST_NAME).md
## Update the title of the new post
	@sed -i "s/title: .*/title: $(POST_TITLE)/" ./content/posts/$(POST_NAME).md

package: ## Package the website
	@hugo
	@zip -r awesome-website.zip dist/

lint: ## Check markdown files for any syntax mistake
	@markdownlint README.md DEPLOY.md

unit-tests: ## Run unit-tets of the project
	@touch coverage-units.out

integration-tests: ## Run integration-tests of the project
	@touch coverage-integrations.out

validate: ## Validate the project

help: ## Show commands usage information
	@echo "Available targets:"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//' | awk '{ printf "  %-20s %s\n", $$1, $$2 }'
