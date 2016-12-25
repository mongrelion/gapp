TAG := $$(git tag -l | tail -1)
IMAGE := mongrelion/gapp

default: test

dist: build-image push-image

build:
	@go build main.go

run:
	@go run main.go

test:
	@go test -v .

build-image:
	@docker build -t $(IMAGE):$(TAG) .

push-image:
	@docker push $(IMAGE):$(TAG)
