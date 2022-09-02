FROM golang:1.17-bullseye

RUN apt-get update \
    && apt-get install -y unzip \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && echo "Asia/Tokyo" > /etc/timezone

WORKDIR /go/src/github.com/Fukkatsuso/blog/og-image

# 最新releaseだけをダウンロードしてttfファイルを取り出す
RUN curl -s https://api.github.com/repos/ookamiinc/kinto/releases/latest \
    | grep "browser_download_url.*zip" \
    | cut -d : -f 2,3 \
    | tr -d '"' \
    | wget -i - \
    && mkdir -p /usr/share/fonts/ \
    && unzip KintoSans.zip Kinto\ Sans/*.ttf -d /usr/share/fonts \
    && chmod 644 /usr/share/fonts/Kinto\ Sans \
    && rm KintoSans.zip

COPY go.* ./
RUN go mod download

COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -x -o /go/bin/og-image

ENTRYPOINT [ "/go/bin/og-image" ]
