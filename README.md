# typer

⌨️ For when you don't want to type things

## Usage

```sh
$ echo 'echo "Hello world"' | typer
echo "Hello world"

$ echo "Hello world"
Hello world
```

## Installation

```sh
go install github.com/jonahgeorge/typer@latest
```

## But, why?

Recently I've found myself playing with Linux VMs in less-friendly environments (Hyper-V and Parallels). In these environments, the console typically doesn't support copy-paste which makes setting up things like OpenSSH and Tailscale quite annoying as I've needed to manually type in install commands and auth tokens. 

This little script allows me to pipe in a local file, focus the Linux VM console, and see it type away. 😂
