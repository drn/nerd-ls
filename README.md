# Nerd LS

A golang LS implementation that provides color and icons to `ls`.

[![CircleCI](https://circleci.com/gh/drn/nerd-ls.svg?style=svg)](https://circleci.com/gh/drn/nerd-ls)
[![Go Report Card](https://goreportcard.com/badge/github.com/drn/nerd-ls)](https://goreportcard.com/report/github.com/drn/nerd-ls)

* * *

<p align="center"><img src="screenshot.png" width="600"></p>

## Usage

Nerd LS is a drop-in replacement for the standard ls command.

```
Usage:
  nerd-ls [OPTIONS]

Application Options:
  -a, --all   Include directory entries whose names begin with a dot (.)
  -l, --long  List in long format
  -i, --icon  Display nerd-font icons

Help Options:
  -h, --help  Show this help message
```

## Install

```
go get github.com/drn/nerd-ls
alias ls=nerd-ls # optional
```

In order to have icons included with the -i/--icon flag, a patched
[Nerd Font](http://nerdfonts.com/)
is required.

Font Options:

* [Patched Font Directory](https://github.com/ryanoasis/nerd-fonts#patched-fonts)
* [Menlo Regular Nerd Font Complete](https://github.com/drn/dots/blob/master/fonts/Menlo%20Regular%20Nerd%20Font%20Complete.otf)

## Release

```
GITHUB_TOKEN=... goreleaser
```

## Benchmarks

<details>
<summary>click for details</summary>
<p>

ls
```
❯ hyperfine "ls" --warmup 5
Benchmark #1: ls
  Time (mean ± σ):       1.6 ms ±   0.5 ms    [User: 0.6 ms, System: 0.8 ms]
  Range (min … max):     1.0 ms …   3.3 ms    572 runs
```

[nerd-ls](https://github.com/drn/nerd-ls)
```
❯ hyperfine "nerd-ls" --warmup 5
Benchmark #1: nerd-ls
  Time (mean ± σ):       6.0 ms ±   0.7 ms    [User: 2.4 ms, System: 2.3 ms]
  Range (min … max):     5.0 ms …   8.0 ms    365 runs
```

[exa](https://github.com/ogham/exa)
```
❯ hyperfine "exa" --warmup 5
Benchmark #1: exa
  Time (mean ± σ):       8.5 ms ±   0.7 ms    [User: 4.8 ms, System: 3.0 ms]
  Range (min … max):     7.6 ms …  11.2 ms    252 runs
```

[colorls](https://github.com/athityakumar/colorls)
```
❯ hyperfine "colorls" --warmup 5
Benchmark #1: colorls
  Time (mean ± σ):     387.6 ms ±   3.4 ms    [User: 274.7 ms, System: 103.3 ms]
  Range (min … max):   381.6 ms … 391.0 ms    10 runs
```

</p>
</details>

## Credit

The following projects have provided either inspiration, configuration, or
implementation guidance. Thank you!!

- [ryanoasis/nerd-fonts](https://github.com/ryanoasis/nerd-fonts)
- [athityakumar/colorls](https://github.com/athityakumar/colorls)
- [reganm/ls](https://github.com/reganm/ls)
- [illinoisjackson/even-better-ls](https://github.com/illinoisjackson/even-better-ls)
- [illinoisjackson/better-ls](https://github.com/illinoisjackson/better-ls)

## References

- [nerdfonts.com](http://nerdfonts.com/#cheat-sheet)

## License

This project is licensed under the [MIT License](LICENSE.md)
