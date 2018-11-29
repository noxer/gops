# GoPS
The goal of GoPS to offer a fast powerline-like prompt.
![image showing the effects of the new prompt](https://raw.githubusercontent.com/noxer/gops/master/img/example.png)

## Installation
Since GoPS is under heavy development right now, I'm not offering a binary to download. To install it, you've got to build it yourself.

### Building from source
GoPS relies on modded fonts (namely powerline-fonts) to display some of the symbols. Make sure to install it and activate it for your terminal: <https://github.com/powerline/fonts#installation>

Make sure the current Go version is installed. If not, refer to this guide: <https://golang.org/doc/install#install>

Also make sure your `$GOPATH/bin` directory is in the `$PATH`.

Now you can install GoPS by running

```go
go get -u github.com/noxer/gops
```

After the downloading and building of GoPS is complete, add it as your prompt. Is this case we are editing the `~/.bashrc` file and add the following line

```bash
PS1='$(gops)'
```

To apply the change, you need to either close your console and open it again or source the updated `.bashrc` file as follows:

```bash
source ~/.bashrc
```

## FAQ
(Questions that have never been asked but I think the answers may help you)

### Where do it put my config file? How do I configure GoPS?
You don't. Right now GoPS does not support any config files or command line parameters. Instead you can edit `main.go` or the segments themselfs to adapt the prompt to your needs.

### I've successfully built and activated GoPS but it prints strange characters in the prompt. How do I fix that?
You've either not installed `powerline-fonts` or it is not configured to be used in your terminal. Make sure you've installed installed it from [here](https://github.com/powerline/fonts#installation) and you've activated one of the fonts. The font used for the screenshot is `Source Code Pro for Powerline Medium`.

### I found a bug, how do I report it?
Open an issue in this repository. Make sure you include information about the used operating system and terminal.

### Does GoPS support shell XYZ?
No idea. I've tested it with `bash` in a GNOME-Terminal. It probably works with `zsh` as well.

### How do I write plugins for GoPS?
That is pretty easy. There is an `example.go` file in the `segments/` folder with a skeleton for a segment. You basically add as many segments as you wish to the provided list of segments and return it. Make sure you add your custom package to import list of the `main.go` file and call your `Add` function in the `main` function. Each segment defines a foreground and background color. If two consecutive segments have the same background color, a small separator in the foreground color is added, otherwise the full separator is inserted.
