FROM golang
WORKDIR /go/src/github.com/CheoGnauhs/http-test
ADD . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix .
EXPOSE 8080
ENTRYPOINT ./http-test