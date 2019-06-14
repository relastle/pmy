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

![pmy_git_checkout](https://user-images.githubusercontent.com/6816040/59523320-ef142d80-8f0b-11e9-9ec9-f957154dbf70.gif)

#### git cherry-pick(cp)

![pmy_git_cherry-pick](https://user-images.githubusercontent.com/6816040/59523324-f0455a80-8f0b-11e9-8dbb-b327b7c0af0e.gif)

### cd

![pmy_cd](https://user-images.githubusercontent.com/6816040/59524373-a9a52f80-8f0e-11e9-91e5-00fbdb8c6784.gif)


### Postfix completion


#### <number>.**for**
    
![pmy_postfix_10for](https://user-images.githubusercontent.com/6816040/59524893-f50c0d80-8f0f-11e9-859c-9dc89cb1f00a.gif)

    
#### <general command>.**for**
    
![pmy_postfix_general_for](https://user-images.githubusercontent.com/6816040/59525546-98115700-8f11-11e9-8a4f-26453e86feb8.gif)

    

[License](LICENSE)
------------------

The MIT License (MIT)
