FROM golang:1.11 

RUN apt-get update && apt-get upgrade -y && apt-get -y install libssl-dev

WORKDIR /go/src/github.com/quorumcontrol/ipfsplay

COPY . .

RUN go install -v

CMD ["/bin/bash"]

