go-build:	## Builds the API binary file using go
	go mod init github.com/holbertonschool-validation/tree/main/module2_task0/awesome-api
	go get github.com/gorilla/mux
	go build

hugo-build: check ## Build the website using gohugo on the dist folder
	hugo

build:	## Builds with go and hugo
	make go-build
	make hugo-build

check: ## Check markdown files for any syntax mistake
	markdownlint ./content/post/*.md
	markdown-link-check ./content/post/*.md

post: ## Create a new blog post
	hugo new posts/$(POST_NAME).md
## Update the title of the new post
	sed -i "s/title: .*/title: $(POST_TITLE)/" ./content/posts/$(POST_NAME).md

validate: ## W3C validator
	make build
	-git clone https://github.com/holbertonschool/W3C-Validator.git
	-./W3C-Validator/w3c_validator.py ./dist/index.html
	-rm -rf ./w3c_validator

run: ## Runs the program
	@./awesome-api >./awesome.log 2>&1 &

stop: ## Kills awesome-api's current process
	@kill "$(shell pgrep awesome-api)" 2>/dev/null &

clean: ## Delete the binary done with build command
	@make -s stop
	-rm -rf awesome-api awesome.log coverage* go* ./dist

test: ## Tests if the website is currently running properly
	-make build
	@make unit-tests
	@make integration-tests

lint: ## Helps to make a lint on the files
	@golangci-lint run

unit-tests: ## Runs the unit testy
	@go test -v -short -coverprofile=coverage-units.out

integration-tests: ## make unit tests on the files
	@go test -v -coverprofile=coverage-integrations.out

help: ## Show commands usage information
	@echo "Available targets:"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//' | awk '{ printf "  %-20s %s\n", $$1, $$2 }'
