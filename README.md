# learn-sqs
sqsの学習を目的にしたwebアプリケーション開発リポジトリ

# require
- docker
- docker-compose
- atlas

# システム設計図

TBD

# Usage

```shell
docker-compose up -d
curl localhost:8080/health
```

# migration

```shell
cd app/cmd/migrate
atlas migrate diff --env "gorm"
```

```shell
docker-compose up migrate
```
