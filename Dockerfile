FROM golang:latest as build
RUN apt-get update && apt-get install ca-certificates -yq
WORKDIR /app
COPY go.mod go.sum ./ 
RUN go mod download
COPY . .
RUN go build -ldflags "-linkmode external -extldflags -static" -o digaws .

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /app/digaws /digaws
ENTRYPOINT ["/digaws"]
