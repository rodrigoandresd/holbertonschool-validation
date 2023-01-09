.PHONY: build clean post

build:
	hugo -d dist/

clean:
	rm -r dist/*

post:
	POST_TITLE=$(shell cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)
	POST_NAME=$(shell cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)
	hugo new posts/$(POST_NAME).md
	sed -i "s/title: .*/title: $(POST_TITLE)/" ./content/posts/$(POST_NAME).md
