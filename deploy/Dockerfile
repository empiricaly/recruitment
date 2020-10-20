FROM golang:alpine AS build-env

RUN apk add --update nodejs yarn build-base git
RUN go get github.com/GeertJohan/go.rice github.com/GeertJohan/go.rice/rice

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

ADD . .

RUN go generate -x ./internal/server
RUN go generate -x ./internal/mturk

RUN go build -race -ldflags="-extldflags=-static" -tags sqlite_omit_load_extension -o recruitment ./cmds/recruitment

FROM scratch

COPY --from=build-env /build/recruitment /

ENTRYPOINT [ "/recruitment" ] 
