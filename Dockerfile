FROM golang:latest
MAINTAINER chrislu chrislu30604@gmail.com

# create the directory
RUN mkdir /app
# go ENV
ENV GOPATH /app
# Work directory
WORKDIR $GOPATH
