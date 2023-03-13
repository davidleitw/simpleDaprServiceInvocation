FROM golang:1.18 AS build

WORKDIR /func
COPY . .  
RUN go mod tidy
RUN go build -o handler main.go

FROM alpine:latest
WORKDIR /func
RUN apk add --no-cache libc6-compat 
COPY --from=build /func/ .
RUN chmod +x handler
CMD [ "./handler" ]