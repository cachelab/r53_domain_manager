NAME := r53_domain_manager
MAINTAINER := cachelab
VERSION := $(shell grep "const version =" service.go | cut -d\" -f2)
AWS_REGION := us-east-1
AWS_ACCESS_KEY_ID=$(shell aws configure get aws_access_key_id)
AWS_SECRET_ACCESS_KEY=$(shell aws configure get aws_secret_access_key)

.PHONY: *

default: run

run: build
	docker run \
	-e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
	-e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
	-e AWS_REGION=${AWS_REGION} \
	${MAINTAINER}/${NAME}

build: vet
	@echo Building Binary and Container
	@go build -o ${NAME}
	@docker build -t ${MAINTAINER}/${NAME} .
	@GOOS= go build -o ${NAME}

vet:
	@echo Formatting Code
	@go fmt ./...
	@echo Vetting Code
	@go vet .

push: build
	docker tag ${MAINTAINER}/${NAME}:latest ${MAINTAINER}/${NAME}:${VERSION}
	docker push ${MAINTAINER}/${NAME}:latest
	docker push ${MAINTAINER}/${NAME}:${VERSION}

test:
	@echo Running Unit Tests
	@mkdir -p .coverage
	@GOOS=darwin \
		go test -tags test -coverprofile=/tmp/cov.out ./...
	@go tool cover -html=/tmp/cov.out -o=.coverage/cov.html
	@open .coverage/cov.html

tag:
	git tag v${VERSION}
	git push origin --tags
