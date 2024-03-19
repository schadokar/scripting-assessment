FROM golang:1.21

WORKDIR /

COPY ./gosoln/. /

RUN go mod download

RUN go build -o exc .

CMD [ "/exc" ]