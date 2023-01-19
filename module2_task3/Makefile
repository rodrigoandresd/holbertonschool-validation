build: ## compile the project source code into a binary named awesome-api
	@make lint && go build || exit -1

run: ## Runs the program
	@./awesome-api >./awesome.log 2>&1 &

stop: ## Kills awesome-api's current process
	@kill "$(shell pgrep awesome-api)" 2>/dev/null &

clean: ## Delete the binary done with build command
	@make -s stop
	@rm -r awesome.log 2>&1 &
	@rm -r awesome-api 2>&1 &
	@rm -r coverage-integrations.out
	@rm -r coverage-units.out

test: ## Tests if the website is currently running properly
	@make unit-tests &
	@make integration-tests

help: ## Show this
	@sed -ne '/@sed/!s/## //p' Makefile

lint: ## Helps to make a lint on the files
	@golangci-lint run

unit-tests: ## Runs the unit testy
	@go test -v -short -coverprofile=coverage-units.out || echo $1

integration-tests: ## make unit tests on the files
	@go test -v -coverprofile=coverage-integrations.out || echo $1
