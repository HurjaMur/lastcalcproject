FROM golang

COPY go.mod go.sum ./
WORKDIR /

RUN go mod download && go mod verify

RUN go build -v -o / ./

CMD [ "golang" ]