# gocat

_A tiny, fast rainbow text printer for your terminal._

`gocat` takes any text, via **pipes**, **text files** or **command-line arguments**, and prints it as smoothly colored **rainbow text** using truecolor ANSI escape codes.

---

## Features

-   Smooth rainbow coloring
-   Works with pipes, files, or CLI args
-   Fast and tiny go binary
-   Pure Go implementation

## Install

### Install via go

If you have Go installed:

```sh
go install github.com/anas-shakeel/gocat@latest
```

### Download Precompiled Binaries

Visit the [Releases](https://github.com/Anas-Shakeel/gocat/releases/latest) page to download precompiled binaries. Binaries are **available only for linux**, for now atleast.

After downloading, unpack and move the binary to a folder in your `PATH`.

## Usage

### Pipe text into it:

```sh
echo "Hello, world!" | gocat
```

### Pass a text file:

```sh
gocat -file "file.txt"
```

### Or pass text directly:

```sh
gocat "Make your terminal colorful!"
```

## Compatibility

`gocat` only supports **Linux**, _(maybe `Mac`)_ for now.

## Inspiration

[`lolcat`](https://github.com/busyloop/lolcat) obviously!
