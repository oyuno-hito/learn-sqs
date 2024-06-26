FROM golang:1.22 as api

WORKDIR /opt/app
COPY . .

CMD [ "go", "run", "app/cmd/api/main.go" ]


FROM golang:1.22 as migrate

WORKDIR /opt/app
COPY . .

RUN curl -sSf https://atlasgo.sh | sh

CMD [ "go", "run", "app/cmd/migrate/main.go" ]

FROM golang:1.22 as lambda

WORKDIR /opt/app
COPY . .

RUN apt-get update
RUN apt-get -y install zip
