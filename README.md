# enkrypt
##### enkrypt is a CLI program and http server. It creates an encrypted folder in the local system that mirrors an existing folder.

### Build and run

#### Requirements

`enkrypt` requires a version of [golang](https://golang.org/) with [go modules](https://github.com/golang/go/wiki/Modules) support.

#### Build

```
# Clone.
$ git clone git https://github.com/Kifen/enkrypt.git
$ cd enkrypt

# Build
$ make build 
```

#### Run

```
# Run enkrypt
$ ./enkrypt -s [source directory] -t [target directory] -k [password]
```

#### API
http://localhost:5000/listencryptedfiles

http://localhost:5000/downloadfile?file={path-to-file}

Enkrypt uses the default port 5000 
