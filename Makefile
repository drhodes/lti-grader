HOST=www.mathtech.org

sendit:
	rsync -ravP lti-grader derek@$(HOST):~/

test-connect:
	curl -X POST https://$(HOST):5001

build: FORCE
	$(shell go build)

deploy: build sendit

FORCE:
