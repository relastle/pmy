# pmy

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://opensource.org/licenses/Apache-2.0)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Frelastle%2Fpmy.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Frelastle%2Fpmy?ref=badge_shield)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/b4d31630d2f64ef1892d74dcc2e3105e)](https://www.codacy.com/app/relastle/pmy?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=relastle/pmy&amp;utm_campaign=Badge_Grade)
[![CircleCI](https://circleci.com/gh/relastle/pmy.svg?style=shield)](https://circleci.com/gh/relastle/pmy)
[![Go Report Card](https://goreportcard.com/badge/github.com/relastle/pmy)](https://goreportcard.com/report/github.com/relastle/pmy)

pmy is a highly customizable context-aware shell(zsh)-completion scheme utilizing fuzzy finder such as
[fzf](https://github.com/junegunn/fzf).
I'm fully in love with fzf, and I think [zsh's completion system](http://zsh.sourceforge.net/Doc/Release/Completion-System.html#Completion-System) is so complicated (I admit it is very powerful), so I developed this system.

## :bulb: Dependency

-   [fzf](https://github.com/junegunn/fzf) (You can of course use other fuzzy finder such as [peco](https://github.com/peco/peco) and [fzy](https://github.com/jhawthorn/fzy) instead of fzf.)
-   [go](https://github.com/golang/go)
    -   [color](https://github.com/fatih/color)

-   awk (available in almost all environment.)

## :hammer: Installation

First, please get pmy(backend system written in Go)
and its dependency tool `taggo` using go get.
```sh
go get -u github.com/relastle/pmy github.com/relastle/taggo
```

Then, source a zsh script which simply configure brief settings.

```zsh
source "${GOPATH:-${HOME}/go}/src/github.com/relastle/pmy/shell/pmy.zsh"
```

You can also add the line into your ~/.zshrc if you want.

## :black_nib: Basic Usage

Pmy can be invoked by <kbd>Ctrl</kbd> + <kbd>Space</kbd>.
If the current left buffer (the part of the buffer that lies to the left of the cursor position) and the right buffer (the right part) match pre-defined rule (described below), fuzzy-finder launches against outputs of the corresponding command.

## :gear: Basic Configuration

### pmy's rule

Basically, pmy's compleion behavior is solely based on one json file specified with `${PMY_RULE_PATH}`.
Default setting is [here](https://github.com/relastle/pmy/blob/master/rules/pmy_rules.json).

Rule unit is described as follows

```json
{
  "regexpLeft": "git (cp|cherry-pick) *$",
  "regexpRight": "",
  "cmdGroups": [
    {
      "tag": "🍒:commit",
      "stmt": "git log --oneline --branches --tags",
      "after": "awk '{print $1}'"
     }
   ],
   "fuzzyFinderCmd": "fzf -0 -1",
   "bufferLeft": "[]",
   "bufferLeft": "[]",
   "bufferRight": "[]"
}
```
| property name         | description                                                                                                  |
| ---                   | ---                                                                                                          |
| ***regexpLeft***      | If this regexp matches the current left buffer, this rule will be activated.                                 |
| ***regexpRight***     | Same as left one, but in many cases you don't have to set it becasue you usually work in line left to right. |
| ***cmdGroups.tag***   | tag string which will be inserted ahead of each line of outputs of the corresponding command.                |
| ***cmdGroups.stmt***  | command that will be executed to make sources for fuzzy-finder.                                              |
| ***cmdGroups.after*** | command that will be executed against line after fuzzy-finder selection (using pipe).                        |
| ***fuzzyFinderCmd***  | Fuzzy finder command that will be executed (piped) against obtained command                                  |
| ***bufferLeft***      | Buffer left values after completion. [] denotes the original left buffer.                                    |
| ***bufferRight***     | Buffer right values after completion. [] denotes the original right buffer.                                  |

### command specific rule

In many cases, your own rule would be command specific ones (i.g. git-specific rule and cd-spefici-rule),
which means that setting such rules into a single one file will increase searching time and slow pmy unreasonably.
Therefore, you can define command specific rule by putting command-specific rules in the same directory as
${PMY_RULE_PATH} with an appropriate file name as follows.

```bash
├── pmy_rules.json
├── git_pmy_rules.json
├── cd_pmy_rules.json
└── tmux_pmy_rules.json
```

In this case, if your current left buffer starts with git command and pmy is invoked,
it searched for matched rule first in git_pmy_rules.json, and then pmy_rules.json.

### Magic command

You sometimes want to define the fizzy-finder sources as while content of a file.
pmy supports such function by a magic command "%".
Suppose that your rule below is defined.

```json
{
  "regexpLeft": "^git $",
  "cmdGroups": [
    {
      "tag": "",
      "stmt": "%git_sub",
      "after": "awk '{print $1}'"
    }
  ],
  "fuzzyFinderCmd": "fzf -0 -1 ",
  "bufferLeft": "git ",
  "bufferRight": "[]"
}
```

Then the sources is defined as content of file

```bash
${PMY_SNIPPET_ROOT}/git_sub.txt
```

This file's content is as follows

```txt
clone      Clone a repository into a new directory
init       Create an empty Git repository or reinitialize an existing one
add        Add file contents to the index
.
.
.
```

You can define such completion (with sub command description) in an very readable way :smile:.

### Environment variables

| variable name                | description                                                                                          | default values                                                           |
| ---                          | ---                                                                                                  | ---                                                                      |
| PMY_RULE_PATH                | It defines the path of main rule json file. Command specific json files are also defined by its path | "${GOPATH:-${HOME}/go}/src/github.com/relastle/pmy/rules/pmy_rules.json" |
| PMY_TAG_DELIMITER            | Delimiter between tag and a line of sources.                                                         | tab character ("\t")                                                     |
| PMY_FUZZY_FINDER_DEFAULT_CMD | Default fuzzy finder command used when "fuzzyFinderCmd" is not set in a rule                         | "fzf -0 -1"                                                              |
| PMY_TRIGGER_KEY              | Trigger key that invokes pmy completion                                                              | '^ '                                                                     |
| PMY_SNIPPET_ROOT             | The root directory in which pmy's snippets for magic command is located                              | "${GOPATH:-${HOME}/go}/src/github.com/relastle/pmy/snippets"             |

If you want to change these values, you should export them in .zshrc before you execute

```zsh
source "${GOPATH:-${HOME}/go}/src/github.com/relastle/pmy/shell/pmy.zsh"
```

## :trumpet: Demonstration

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
      "stmt":  "command ls -F -1 <path> | egrep '/$'",
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

## :memo: [License](LICENSE)

The MIT License (MIT)

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Frelastle%2Fpmy.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Frelastle%2Fpmy?ref=badge_large)
