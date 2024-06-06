FROM golang:alpine as build-stage
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o stakeholders-service

FROM alpine
COPY --from=build-stage app/stakeholders-service /usr/bin
EXPOSE 8083
ENTRYPOINT [ "stakeholders-service" ]