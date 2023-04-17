# goバージョン
FROM golang:1.19.1-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# boiler-plateディレクトリの作成
RUN mkdir -p /go/src/github.com/server
# ワーキングディレクトリの設定
WORKDIR /go/src/github.com/server
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/github.com/server