# Targets:
# all: Build code.
# clean: Remove build artifacts.
# $FILENAME.go: Build just $FILENAME.
# build: Create build directory.

GOC=go build -compiler gccgo
GOFLAGS=

.PHONY: all

all: build basename.go cat.go cp.go dirname.go echo.go false.go head.go mkdir.go nl.go pwd.go seq.go sleep.go tail.go tee.go touch.go true.go xxd.go yes.go

.PHONY: clean

clean:
	rm -rf build

build:
	mkdir -p build

basename.go:
	${GOC}  -o build/basename basename.go

cat.go:
	${GOC} -o build/cat cat.go

cp.go:
	${GOC} -o build/cp cp.go

dirname.go:
	${GOC} -o build/dirname dirname.go

echo.go:
	${GOC} -o build/echo echo.go

false.go:
	${GOC} -o build/false false.go

head.go:
	${GOC} -o build/head head.go

mkdir.go:
	${GOC} -o build/mkdir mkdir.go

nl.go:
	${GOC} -o build/nl nl.go

pwd.go:
	${GOC} -o build/pwd pwd.go

seq.go:
	${GOC} -o build/seq seq.go

sleep.go:
	${GOC} -o build/sleep sleep.go

tail.go:
	${GOC} -o build/tail tail.go

tee.go:
	${GOC} -o build/tee tee.go

touch.go:
	${GOC} -o build/touch touch.go

true.go:
	${GOC} -o build/true true.go

xxd.go:
	${GOC} -o build/xxd xxd.go

yes.go:
	${GOC} -o build/yes yes.go
