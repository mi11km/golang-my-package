FROM golang:1.18-alpine as dev
WORKDIR /go/src/app
RUN apk add --no-cache gcc musl-dev
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
#WORKDIR /go/src/app/cmd
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/main -ldflags '-s -w'
#CMD ["go", "run", "cmd/main.go"]

#FROM scratch as runner
#EXPOSE 8080
#COPY --from=dev /go/bin/main /app/main
#COPY --from=dev /go/src/app/chat/templates /app/chat/templates
#CMD ["/app/main"]