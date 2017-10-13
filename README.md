# Loom Client [![Build Status](https://travis-ci.org/loomnetwork/client.svg?branch=master)](https://travis-ci.org/loomnetwork/client)
Client for deploying apps on the loom network 
For more info about Loom Network please see [Loom Network](https://loomx.io)

[API Docs](https://loomx.io/docs/)

```
loom setapikey 1234
loom deploy application.zip application_name
```

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