# pmy

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://opensource.org/licenses/Apache-2.0)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Frelastle%2Fpmy.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Frelastle%2Fpmy?ref=badge_shield)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/b4d31630d2f64ef1892d74dcc2e3105e)](https://www.codacy.com/app/relastle/pmy?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=relastle/pmy&amp;utm_campaign=Badge_Grade)
[![CircleCI](https://circleci.com/gh/relastle/pmy.svg?style=shield)](https://circleci.com/gh/relastle/pmy)
[![Go Report Card](https://goreportcard.com/badge/github.com/relastle/pmy)](https://goreportcard.com/report/github.com/relastle/pmy)

pmy is a highly customizable context-aware shell(zsh)-completion scheme utilizing fuzzy finder such as
[fzf](https://github.com/junegunn/fzf).
I'm fully in love with fzf, and I think [zsh's completion system](http://zsh.sourceforge.net/Doc/Release/Completion-System.html#Completion-System) is so complicated (I admit it is very powerful), so I developed this system.

## Dependency

-   [fzf](https://github.com/junegunn/fzf) (You can of course use other fuzzy finder such as [peco](https://github.com/peco/peco) and [fzy](https://github.com/jhawthorn/fzy) instead of fzf.)

## Installation

First, please get pmy by go get (because the backend system is written in Go)

```sh
go get -u github.com/relastle/pmy
```

Then, only you have to do is executing folloing zsh command.

```zsh
eval "$(pmy init)"
```

You can also add the line into your ~/.zshrc if you want.

## Quick Start

Try downloading very simple pmy rule files into `$HOME/.pmy` (where pmy searches for rules by default).

```zsh
git clone https://github.com/relastle/pmy-config $HOME/.pmy
```

Then, you are already able to enjoy path completion using fuzzy finder.

![Sample GIF](https://user-images.githubusercontent.com/6816040/67204810-48a3e000-f449-11e9-8195-ee0ae1282bdb.gif)

This path-completion befavior is realized by simple yml configurations below

```yml
- regexp-left: ^(?P<body>.*?)(?P<path>~{0,1}([0-9A-Za-z_\-.]*/)+)(?P<query>[0-9A-Za-z_\-.]*)$
  cmd-groups:
  - stmt: \ls -AlFG --color=always <path> | tail -n +2 | grep --color=always "<query>"
    after: awk '{print $8}'
  fuzzy-finder-cmd: fzf -0 -1 --ansi -n8
  buffer-left: <body><path>
  buffer-right: '[]'

- regexp-left: ^(?P<body>.*?) (?P<query>[0-9A-Za-z_\-.]*)$
  cmd-groups:
  - stmt: \ls -AlFG --color=always | tail -n +2 | grep --color=always "<query>"
    after: awk '{print $8}'
  fuzzy-finder-cmd: fzf -0 -1 --ansi -n8
  buffer-left: '<body> '
  buffer-right: '[]'
```

Customization is very easy.

-  Wrtie your own rule in JSON/YML format.
-  Save it in the name of `pmy_rules.[json|yml]`
-  Locate the file under one of `$PMY_RULE_PATH`.

If you want to investigate into further examples, see [Gallery](https://github.com/relastle/pmy/wiki/Gallery).

## Basic Usage

Pmy can be invoked by <kbd>Ctrl</kbd> + <kbd>Space</kbd>.
If the current left buffer (the part of the buffer that lies to the left of the cursor position) and the right buffer (the right part) match pre-defined rule (described below), fuzzy-finder launches against outputs of the corresponding command.

## Basic Configuration

### pmy's rule

Basically, pmy's compleion behavior is solely based on `rule`.

A single rule is described as follows

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
| property name (JSON / YML)                     | description                                                                                                  |
| ---                                            | ---                                                                                                          |
| ***regexpLeft*** / ***regexp-left***           | If this regexp matches the current left buffer, this rule will be activated.                                 |
| ***regexpRight*** / ***regexp-right***         | Same as left one, but in many cases you don't have to set it becasue you usually work in line left to right. |
| ***cmdGroups.tag*** / ***cmd-groups.tag***     | tag string which will be inserted ahead of each line of outputs of the corresponding command.                |
| ***cmdGroups.stmt*** / ***cmd-groups.stmt***   | command that will be executed to make sources for fuzzy-finder.                                              |
| ***cmdGroups.after*** / ***cmd-groups.after*** | command that will be executed against line after fuzzy-finder selection (using pipe).                        |
| ***fuzzyFinderCmd*** / ***fuzzy-finder-cmd***  | Fuzzy finder command that will be executed (piped) against obtained command                                  |
| ***bufferLeft*** / ***buffer-left***           | Buffer left values after completion. [] denotes the original left buffer.                                    |
| ***bufferRight*** / ***buffer-right***         | Buffer right values after completion. [] denotes the original right buffer.                                  |

### Rule configuration

pmy searchs for its rule setting in `${HOME}/pmy/rules` by default.
So you can use common settings by executing

```zsh
git clone https://github.com/relastle/pmy-config ~/.pmy
```

If you want to customize pmy's rule.
You can define environment variable `PMY_RULE_PATH`.
if you want to add `/path/to/1` and `/path/to/2` into pmy's rules paths,
execute
```zsh
export PMY_RULE_PATH="/path/to/1:/path/to/2"
```
This setting is similar that of that of `$PATH` variable (, which controlls paths where executable binaries and scripts are located).

In this situation, priorities as follows:

-  1. `/path/to/1/hoge_pmy_rules.json`
-  2. `/path/to/2/hoge_pmy_rules.json`
-  3. `${HOME}/.pmy/rules/hoge_pmy_rules.json`

### command specific rule

In many cases, your own rule would be command specific ones (i.g. git-specific rule and cd-spefici-rule),
which means that setting such rules into a single one file will increase searching time and slow pmy unreasonably.
Therefore, you can define command specific rule by putting command-specific rules in the same directory as
`${PMY_RULE_PATH}` with an appropriate file name as follows.

```zsh
├── pmy_rules.json
├── git_pmy_rules.json
├── cd_pmy_rules.json
└── tmux_pmy_rules.json
```

In this case, if your current left buffer starts with git command and pmy is invoked,
it searched for matched rule first in `git_pmy_rules.json` (git-specific rules), and then `pmy_rules.json` (globally active rules).

### Environment variables

| variable name                | description                                                                                          | default values                                                           |
| ---                          | ---                                                                                                  | ---                                                                      |
| PMY_FUZZY_FINDER_DEFAULT_CMD | Default fuzzy finder command used when "fuzzyFinderCmd" is not set in a rule                         | "fzf -0 -1"                                                              |
| PMY_TRIGGER_KEY              | Trigger key that invokes pmy completion                                                              | '^ '                                                                     |

If you want to change these values, you should export them in .zshrc before you execute

```zsh
eval "$(pmy init)"
```

## Features

-  [x] JSON/YML rule-configurations.
-  [x] Customize fuzzy finder command used.
-  [x] Combining multiple command into one source.
-  [ ] Caching compiled regular expression.
-  [ ] Customizing priority of rules.

## [License](LICENSE)

The MIT License (MIT)

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Frelastle%2Fpmy.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Frelastle%2Fpmy?ref=badge_large)
