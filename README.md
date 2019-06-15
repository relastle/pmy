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

![pmy_git_checkout_resized](https://user-images.githubusercontent.com/6816040/59544897-a5e6cc80-8f51-11e9-8b6a-656734d159b0.gif)


#### git cherry-pick(cp)

![pmy_git_cherry-pick_resized](https://user-images.githubusercontent.com/6816040/59544901-a67f6300-8f51-11e9-91f9-16e668b25af7.gif)

### cd

![pmy_cd_resized](https://user-images.githubusercontent.com/6816040/59544895-a54e3600-8f51-11e9-894a-22beac49014e.gif)

### Postfix completion


#### generate for loop iterating from 1 to a given number

![pmy_postfix_numfor_resized](https://user-images.githubusercontent.com/6816040/59544899-a5e6cc80-8f51-11e9-82ca-a149620264cb.gif)
    
#### generate for loop iterating each line of outputs from a given command

![pmy_postfix_general_resized](https://user-images.githubusercontent.com/6816040/59544900-a5e6cc80-8f51-11e9-8c86-1a88a417b11e.gif)
    

[License](LICENSE)
------------------

The MIT License (MIT)
