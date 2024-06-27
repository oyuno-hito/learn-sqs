# learn-sqs
sqsの学習を目的にしたwebアプリケーション開発リポジトリ

# require
- docker
- docker-compose
- atlas

# システム構成図

![architecture](doc/architecture.jpg)

# usage

```shell
docker-compose up -d
curl -XPOST http://localhost:8080/messages -d '{"message": "test message"}'
```

# migration

```shell
cd app/cmd/migrate
atlas migrate diff --env "gorm"
cd -
```

```shell
docker-compose up migrate
```


# lambda
ログ確認
```shell
aws --endpoint-url=http://localhost:4566 logs filter-log-events \
--log-group-name /aws/lambda/handler \
--query "events[].[message]" \
--output text

```
