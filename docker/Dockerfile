FROM golang:alpine
LABEL maintainer "Hiroki Konishi <relastle@gmail.com>"

# Install zsh
RUN apk update && \
    apk add --no-cache libc-dev zsh git gcc make && \
    rm -f /tmp/* /etc/apk/cache/*
ENV SHELL /bin/zsh
ENV GOPATH /go

# Install fzf
RUN go get -u github.com/junegunn/fzf

# Install pmy
WORKDIR /go/src/github.com/relastle/pmy
COPY . .
RUN go get && go build && go install
RUN ln -s "/go/src/github.com/relastle/pmy/shell/zshrc.minimal" "${HOME}/.zshrc"
