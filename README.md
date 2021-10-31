[![Go Report Card](https://goreportcard.com/badge/github.com/noxer/gops)](https://goreportcard.com/report/github.com/noxer/gops)
[![Build Status](https://travis-ci.org/noxer/gops.svg?branch=master)](https://travis-ci.org/noxer/gops)

# GoPS
The goal of GoPS to offer a fast powerline-like prompt.
![image showing the effects of the new prompt](https://raw.githubusercontent.com/noxer/gops/master/img/example.png)

## Installation
There are two approaches of getting GoPS: Using a release or building it from source.

GoPS relies on modded fonts (namely powerline-fonts) to display some of the symbols. Make sure to install them and activate one of them in your terminal: <https://github.com/powerline/fonts#installation>

### Using a release
- Click on ["release"](https://github.com/noxer/gops/releases) and download the newest binary for your OS.
- Rename it to `gops`.
- Place it in a folder which is in `$PATH`.

### Building from source
Make sure the current Go version is installed. If not, refer to this guide: <https://golang.org/doc/install#install>

Also make sure your `$GOPATH/bin` directory is in the `$PATH`.

Now you can install GoPS by running

```go
go get -u github.com/noxer/gops
```

## Setup
### bash
After the downloading and building of GoPS is complete, add it as your prompt. For Bash edit `~/.bashrc` and add the following lines

```bash
PS1='$(gops)'
PS2='$(gops -p 2)'
```

The first line activates GoPS as your default prompt, the second line for continuation prompts (when the command ends with `\` it will print a continuation prompt to the next line).

To apply the change, you need to either close your shell and open it again or source the updated `.bashrc` file as follows:

```bash
source ~/.bashrc
```

### zsh
After the downloading and building of GoPS is complete, add it as your prompt. For ZSH edit `~/.zshrc` and add the following lines

```zsh
setopt PROMPT_SUBST
PROMPT='$(gops -s zsh)'
```

The first line tells ZSH to evaluate the commands in the `$PROMPT` variable. The second line calls GoPS.

To apply the change, you need to either close your shell and open it again or source the updated `.zshrc` file as follows:

```zsh
source ~/.zshrc
```

## Currently available segments
Currently GoPS comes with a sall selection of segments for your prompt, which can be found in segments/

| Segment    | Displayed Information |
| ---------- | ---------- |
| git        | current git branch |
| dir        | path of the current dir (can be shortened) |
| host       | hostname (FQDN) (can be shortened) |
| nodejs     | version of the currently installed node.js |
| user       | name of the current user |
| userathost | name of the current user and the hostname combined as user@host (can be shortened) |
| virtualenv | name of the currently active python viratualenv |

## FAQ
(Questions that have never been asked but I think the answers may help you)

### Where do it put my config file? How do I configure GoPS?
You don't. Right now GoPS does not support any config files or command line parameters. Instead you can edit `main.go` or the segments themselves to adapt the prompt to your needs.

### I've successfully built and activated GoPS but it prints strange characters in the prompt. How do I fix that?
You've either not installed `powerline-fonts` or it is not configured to be used in your terminal. Make sure you've installed it from [here](https://github.com/powerline/fonts#installation) and you've activated one of the fonts. The font used for the screenshot is `Source Code Pro for Powerline`.

### I found a bug, how do I report it?
Open an issue in this repository. Make sure you include information about the used operating system and terminal.

### Does GoPS support shell XYZ?
No idea. I've tested it with `bash` in a GNOME-Terminal, and `zsh` in iTerm2. Feel free to try more combinations and let me know.

### How do I write plugins for GoPS?
That is pretty easy. There is an `example.go` file in the `segments/` folder with a skeleton for a segment. You basically add as many segments as you wish to the provided list of segments and return it. Make sure you add your custom package to import list of the `main.go` file and call your `Add` function in the `main` function. Each segment defines a foreground and background color. If two consecutive segments have the same background color, a small separator in the foreground color is added, otherwise the full separator is inserted.

## API
You can find the packages rendering the different parts here:

* Common Segments: [![GoDoc](https://godoc.org/github.com/noxer/gops/segments/common?status.svg)](https://godoc.org/github.com/noxer/gops/segments)
* Separator: [![GoDoc](https://godoc.org/github.com/noxer/gops/separator?status.svg)](https://godoc.org/github.com/noxer/gops/separator)
* Symbols: [![GoDoc](https://godoc.org/github.com/noxer/gops/symbols?status.svg)](https://godoc.org/github.com/noxer/gops/symbols)
