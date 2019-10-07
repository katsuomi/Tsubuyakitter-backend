# Gin Like Twitter API

## 💬 About

goのWebフレームワークGinを使ったtwitterもどきのバックエンドAPIです。

## 🌻 Version

||Name|Version|What|
|:-:|:-:|:-:|:-|
|backend|golang|1.12.5|高級言語|
||gin|1.4.0|Webフレームワーク|
||gorm|1.9.8|ORマッパー|
|DB|Postgresql|11.5|データベース|

## 🔰 Install & Setup

#### 1. Dockerのダウンロード

下記より、`Docker For Mac` か `Docker For Windows`をインストールして下さい。  
[https://docs.docker.com/install/](https://docs.docker.com/install/)

#### 2. ソースコードの取得

```bash
git clone git@github.com:katsuomi/gin-like-twitter-api.git
cd gin-like-twitter-api
```

#### 3. 起動

下記の手順で、コンテナを起動させて下さい。

```bash
# Dockerイメージの作成
$ docker-compose build

# Dockerコンテナ起動
$ docker-compose up -d

# 確認
$ docker-compose ps

# コンテナのシェルに入る
$ docker-compose exec gin-like-twitter-api /bin/bash 

# サーバーの立ち上げ
$ go run main.go
```

