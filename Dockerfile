FROM golang:1.23-alpine

WORKDIR /pb

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o /pb/pocketbase

EXPOSE 8090

ENV GOMEMLIMIT=500MiB

ADD docker_cmd.sh /docker_cmd.sh

CMD /docker_cmd.sh
