Trivial go redirect app
=======================

This small application will listen to a {ip,}:socket to a given
http url. So if you have a public "example.com" IP and want to
redirect requests to it to "www.example.com" run on example.com:
```
$ redirect-app -from ":80" -to "http://www.example.com"
```
press ctrl-c to stop it again. As its written in GO it can be build
for linux/windows/macos (and some BSDs).


Build & Test
------------

To get started run:
```
$ . env.sh
```
to setup the GO workspace environment (its just GOPATH really)

Then build with:
```
$ go build redirect-app
```

You can run the tests with:
```
$ go test redirect-lib
```

And actually run the code via:
```
$ ./redirect-app
```

To cross-build for windows do:
```
$ sudo apt-get install golang-go-windows-386
$ GOARCH=386 GOOS=windows go build redirect-app 
```

