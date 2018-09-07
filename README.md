# Nerd LS

A golang LS implementation that provides color and icons to `ls`.

* * *

<img src="screenshot.png" width="500">

## In Development

This project is in active development and not yet fully featured. Feel free to
check it out as development progresses.

## Install

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
0.21s user 0.29s system 76% cpu 0.656 total
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

- [athityakumar/colorls](https://github.com/athityakumar/colorls)
- [reganm/ls](https://github.com/reganm/ls)
- [illinoisjackson/even-better-ls](https://github.com/illinoisjackson/even-better-ls)
- [illinoisjackson/better-ls](https://github.com/illinoisjackson/better-ls)
