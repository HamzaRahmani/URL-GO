FROM golang:1.21

WORKDIR /usr/src/urlShortner

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

ARG PORT
ARG DB_HOST
ARG DB_USER
ARG DB_PASS
ENV LISTENING_PORT=$PORT
ENV DATABASE_HOST=$DB_HOST
ENV DATABASE_USER=$DB_USER
ENV DATABASE_PASSWORD=$DB_PASS

EXPOSE 5050

COPY . .
RUN go build -v

CMD ["./urlShortner"]