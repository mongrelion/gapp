TAG := $$(git tag -l | tail -1)
IMAGE := mongrelion/gapp

build:
	@go build main.go

run:
	@go run main.go

test:
	@go test .

dist: build-image push-image

build-image:
	@docker build -t $(IMAGE):$(TAG) .

push-image:
	@docker push $(IMAGE):$(TAG)
