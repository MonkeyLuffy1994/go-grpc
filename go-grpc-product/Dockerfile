FROM golang:1.19

WORKDIR /app/product

COPY go.* ./

RUN go mod tidy

COPY . ./

RUN cd /app/product/cmd/ && go build -o /server

EXPOSE 50052

CMD ["/server"]