# LikeTwitterApp-backend

## 💬 About

This repository is backend on LikeTwitterApp.

## 🌻 Version

||Name|Version|What|
|:-:|:-:|:-:|:-|
|backend|golang|1.12.5|High-level languages|
||gin|1.4.0|Web FrameWork|
||gorm|1.9.8|ORM|
|DB|Postgresql|11.5|Database|

## 🔰 Install & Setup

#### 1. Download Docker

The following procedure, please install Docker For Mac or Docker For Windows

[https://docs.docker.com/install/](https://docs.docker.com/install/)

#### 2. Getting source code

```bash
git clone git@github.com:katsuomi/LikeTwitterApp-backend.git
cd LikeTwitterApp-backend
```

#### 3. Start-up

The following procedure, start the container.

```bash
# Create Docker image
$ docker-compose build

# Start Docker container
$ docker-compose up

# confirm
$ docker-compose ps

# in container shell
$ docker-compose exec like_twitter_app_backend /bin/bash 

# run server
$ go run main.go
```
