# Loom Client [![Build Status](https://travis-ci.org/loomnetwork/client.svg?branch=master)](https://travis-ci.org/loomnetwork/client)
Client for deploying apps on EthDeploy (private hosted blockchains)
For more info about Loom Network please see [Loom Network](https://loomx.io)

```
loom login
loom deploy application.zip application_name
```

Currently you can log into loom network with Github or Linkedin.

## Install via Homebrew (OSX)

```
brew install loomnetwork/homebrew-client/loom
```

## Install via Wget (Linux/OSX)

# Building 

```
glide install
go build -o loom .
```