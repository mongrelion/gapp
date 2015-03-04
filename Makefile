app-build:
	go build -v -i
xbuild:
	@@CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -v -i -a -o gapp.386 -tags netgo -ldflags '-w' 
docker-build:
	docker build -t mongrelion/gapp .
run:
	docker run --name gapp -p :8080 mongrelion/gapp
