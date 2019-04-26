# FROM golang:8.15.1-alpine
FROM golang:latest

# ADD ./src /go/src/godockerjson
WORKDIR /go/src/godockerjson
COPY . /go/src/godockerjson
# ENV PORT=8080
# EXPOSE 8080
CMD ["go","run","main.go"]
