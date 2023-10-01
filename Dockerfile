FROM golang:1.21-bookworm as build 
WORKDIR /app

COPY . /app

RUN go build . 


FROM golang:1.21-bookworm as dev 
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest && \
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


COPY go.mod go.sum /app/

RUN go mod download

COPY . /app

CMD air


FROM golang:1.21-bookworm as prod 

WORKDIR /app

COPY --from=build /app/chat ./chat

CMD [ "./chat" ]

