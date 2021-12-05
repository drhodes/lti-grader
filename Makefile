SHELL := /bin/bash

HOST=www.mathtech.org
KEY=./acme-utils/key.pem
CERT=./acme-utils/certificiate.pem
FILES=lti-grader Makefile env-secret.bash acme-utils/
GO=/home/derek/bin/go/bin/go

include env-secret.bash

clean:
	$(GO) clean
build: 
	$(GO) build

deploy: build
	rsync -ravP $(FILES) derek@$(HOST):~/lti-grader/

test-connect:
	$(shell source env-secret.bash)	
	curl -X POST https://$(HOST):5001

work:
	emacs *.go problem.xml 

serve: FORCE
	@killall lti-grader || true
	@./lti-grader -secret ${GRADER_SECRET} -consumer ${GRADER_CONSUMER}

# http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk \
	'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

FORCE:
