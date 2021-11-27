FROM golang:1.16-buster as susy

RUN mkdir /node

COPY . /node

RUN cd /node && make node
