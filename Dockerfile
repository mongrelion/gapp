FROM google/golang

MAINTAINER Carlos León, mail@carlosleon.info

ADD . /gopath/src/github.com/mongrelion/gapp

RUN ["go", "install", "github.com/mongrelion/gapp"]

ENTRYPOINT ["gapp"]

EXPOSE 8080
