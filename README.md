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
