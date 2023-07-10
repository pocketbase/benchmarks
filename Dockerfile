FROM golang:1.20-alpine

WORKDIR /pb

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o /pb/pocketbase

EXPOSE 8090

ADD docker_cmd.sh /docker_cmd.sh

CMD /docker_cmd.sh
