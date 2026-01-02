FROM golang:1.25.5-alpine AS build
ENV CGO_ENABLED=0

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY ./*.go /workspace/
COPY ./internal /workspace/internal
RUN go build -ldflags='-s -w -extldflags "-static"'  -o "default-app"

FROM scratch AS deploy

WORKDIR /app/
COPY --from=build /workspace/default-app /usr/local/bin/default-app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /ca-certificates.crt

ENTRYPOINT [ "default-app" ]
