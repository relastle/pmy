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

![pmy_git_checkout](https://user-images.githubusercontent.com/6816040/59527158-cdb83f00-8f15-11e9-84b1-807360610ad3.gif)

#### git cherry-pick(cp)

![pmy_git_cherry-pick](https://user-images.githubusercontent.com/6816040/59527161-cee96c00-8f15-11e9-9b8f-749a7997ced0.gif)

### cd

![pmy_cd](https://user-images.githubusercontent.com/6816040/59527164-d0b32f80-8f15-11e9-97ac-0127bccea233.gif)


### Postfix completion


#### <number>.**for**
    
![pmy_postfix_numfor](https://user-images.githubusercontent.com/6816040/59527168-d27cf300-8f15-11e9-9450-6a930b248f74.gif)

    
#### <general command>.**for**
    

![pmy_postfix_general](https://user-images.githubusercontent.com/6816040/59527179-d446b680-8f15-11e9-8d34-ffa3713c04e9.gif)

    

[License](LICENSE)
------------------

The MIT License (MIT)
