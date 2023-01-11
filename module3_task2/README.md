# Building a Go-Hugo Website

This project uses the Go programming language and the Hugo static site generator to build a website. In order to build the website, you will need to have the following tools installed on your system:

## Prerequisites:

- GNU Make in version 3.81+
- Go Hugo v0.80+
- Usage of Git Submodules is prohibited.
- Use the theme "ananke" for the website by following:
`Note for non-git users` at the
[Step 3](https://docs.edg.io/guides/sites_frameworks/getting_started/hugo).
- The website is expected to be generated into ./dist folder but this folder

## Lifecycle:


To build the website, run the following command from the root directory of the project:

- make build

This will generate the website from the source files and configuration files, and store the generated files in the "dist/" directory.

## Cleaning the Build Directory

To clean up the build directory, run the following command:

- make clean

This will delete all files in the "dist/" directory.

## Creating a New Blog Post
To create a new blog post, run the following command:

- make post POST_NAME=<name of post> POST_TITLE=<title of post>

This will create a new blog post with the specified name and title in the "posts/" directory.

## Help
For more information on the available targets and their usage, run the following command:

- make help

This will display a list of available targets and a brief description of each one.

## Code Comments

The source code for this project includes comments to provide additional context and information for readers. Please refer to these comments for more information on the inner workings of the code.
