

DESTDIR := /usr/local
export GOPATH := $(shell pwd)


all:
	go build redirect-app

install:
	# binary
	install -d -m644 $(DESTDIR)/bin
	install -m755 redirect-app $(DESTDIR)/bin
	# apparmor
	# FIXME: path in apparmor hardcoded
	mkdir -p $(DESTDIR)/../etc/apparmor.d
	install -m644 redirect-app.apparmor $(DESTDIR)/../etc/apparmor.d/
	/etc/init.d/apparmor reload || true
	# init script
	# FIXME:

test:
	go test redirect-lib

win:
	GOARCH=386 GOOS=windows go build redirect-app

clean:
	rm -f redirect-app redirect-app.exe
