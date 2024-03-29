FROM fedora
RUN dnf install git -y
FROM golang:1.17
ENV - GOPATH=/go_service/ .
RUN mkdir /go_service
WORKDIR /
COPY . /go_service/
EXPOSE 3001/udp
EXPOSE 3001/tcp
