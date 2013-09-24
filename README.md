Trivial go redirect app
=======================

To get started run:
```
$ . env.sh
```

To build:
```
$ go build redirect-app
```

To test:
```
$ go test redirect-lib
```

To run:
```
$ ./redirect-app
```

To cross-build for windows do:
```
$ sudo apt-get install golang-go-windows-386
$ GOARCH=386 GOOS=windows go build redirect-app 
```

