# pmy

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://opensource.org/licenses/Apache-2.0)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Frelastle%2Fpmy.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Frelastle%2Fpmy?ref=badge_shield)
[![CircleCI](https://circleci.com/gh/relastle/pmy.svg?style=svg)](https://circleci.com/gh/relastle/pmy)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/b4d31630d2f64ef1892d74dcc2e3105e)](https://www.codacy.com/app/relastle/pmy?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=relastle/pmy&amp;utm_campaign=Badge_Grade)

pmy is a highly customizable context-aware shell(zsh)-completion scheme utilizing fuzzy finder such as
[fzf](https://github.com/junegunn/fzf).
I'm fully in love with fzf, and I think [zsh's completion system](http://zsh.sourceforge.net/Doc/Release/Completion-System.html#Completion-System) is so complicated (I admit it is very powerful), so I developed this system.

## Dependency

-   [fzf](https://github.com/junegunn/fzf) (You can of course use other fuzzy finder such as [peco](https://github.com/peco/peco) and [fzy](https://github.com/jhawthorn/fzy) instead of fzf.)
-   [go](https://github.com/golang/go)
    -   [go-shellwords](https://github.com/mattn/go-shellwords)
    -   [go-pipeline](https://github.com/mattn/go-pipeline)
    -   [color](https://github.com/fatih/color) (Used in test)

-   awk (available in almost all environment.)

## Installation

First, please get pmy(backend system written in Go) using go get.
```sh
go get -u github.com/relastle/pmy
```

Then, source a zsh script which simply configure brief settings.

```zsh
source "${GOPATH}/src/github.com/relastle/pmy/shell/pmy.zsh"
```

You can also add the line into your ~/.zshrc if you want.

## Basic Usage

Pmy can be invoked by <kbd>Ctrl</kbd> + <kbd>Space</kbd>.
If the current left buffer (the part of the buffer that lies to the left of the cursor position) and the right buffer (the right part) match pre-defined rule (described below), fuzzy-finder launches against outputs of the corresponding command. 

## Basic Configuration

Basically, pmy's compleion behavior is solely based on one json file specified with `${PMY_RULE_PATH}`.
Default setting is [here](https://github.com/relastle/pmy/blob/master/resources/pmy_rules.json).

Rule unit is described as follows

```json
{
  "regexpLeft": "git (cp|cherry-pick) *$",
  "cmdGroups": [
    {
      "tag": "🍒:commit",
      "stmt": "git log --oneline --branches --tags",
      "after": "awk '{print $1}'"
     }
   ],
   "bufferLeft": "[]",
   "bufferRight": "[]"
}
```
| property name         | description                                                                                                             |
| ---                   | ---                                                                                                                     |
| ***regexpLeft***      | If this regexp matches the current left buffer, this rule will be activated.                                            |
| ***regexpRight***     | Same as left one. But in many cases it is not set as default value '' (becasue you usually work in line left to right). |
| ***cmdGroups.tag***   | tag string which will be inserted ahead of each line of outputs of the corresponding command.                           |
| ***cmdGroups.stmt***  | command that will be executed to make sources for fuzzy-finder.                                                         |
| ***cmdGroups.after*** | command that will be executed against line after fuzzy-finder selection (using pipe).                                   |
| ***bufferLeft***      | Buffer left values after completion. [] denotes the original left buffer.                                               |
| ***bufferRight***     | Buffer right values after completion. [] denotes the original right buffer.                                             |

## Demonstration

Here, some of examples of pmy's completion are provided as GIF with its rule(json format).
They are just a few examples of all possible pattern-matching based completion, but I think it help you to create new pmy's rule.

### git checkout(co)

![pmy_git_checkout_resized](https://user-images.githubusercontent.com/6816040/59544897-a5e6cc80-8f51-11e9-8b6a-656734d159b0.gif)

```json
{
  "regexpLeft": "(?P<body>git (co|checkout)) *(?P<query>.*)$",
  "cmdGroups": [
    {
      "tag": "🌱:branch",
      "stmt": "git branch --format=\"%(refname:short)\"",
      "after": "awk '{print $0}'"
    },
    {
      "tag": "🍺:commit",
      "stmt": "git log --oneline -10",
      "after": "awk '{print $1}'"
    }
  ],
  "fuzzyFinderCmd": "fzf -0 -1 -q \"<query>\"",
  "bufferLeft": "<body> ",
  "bufferRight": "[]"
}

```

### git cherry-pick(cp)

![pmy_git_cherry-pick_resized](https://user-images.githubusercontent.com/6816040/59544901-a67f6300-8f51-11e9-91f9-16e668b25af7.gif)

```json
{
  "regexpLeft": "git (cp|cherry-pick) *$",
  "cmdGroups": [
    {
      "tag": "🍒:commit",
      "stmt": "git log --oneline --branches --tags",
      "after": "awk '{print $1}'"
    }
  ],
  "bufferLeft": "[]",
  "bufferRight": "[]"
}

```

### cd

![pmy_cd_resized](https://user-images.githubusercontent.com/6816040/59544895-a54e3600-8f51-11e9-894a-22beac49014e.gif)

```json
{
  "regexpLeft": "^cd +(?P<path>([^/]*/)*)(?P<query>[^/]*)$",
  "cmdGroups": [
    {
      "tag": "",
      "stmt":  "command ls ${PMY_LS_OPTION} -1 <path> | egrep '/$'",
      "after": "awk '{print $0}'"
    }
  ],
  "fuzzyFinderCmd": "fzf -0 -1 -q \"<query>\"",
  "bufferLeft": "cd <path>",
  "bufferRight": "[]"
}

```

### Postfix completion

Pmy's completion rule is highly customizable and flexible, you can easily create a rule that performs ***postfix-completion*** .

#### generate for loop iterating from 1 to a given number

![pmy_postfix_numfor_resized](https://user-images.githubusercontent.com/6816040/59544899-a5e6cc80-8f51-11e9-82ca-a149620264cb.gif)

```json
{
  "regexpLeft": "^(?P<num>[1-9][0-9]*).for$",
  "cmdGroups": [
    {
      "tag": "",
      "stmt":  "echo ''",
      "after": "awk '{print $0}'"
    }
  ],
  "bufferLeft": "for x in $(seq 1 <num>); do ",
  "bufferRight": "; done"
}

```

#### generate for loop iterating each line of outputs from a given command

![pmy_postfix_general_resized](https://user-images.githubusercontent.com/6816040/59544900-a5e6cc80-8f51-11e9-8c86-1a88a417b11e.gif)

```json
{
  "regexpLeft": "(?P<cmd>.+)\\.for$",
  "cmdGroups": [
    {
      "tag": "",
      "stmt":  "echo ''",
      "after": "awk '{print $0}'"
    }
  ],
  "bufferLeft": "for x in $(<cmd>); do ",
  "bufferRight": "; done"
}
```

## [License](LICENSE)
------------------

The MIT License (MIT)
