pmy
---

pmy is a highly customizable context-aware shell(zsh)-completion system using
[fzf](https://github.com/junegunn/fzf).
I'm fully in love with fzf, and I think [zsh's completion system](http://zsh.sourceforge.net/Doc/Release/Completion-System.html#Completion-System) is so complicated (I admit it is very powerful), so I developed this system.

Dependency
---

- [fzf](https://github.com/junegunn/fzf)
- [go](https://github.com/golang/go)
- awk (available in almost all environment)


Getting started
---

## Installation

First, pls get pmy(backend system written in Go) using go get.
```sh
go get github.com/relastle/pmy
```

Then, source a zsh script which simply configure brief settings.

```zsh
source "${GOPATH}/src/github.com/relastle/pmy/shell/pmy.zsh"
```

You can also add the line into your ~/.zshrc if you want.

## Basic examples

Default setting is stored in the file ${PMY_CONFIG_PATH}.

Here is its content

```json
[
    {
      "regexpLeft": "^(vi||vim|nvim) ",
      "command": "find . -maxdepth 5",
      "bufferLeft": "[]",
      "bufferRight": "[]"
    },
    {
      "regexpLeft": "^(python|python|python3|pypy|pypy3) ",
      "command": "\\find . | egrep py$",
      "bufferLeft": "[]",
      "bufferRight": "[]"
    },
    {
      "regexpLeft": "^cd +(?P<path>.*)$",
      "command": "\\ls ${PMY_LS_OPTION} -1 <path> | egrep '/$'",
      "bufferLeft": "[]",
      "bufferRight": "[]"
    },
    {
      "regexpLeft": "(?P<cmd>.+)\\.for",
      "command": "echo ''",
      "bufferLeft": "for x in $(<cmd>); do ",
      "bufferRight": "; done"
    }
]
```

Let me describe a way of setting and its behavior one by one.

### 1

```json
{
  "regexpLeft": "^(vi||vim|nvim) ",
  "command": "find . -maxdepth 5",
  "bufferLeft": "[]",
  "bufferRight": "[]"
}
```

Anyway, let's launch this comletion system by
hitting Ctrl + Space key (, which is common key bind for completion in many popular environments such as Eclipse and IntelliJ IDEA) in the following context

```zsh
vim <your cursor here>
```

You can see that fzf launches from results of command
```zsh
"find . -maxdepth 5"
```

By selecting a candidate from them, the item selected is inserted into the cursor position.

Now, let me describe the meaning of settings

- "regexpLeft" configures when invokes pmy by defining regular expression of left buffer (left text to the cursor).
- "regexpRight" configures when invokes pmy by defining regular expression of right buffer (right text to the cursor). I think this option is rarely specified because you usually type command in a left-to-right manner.
- "command" configures the shell command, whose results will be passed to fzf. A selected candidate by fzf will be inserted to the cursor position.
- "bufferLeft" configures the resulted text in left buffer. if you want the buffer at it is, you should set its value to "[]", which indicates the original left buffer text.
- "bufferRight" configures the resulted text in right buffer.

This example resemble the fzf default <Ctrl-T> widget, so let us move to the next example.

### 2

```json
{
  "regexpLeft": "^(python|python|python3|pypy|pypy3) ",
  "command": "\\find . | egrep py$",
  "bufferLeft": "[]",
  "bufferRight": "[]"
}
```

when you use python interpreter, a target script has python extension in almost all cases.
So this setting enables you to select from python script when the command is a kind of python interpreter inclusing pypy. Such context-aware completion can also be invoked in the same key Ctrl + Space.


### 3

```json
{
  "regexpLeft": "^cd +(?P<path>.*)$",
  "command": "\\ls ${PMY_LS_OPTION} -1 <path> | egrep '/$'",
  "bufferLeft": "[]",
  "bufferRight": "[]"
}
```
This example realizes more powerful completion. This will make your cd(change directory)-procedure more productive.
You can see that parametrized sub-regular expression in `regexpLeft`. If the query matches the regular expression,


