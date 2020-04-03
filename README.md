# Podnamez - wordlist generation for Kubernetes pod names

## What?

Podnames generates wordlists for all possible pod names for Kubernetes clusters, which by default append a five character alpha-numeric "trail".

For example, we turn this:

```
pod-name
```

Into this:

```
pod-name-bbbbb
pod-name-bbbbc
pod-name-bbbbd
pod-name-bbbbf
pod-name-bbbbg
pod-name-bbbbh
pod-name-bbbbj
pod-name-bbbbk
pod-name-bbbbl
pod-name-bbbbm
...
```

## How?

```
$ go run podnamez.go pod-name > pod-name-permutations.txt
```

## Why?

There are some cases where applications (incorrectly!) use Kubernetes pod names in a security sensitive context. 

In this case, it is useful to have a wordlist builder.

The wordlist has been pre-generated and included in this repo in the event that it can be used manually (think: TurboIntruder).


