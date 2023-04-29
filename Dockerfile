FROM golang:1.19 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY index.html .
RUN CGO_ENABLED=0 GOOS=linux go build -o /file_server

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /file_server /file_server
EXPOSE 6060
ENTRYPOINT ["/file_server"]