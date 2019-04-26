# FROM golang:8.15.1-alpine
FROM golang:latest

# ADD ./src /go/src/godockerjson
WORKDIR /go/src/godockerjson/ser2
COPY . /go/src/godockerjson/ser2
# ENV PORT=8080
# EXPOSE 8080
CMD ["go","run","main2.go"]