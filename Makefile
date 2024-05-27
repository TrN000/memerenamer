

TARGET := /usr/bin/memerenamer

default:
	go build .

install:
	ln -sf $(CURDIR)/memerenamer $(TARGET)

uninstall:
	rm -f $(TARGET)
