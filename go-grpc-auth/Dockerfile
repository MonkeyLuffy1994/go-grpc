FROM golang:1.19

WORKDIR /app/auth

COPY go.* ./

RUN go mod tidy

COPY . ./

RUN cd /app/auth/cmd/ && go build -o /server

EXPOSE 50051

CMD ["/server"]