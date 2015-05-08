# gonix
Clones of the *nix tools written in go

# Status
If a program is marked complete, then most/all the standard features are implemented.
* basename [works]
* cat [works]
* cp [works, very incomplete]
* dirname [**complete**]
* echo [works]
* false [**complete**]
* head [works]
* mkdir [works]
* pwd [works]
* seq [works]
* sleep [**complete**]
* tail [works]
* tee [**complete**]
* true [**complete**]
* xxd [works]
* yes [**complete**]

Unimplemented:
* bc
* dc
* ed
* ls
* more
* rm

# Bugs
Many tools do not exist and many features are not implemented.
I used "flag" to do the flag-parsing, it is a little strict (ie. no combining flags like -xyz instead of -x -y -z).

# Building
To keep the binary sizes small (around 20-30kb each, on my PC) I use the gccgo compiler, so you need to have that installed.

I haven't wrote a Makefile yet, but in the meantime I am using a very simple build script.

You *must* be in the same directory that the source code is in when running the build script.

To run it:

    sh build.sh

It will compile each program and the binaries will be placed in the same directory.

**TODO:** Write a Makefile.

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
