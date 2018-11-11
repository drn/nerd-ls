# Nerd LS

A golang LS implementation that provides color and icons to `ls`.

[![CircleCI](https://circleci.com/gh/drn/nerd-ls.svg?style=svg)](https://circleci.com/gh/drn/nerd-ls)

* * *

<p align="center"><img src="screenshot.png" width="600"></p>

## Install

Requires a
[nerd-font patched font](https://github.com/ryanoasis/nerd-fonts#patched-fonts)
for the custom icons. My preferred font is a patched Menlo font that you can
find
[here](https://github.com/drn/dots/blob/master/lib/fonts/Menlo%20Regular%20Nerd%20Font%20Complete.otf).

```
go get github.com/drn/nerd-ls
alias ls=nerd-ls
```

## Benchmarks

ls
```
❯ time (repeat 100 { ls })
0.13s user 0.22s system 90% cpu 0.386 total
```

[nerd-ls](https://github.com/drn/nerd-ls)
```
❯ time (repeat 100 { nerd-ls })
0.21s user 0.29s system 86% cpu 0.574 total
```

[exa](https://github.com/ogham/exa)
```
❯ time (repeat 100 { exa })
0.60s user 0.60s system 90% cpu 1.338 total
```

[better-ls](https://github.com/illinoisjackson/better-ls)
```
time ( repeat 100 { ./lsicons.py })
6.43s user 6.96s system 91% cpu 14.699 total
```

[colorls](https://github.com/athityakumar/colorls)
```
❯ time (repeat 100 { colorls })
24.74s user 11.75s system 96% cpu 37.846 total
```

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
