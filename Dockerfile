FROM golang:1.23 AS base

COPY . /go/src/github.com/microservices-demo/catalogue/

RUN go install github.com/DataDog/orchestrion@latest

RUN cd /go/src/github.com/microservices-demo/catalogue/ \
    && orchestrion pin \
    # && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app github.com/microservices-demo/catalogue/cmd/cataloguesvc
    && CGO_ENABLED=0 GOOS=linux go build -toolexec="orchestrion toolexec" -a -installsuffix cgo -o /app github.com/microservices-demo/catalogue/cmd/cataloguesvc

FROM golang:1.23
WORKDIR /
COPY --from=base /app /
EXPOSE 80

CMD ["/app", "-port=80"]
