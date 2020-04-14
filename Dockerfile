FROM golang:1.14.1 AS mod
WORKDIR $GOPATH/src/github.com/kodesmil/go-patient-registry
COPY Gopkg.lock .
COPY Gopkg.toml .
COPY vendor vendor
RUN ls vendor
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

FROM golang:1.14.1 as build
COPY --from=mod $GOCACHE $GOCACHE
COPY --from=mod $GOPATH/pkg/mod $GOPATH/pkg/mod
WORKDIR $GOPATH/src/github.com/kodesmil/go-patient-registry
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -p 8 -o /bin/server ./cmd/server

FROM scratch
COPY --from=build /bin/server /bin/server
ENTRYPOINT ["/bin/server"]