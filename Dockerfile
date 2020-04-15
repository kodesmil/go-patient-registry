FROM golang:1.14.1 AS build
WORKDIR $GOPATH/src/github.com/kodesmil/go-patient-registry
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o /bin/server ./cmd/server

FROM scratch
COPY --from=build /bin/server /bin/server
ENTRYPOINT ["/bin/server"]