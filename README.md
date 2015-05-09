# gonix
Clones of the *nix tools written in go

# Status
If a program is marked complete, then most/all the standard features are implemented.
* base64 [works]
* basename [works]
* cat [works]
* cp [works, very incomplete, not recursive]
* dirname [**complete**]
* echo [works]
* false [**complete**]
* head [works]
* md5sum [works]
* mkdir [works]
* nl [works]
* pwd [works]
* rm [works, very incomplete, not recursive]
* seq [works]
* sha1sum [works]
* sleep [**complete**]
* tail [works inefficiently, needs to be rewritten like head]
* tee [**complete**]
* touch [works]
* true [**complete**]
* xxd [works]
* yes [**complete**]

Unimplemented:
* ls
* more

# Misc.
Someone else had the same idea (of writing all the tools in Go) a while before mine.
https://github.com/EricLagerg/go-coreutils
His project is way more mature than mine and he said he'd love some help.


## Plans
This is just my little weekend project, so I won't be active much besides a few hours on the weekends when I have free time.
I may eventually lose interest in the project, probably in a few months (it's happened to me before), but I'm doing this for fun anyway, not as a fulltime job.
Things I'm not planning on ever writing:
* a kernal
* a shell
* networking tools (ping, ip...)
* hardware tools (dd, du, df...)
* text editors (ed, vi, emacs...)
* small languages (bc, dc, sed, awk...)

and other complicated things. I'm focusing mainly on the text tools, some of the file tools (ls, cp, mkdir...), and a few other simple tools.

# Bugs
Many tools do not exist and many features are not implemented.
I used "flag" to do the flag-parsing, it is a little strict (ie. no combining flags like -xyz instead of -x -y -z).

# Building
To keep the binary sizes small (around 20-30kb each, on my PC) I use the gccgo compiler, so you need to have that installed.

Build the project:
```sh
$ make
```

It will compile each program and the binaries will be placed in the build directory.
```sh
$ cd build
```

# Running
As long as you are in the same directory as the compiled program, you can run it like so:

    ./x
Where x refers to the name of the program

**WARNING:** Keep the programs in a directory **NOT** in your $PATH or they may end up being run *instead* of your system's commands.

# License
I am releasing this under the terms of the GNU GPLv3 license with **ABSOLUTELY NO WARRANTY OR LIABILITY**.
I have included the license in the file LICENSE.md

This is just one of my coding experiments and should not be used to actually replace your system's tools.
I am not liable for whatever these programs might do. Use them at your own risk.
