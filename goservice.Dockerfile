FROM fedora
RUN dnf install git -y
FROM golang:1.15
ENV - GOPATH=/go_service/ .
RUN mkdir /go_service
WORKDIR /
COPY . /go_service/
EXPOSE 60/tcp
