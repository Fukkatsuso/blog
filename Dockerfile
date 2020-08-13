FROM golang:1.14.4-alpine3.12

RUN apk update && \
    apk --no-cache add \
        git

RUN mkdir $HOME/src && \
    cd $HOME/src && \
    git clone https://github.com/gohugoio/hugo.git && \
    cd hugo && \
    go install

EXPOSE 1313

WORKDIR /go/src/github.com/Fukkatsuso/blog