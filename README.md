# EthDeploy Client [![Build Status](https://travis-ci.org/loomnetwork/client.svg?branch=master)](https://travis-ci.org/loomnetwork/client)

*warning: EthDeploy is only lightly maintained as we're only focusing on our core product. This is not related to DAppchain tech.*

Client for deploying apps on EthDeploy (private hosted sandbox-blockchains)
For more info about EthDeploy and our other projects, please see [Loom Network](https://loomx.io)

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
