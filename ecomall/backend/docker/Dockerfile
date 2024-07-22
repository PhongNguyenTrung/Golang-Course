FROM golang:1.20.5

WORKDIR /app

ENV config=docker

COPY .. /app

RUN go mod download

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 9092

ENTRYPOINT CompileDaemon --build="go build cmd/api/main.go" --command=./main

RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
CMD [ "/go/bin/dlv", "--listen=:4000", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/app/main" ]
