# hugoはgo1.16に対応している模様
FROM golang:1.16-buster

RUN apt-get update && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone

# install hugo
ARG HUGO_VERSION
RUN mkdir $HOME/src && \
    cd $HOME/src && \
    git clone https://github.com/gohugoio/hugo.git -b "$HUGO_VERSION" --depth 1 --single-branch && \
    cd hugo && \
    go install

ENV PORT 1313
EXPOSE 1313

WORKDIR /go/src/github.com/Fukkatsuso/blog
