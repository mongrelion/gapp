FROM scratch

MAINTAINER Carlos León, mail@carlosleon.info

ADD ./gapp.i386 /bin/gapp

ENTRYPOINT ["/bin/gapp"]

EXPOSE 8080
