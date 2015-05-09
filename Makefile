# Targets:
# all: Build code.
# clean: Remove build artifacts.
# $FILENAME: Build just $FILENAME.
# build: Create build directory.

GOC=go build -compiler gccgo
GOFLAGS=

.PHONY: all

all: build base64 basename cat cp dirname echo false head mkdir nl pwd seq sleep tail tee touch true xxd yes

.PHONY: clean

clean:
	rm -rf build

build:
	mkdir -p build

base64:
	${GOC} -o build/base64 base64.go

basename:
	${GOC}  -o build/basename basename.go

cat:
	${GOC} -o build/cat cat.go

cp:
	${GOC} -o build/cp cp.go

dirname:
	${GOC} -o build/dirname dirname.go

echo:
	${GOC} -o build/echo echo.go

false:
	${GOC} -o build/false false.go

head:
	${GOC} -o build/head head.go

mkdir:
	${GOC} -o build/mkdir mkdir.go

nl:
	${GOC} -o build/nl nl.go

pwd:
	${GOC} -o build/pwd pwd.go

seq:
	${GOC} -o build/seq seq.go

sleep:
	${GOC} -o build/sleep sleep.go

tail:
	${GOC} -o build/tail tail.go

tee:
	${GOC} -o build/tee tee.go

touch:
	${GOC} -o build/touch touch.go

true:
	${GOC} -o build/true true.go

xxd:
	${GOC} -o build/xxd xxd.go

yes:
	${GOC} -o build/yes yes.go
