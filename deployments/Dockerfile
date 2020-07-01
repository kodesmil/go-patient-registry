# syntax=docker/dockerfile:experimental

FROM golang
ENV GOTRACEBACK=all
WORKDIR /src/
COPY go.mod go.sum ./
RUN go mod download

COPY app app
COPY cmd cmd
COPY internal internal

# RUN --mount=type=cache,target=/go/pkg/mod \
#     --mount=type=cache,target=/root/.cache/go-build \
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux \
    go build -gcflags='all=-N -l' -o /bin/server ./cmd/server
ENTRYPOINT ["/bin/server"]