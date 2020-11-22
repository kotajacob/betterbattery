# betterbattery
# See LICENSE for copyright and license details.
.POSIX:

include config.mk

all: clean build

build:
	go build -ldflags "-X main.Version=$(VERSION)"
	scdoc < betterbattery.1.scd | sed "s/VERSION/$(VERSION)/g" > betterbattery.1

clean:
	rm -f betterbattery
	rm -f betterbattery.1

install: build
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f betterbattery $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/betterbattery
	mkdir -p $(DESTDIR)$(MANPREFIX)/man1
	cp -f betterbattery.1 $(DESTDIR)$(MANPREFIX)/man1/betterbattery.1
	chmod 644 $(DESTDIR)$(MANPREFIX)/man1/betterbattery.1

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/betterbattery
	rm -f $(DESTDIR)$(MANPREFIX)/man1/betterbattery.1

.PHONY: all build clean install uninstall
