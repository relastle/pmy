pmy
---

pmy is a highly customizable context-aware shell(zsh)-completion scheme using
[fzf](https://github.com/junegunn/fzf).
I'm fully in love with fzf, and I think [zsh's completion system](http://zsh.sourceforge.net/Doc/Release/Completion-System.html#Completion-System) is so complicated (I admit it is very powerful), so I developed this system.

Dependency
---

- [fzf](https://github.com/junegunn/fzf) (You can of course use other fuzzy finder such as [peco](https://github.com/peco/peco) and [fzy](https://github.com/jhawthorn/fzy) instead of fzf)
- [go](https://github.com/golang/go)
    - [go-shellwords](https://github.com/mattn/go-shellwords)
    - [go-pipeline](https://github.com/mattn/go-pipeline)
- awk (available in almost all environment)

Getting started
---

## Installation

First, pls get pmy(backend system written in Go) using go get.
```sh
go get -u github.com/relastle/pmy
```

Then, source a zsh script which simply configure brief settings.

```zsh
source "${GOPATH}/src/github.com/relastle/pmy/shell/pmy.zsh"
```

You can also add the line into your ~/.zshrc if you want.

## Demonstration

### Git

#### git checkout(co)

![pmy_git_checkout](https://user-images.githubusercontent.com/6816040/59522651-2c77bb80-8f0a-11e9-9c73-9b530ec58d75.gif)

### cd

### Postfix completion

[License](LICENSE)
------------------

The MIT License (MIT)
