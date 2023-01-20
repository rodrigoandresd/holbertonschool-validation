# Building a Go-Hugo Website

This project uses the Go programming language and the Hugo static site generator to build a website. In order to build the website, you will need to have the following tools installed on your system:

## Prerequisites:
    GNU Make in version 3.81+
    Go Hugo v0.80+


## Lifecycle:

## check

    make check

check markdown files for any syntax mistake

## buid

To build the website, run the following command from the root directory of the project:

    make build

This will generate the website from the source files and configuration files, and store the generated files in the "dist/" directory.

## Cleaning the Build Directory

To clean up the build directory, run the following command:

    make clean

This will delete all files in the "dist/" directory.

## Validate

    make validate

W3C validator

## Creating a New Blog Post
To create a new blog post, run the following command:

    make post POST_NAME=<name of post> POST_TITLE=<title of post>

This will create a new blog post with the specified name and title in the "posts/" directory.

## Help
For more information on the available targets and their usage, run the following command:

    make help

This will display a list of available targets and a brief description of each one.

run: Runs the built binary and send everything to awesome.log stop: Terminates the execution of awesome-api test: Tests API using go test lint: Lints all the go files using golangci-lint unit-tests: Runs implemented unit test using go test integration-tests: Runs integration test using go test

## Code Comments

The source code for this project includes comments to provide additional context and information for readers. Please refer to these comments for more information on the inner workings of the code.
