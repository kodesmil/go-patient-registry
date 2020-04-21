FROM golang:1.14.1 AS build

RUN go get -u -v github.com/go-delve/delve/cmd/dlv

WORKDIR /src/

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN GO111MODULE=on GOTRACEBACK=all CGO_ENABLED=0 GOOS=linux go build -gcflags='all=-N -l' -o /bin/server ./cmd/server

ENTRYPOINT ["dlv", "exec", "/bin/server", "--continue", "--accept-multiclient", "--api-version=2", "--headless", "--listen=:3000", "--log"]

# FROM scratch
# COPY --from=build /bin/server /bin/server
# ENTRYPOINT ["/bin/server"]