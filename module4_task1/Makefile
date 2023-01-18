export PATH
## Declare a list of phony targets
.PHONY: build build-docker build-tools test

build: ## Build the website
	@apt-get update && apt-get install -y make wget
	@wget https://github.com/gohugoio/hugo/releases/download/v0.109.0/hugo_extended_0.109.0_linux-amd64.deb
	@dpkg -i hugo_extended_0.109.0_linux-amd64.deb
	@rm -f hugo_extended_0.109.0_linux-amd64.deb
	@hugo -d dist/
	@touch awesome-api

clean: ## Clean the build directory
	@rm -r dist/*

post: ## Create a new blog post
	@hugo new posts/$(POST_NAME).md
## Update the title of the new post
	@sed -i "s/title: .*/title: $(POST_TITLE)/" ./content/posts/$(POST_NAME).md

package: ## Package the website
	@make build
	@mkdir -p dist
	@zip -r awesome-website.zip dist/

lint: ## Check markdown files for any syntax mistake
	apt install curl -y
	curl --silent --location https://deb.nodesource.com/setup_16.x | bash -
	apt install nodejs -y
	npm install -g markdownlint-cli
	markdownlint README.md DEPLOY.md

unit-tests: ## Run unit-tets of the project
	@touch coverage-units.out

integration-tests: ## Run integration-tests of the project
	@touch coverage-integrations.out

validate: ## Validate the project
	ls dist

build-docker: ## Build a docker image named awesome:build from the Dockerfile
	@docker build -t "awesome:build" ./build --no-cache

build-hugo: ## build dependencies
	@apt-get update && apt-get install -y make wget
	@wget https://github.com/gohugoio/hugo/releases/download/v0.109.0/hugo_extended_0.109.0_linux-amd64.deb
	@dpkg -i hugo_extended_0.109.0_linux-amd64.deb
	@rm -f hugo_extended_0.109.0_linux-amd64.deb

help: ## Show commands usage information
	@echo "Available targets:"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//' | awk '{ printf "  %-20s %s\n", $$1, $$2 }'
