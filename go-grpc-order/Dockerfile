FROM golang:1.19

WORKDIR /app/order

COPY go.* ./

RUN go mod tidy

COPY . ./

RUN cd /app/order/cmd/ && go build -o /server

EXPOSE 50053

CMD ["/server"]