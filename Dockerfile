FROM golang:1.10.3-alpine3.8 AS build-env

# download glide
RUN apk update && \
    apk add --no-cache ca-certificates openssl git && \
    update-ca-certificates && \
    wget -qO- https://github.com/Masterminds/glide/releases/download/v0.13.1/glide-v0.13.1-linux-386.tar.gz | tar xzf - -C /usr/bin --strip-components=1 linux-386/glide

ADD ./glide.* /tmp/
# install dependencies from vendor
RUN cd /tmp && glide install

# copy files to GOPATH
ADD . /go/src/github.com/ns-consulting/fizz-buzz-rest-server
RUN mv /tmp/vendor /go/src/github.com/ns-consulting/fizz-buzz-rest-server/
WORKDIR /go/src/github.com/ns-consulting/fizz-buzz-rest-server

# build statically omitting the symbol table, debug information and the DWARF symbol table
# see https://golang.org/cmd/link/
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags="-s -w" -o fizz-buzz ./main.go


FROM scratch
COPY --from=build-env /go/src/github.com/ns-consulting/fizz-buzz-rest-server/fizz-buzz /fizz-buzz
ENTRYPOINT ["/fizz-buzz"]
