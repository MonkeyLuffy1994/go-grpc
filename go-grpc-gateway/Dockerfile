FROM golang:1.19

WORKDIR /app/gateway

COPY go.* ./

RUN go mod tidy

COPY . ./

RUN cd /app/gateway/cmd/ && go build -o /server

EXPOSE 3000

CMD ["/server"]