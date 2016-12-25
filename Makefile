TAG := $$(git tag -l | tail -1)
IMAGE := mongrelion/gapp

default: test

build:
	@go build main.go

run:
	@go run main.go

test:
	@go test -v .

dist: build-image push-image

build-image:
	@docker build -t $(IMAGE):$(TAG) .

push-image:
	@docker push $(IMAGE):$(TAG)
