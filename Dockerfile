# FROM golang:1.17 AS build

# WORKDIR /src/
# COPY . /src/
# RUN go mod download
# RUN go install
# RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/demo

# FROM scratch
# COPY --from=build /bin/demo /bin/demo
# ENTRYPOINT ["/bin/demo"]

FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY go.sum .

# COPY .env .

RUN go mod download

COPY . .

EXPOSE 8080

RUN go build ./

CMD [ "./go-crud-mvc" ]