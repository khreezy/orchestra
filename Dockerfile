FROM golang:1.20 AS worker
WORKDIR /app/worker
COPY app/worker/go.mod app/worker/go.sum ./
RUN go mod download
COPY app/worker ./
RUN go build -buildvcs=false -ldflags "-s -w -extldflags '-static'" -tags osusergo,netgo -o worker .

