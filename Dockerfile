FROM golang:1.12.5 as builder

WORKDIR /server
RUN mkdir -p /server/bin
ADD go /server/go
ADD vendor /server/vendor
ADD  Makefile /server/Makefile
ADD  go.mod /server/go.mod
ADD  go.sum /server/go.sum
RUN make build

ENTRYPOINT ["/server/bin/server-linux"]