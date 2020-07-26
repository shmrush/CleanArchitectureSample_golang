FROM golang:1.14.2-alpine

RUN apk update && \
    apk add --no-cache git gcc libc-dev

RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY ./src ./

RUN go get github.com/go-delve/delve/cmd/dlv@latest && \
    go get -u -v \
      github.com/pilu/fresh@latest \
      golang.org/x/tools/gopls@latest \
      sourcegraph.com/sqs/goreturns@latest \
      github.com/ramya-rao-a/go-outline@latest \
      github.com/acroca/go-symbols@latest \
      golang.org/x/tools/cmd/guru@latest \
      github.com/motemen/gore/cmd/gore@latest \
      golang.org/x/tools/cmd/gorename@latest \
      github.com/rogpeppe/godef@latest \
      golang.org/x/tools/cmd/goimports@latest \
      github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
      github.com/rubenv/sql-migrate/... \
      golang.org/x/lint/golint@latest 2>&1
