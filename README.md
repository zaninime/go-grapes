# GRAPES: a Generic Environment for P2P Streaming

[![](https://godoc.org/github.com/zaninime/go-grapes?status.svg)](https://godoc.org/github.com/zaninime/go-grapes) [![Travis](https://img.shields.io/travis/zaninime/go-grapes.svg?maxAge=2592000)](https://travis-ci.org/zaninime/go-grapes) [![Coveralls](https://img.shields.io/coveralls/zaninime/go-grapes.svg?maxAge=2592000)](https://coveralls.io/)

This is a (partial) port of [GRAPES library](http://peerstreamer.org/GRAPES/) in pure Go. I developed this as a part of my Bachelor thesis.

This package contains the functions needed to decode the GRAPES protocol: messages, chunks, and RTP stuff. It can be used with [go-ml](https://github.com/zaninime/go-ml) to mimic the features of PeerStreamer. Chunk trading isn't implemented yet, but may be in the future.

