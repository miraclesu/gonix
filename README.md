# gonix
Clones of the *nix tools written in go

# Status
If a program is marked complete, then most/all the standard features are implemented.
* base64 [works]
* basename [**complete**]
* cal [works, formatting needs work]
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
See also:
https://github.com/EricLagerg/go-coreutils

# Bugs
In general, expect there to be many bugs.
Many tools do not exist and many features are not implemented.

# Building
To keep the binary sizes small (around 20-30kb each, on my PC) I recommend using the gccgo compiler, but it is not required.

## Dependencies
* pflag, https://github.com/ogier/pflag

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

**WARNING:** Keep the programs in a directory **OUTSIDE** of your $PATH or they may end up being run *instead* of your system's commands.

# License
I am releasing this under the terms of the GNU GPLv3 license with **ABSOLUTELY NO WARRANTY OR LIABILITY**.
I have included the license in the file LICENSE.md

This is just one of my coding experiments and should not be used to actually replace your system's tools.
I am not liable for whatever these programs might do. Use them at your own risk.
