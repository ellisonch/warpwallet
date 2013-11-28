# WarpWallet
This is an implementation of [https://keybase.io/warp](WarpWallet) in Go.  WarpWallet is a brain wallet generator (bit Bitcoin), originally written by [https://github.com/maxtaco](Max Krohn) and [https://github.com/malgorithms](Chris Coyne).

Except for neeeding to referencing some hash functions (described below), this package is entirely self contained.  It is released under the BSD 2-Clause license, and includes some BSD-style code from [https://github.com/thepiachu](ThePiachu).

This program has been tested under Linux and Windows.

## Install
To install, you'll need to run these from the command line:

```
go get code.google.com/p/go.crypto/scrypt
go get code.google.com/p/go.crypto/pbkdf2
go get code.google.com/p/go.crypto/ripemd160
```

Since those repositories use mercurial, you'll also need to:

```
sudo apt-get install mercurial
```
if you don't have it.

Finally, you should be able to do a:

```
go get github.com/ellisonch/warpwallet
```
If that doesn't work (I haven't had much luck with it, let me know how to fix it!) just download this repo and run `go build` in the `warpwallet` directory.  This will create a `warpwallet` executable that you can then run.

## Test
To run the test suite:

```
go test github.com/ellisonch/warpwallet
```

## Build
To run the program, if you have `$GOPATH` in your `$PATH`, you should be able to run `warpwallet`.  Otherwise, run

```
$GOPATH/bin/warpwallet
```

## Love
If you found this useful, please send me some love at `1GGCFrshLz46tdas9ZtKqX59n5UAFzR6sD` :)
