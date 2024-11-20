FROM golang:1.23.2
ENV GO111MODULE=on

RUN apt-get update && \
    apt-get install -y git bash && \
    apt-get clean

# 作業ディレクトリ指定
WORKDIR /go/src/app
# go.mod作成
RUN go mod init app
# go.mod更新&Airインストール
RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.40.4

# Air起動
CMD ["air", "-c", ".air.toml"]
