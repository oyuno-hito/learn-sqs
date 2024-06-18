FROM golang:1.22 as api

WORKDIR /opt/app
COPY . .

CMD [ "go", "run", "app/cmd/api/main.go" ]


FROM golang:1.22 as migrate

WORKDIR /opt/app
COPY . .

RUN curl -sSf https://atlasgo.sh | sh

CMD [ "go", "run", "app/cmd/migrate/main.go" ]

FROM golang:1.22-alpine as lambda

WORKDIR /opt/app
COPY . .

RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o artifacts/bootstrap app/cmd/lambda/main.go

FROM localstack/localstack as localstack

COPY ./tools/localstack/init/ready.d/init-aws.sh /etc/localstack/init/ready.d/init-aws.sh
COPY --from=lambda /opt/app/artifacts/bootstrap /opt/files/bootstrap

RUN chmod +x /etc/localstack/init/ready.d/init-aws.sh && zip -rj /opt/files/lambda.zip /opt/files/bootstrap
