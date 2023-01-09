.PHONY: build clean post

build:
	hugo -d dist/

clean:
	rm -r dist/*

post:
	hugo new posts/$(POST_NAME).md
	sed -i "s/title: .*/title: $(POST_TITLE)/" ./content/posts/$(POST_NAME).md
