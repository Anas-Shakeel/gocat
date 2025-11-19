# gocat

_A tiny, fast Go-powered rainbow text printer for your terminal._

`gocat` takes any text, via **pipes**, **text files** or **command-line arguments**, and prints it as smoothly colored **rainbow text** using truecolor ANSI escape codes.

---

## Usage

### Pipe text into it:

```sh
echo "Hello, world!" | gocat
```

## Features

-   Smooth rainbow coloring
-   Works with pipes
-   Pure Go implementation

## Compatibility

`gocat` only supports **UNIX-like** operating systems for now. (e.g `Linux`, `Mac`)

## Inspiration

[`lolcat`](https://github.com/busyloop/lolcat) obviously!
