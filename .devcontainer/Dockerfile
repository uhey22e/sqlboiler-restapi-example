FROM golang:1.16

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN apt-get update && \
    export DEBIAN_FRONTEND=noninteractive && \
    apt-get -y install --no-install-recommends \
        zsh
RUN chsh -s /bin/zsh

RUN go install \
    github.com/volatiletech/sqlboiler/v4@latest \
    github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
