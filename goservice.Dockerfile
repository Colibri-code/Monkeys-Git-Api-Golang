FROM fedora
RUN dnf install git -y
FROM golang:1.14
ENV - GOPATH=/code/ .
RUN mkdir /code
WORKDIR /code
COPY . /code/
EXPOSE 60/tcp
