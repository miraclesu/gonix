# gonix
Clones of the *nix tools written in go

# Status
If a program is marked complete, then most/all the standard features are implemented.
* base64 [works]
* basename [works]
* cat [works]
* clear [**complete**]
* cp [works, very incomplete, not recursive]
* dirname [**complete**]
* echo [works]
* false [**complete**]
* head [works]
* ln [works]
* ls [doesn't compile for all of us]
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
* whoami [**complete**]
* xxd [works]
* yes [**complete**]

Unimplemented:
* more

# Misc.
Someone else had the same idea (of writing all the tools in Go) a while before mine.
https://github.com/EricLagerg/go-coreutils
His project is way more mature than mine and he said he'd love some help.


## Plans
This is just my little weekend project, so I won't be active much besides a few hours on the weekends when I have free time.
I will eventually lose interest in the project, probably in a month or two, but I'm doing this for fun anyway, not as a fulltime job or to ever actually replace anything.
Things I'm not planning on ever writing:
* a kernel
* a shell
* networking tools (ping, ip...)
* hardware tools (dd, du, df...)
* text editors (ed, vi, emacs...)
* small languages (bc, dc, sed, awk...)

and other complicated things. I'm focusing mainly on the text tools, some of the file tools (ls, cp, mkdir...), and a few other simple tools. Earlier I was thinking of doing some of those things in the list, but changed my mind after realizing the amount of work it would take. I want this to stay a simple little weekend project.

### Why publish something you aren't planning on finishing?
By putting my project on the Internet, I am getting valuable advice and criticism that I wouldn't get if I only worked on my own. I'm not making a real operating system here or commiting myself to a years long project, but this is educational for me and anyone else working on it, which is better than leaving it to rot on my hard drive.

#### Why do you keep talking about this anyway?
This project has gotten much more attention than I was expecting, and I don't want to mislead anyone on what my plans for the project are.

# Bugs
In general, expect there to be many bugs.
Many tools do not exist and many features are not implemented.
I used "flag" to do the flag-parsing, it is a little strict (ie. no combining flags like -xyz instead of -x -y -z).

# Building
To keep the binary sizes small (around 20-30kb each, on my PC) I recommend using the gccgo compiler, but it is not required.

Build the project:
```sh
$ make
```

Alternatively, build with gccgo:
```sh
$ make GCFLAGS="-compiler gccgo"
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
