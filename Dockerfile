FROM golang:1.17-alpine
WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .


#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/main -ldflags '-s -w'
#CMD ["go", "run", "cmd/main.go"]
#FROM scratch as runner
#EXPOSE 8080
#COPY --from=dev /go/bin/main /app/main
#COPY --from=dev /go/src/app/chat/templates /app/chat/templates
#CMD ["/app/main"]