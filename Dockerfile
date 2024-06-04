FROM golang:1.22 as api

WORKDIR /opt/app
COPY . .

CMD [ "go", "run", "app/cmd/api/main.go" ]
