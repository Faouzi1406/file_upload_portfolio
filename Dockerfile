FROM golang:1.19

WORKDIR /Users/faouzibouchkachekh/coding/some_go

COPY go.mod go.sum  ./
RUN go mod download && go mod verify

COPY . .
RUN go build 

EXPOSE 3000
CMD [ "./learning" ]
