FROM golang:1.18.2 as builder

#define build env
ENV GOOS linux
ENV CGO_ENABLED 0

#Setup folder
RUN mkdir /jumiamicroservices

#Add a work directory
WORKDIR   /jumiamicroservices

#cache and install dependencies
COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

#Copy app files
COPY . .

#install Reflex for development
RUN go install github.com/cespare/reflex@latest

#Expose port
EXPOSE   80
EXPOSE   60001
WORKDIR  /jumiamicroservices



