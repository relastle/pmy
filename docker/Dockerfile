FROM golang:alpine
LABEL maintainer "Hiroki Konishi <relastle@gmail.com>"

# Install zsh
RUN apk update && \
    apk add --no-cache zsh git && \
    rm -f /tmp/* /etc/apk/cache/*
ENV SHELL /bin/zsh

# Install fzf
RUN go get -u github.com/junegunn/fzf
RUN go get -u github.com/relastle/taggo


# Install pmy
WORKDIR /go/src/github.com/relastle/pmy
RUN go get github.com/fatih/color
COPY . .
RUN go install
RUN ln -s "/go/src/github.com/relastle/pmy/shell/zshrc.minimal" "${HOME}/.zshrc"


