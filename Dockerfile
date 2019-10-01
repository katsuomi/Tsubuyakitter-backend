# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
RUN mkdir -p $GOPATH/src/github.com/katsuomi/gin-gorm-twitter-api
# コンテナログイン時のディレクトリ指定
WORKDIR $GOPATH/src/github.com/katsuomi/gin-gorm-twitter-api
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . $GOPATH/src/github.com/katsuomi/gin-gorm-twitter-api
# 必要なパッケージをイメージにインストールする
RUN go get -u github.com/gin-gonic/gin && \
  go get github.com/jinzhu/gorm && \
  go get github.com/jinzhu/gorm/dialects/postgres
