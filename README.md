# gonix
Clones of the *nix tools written in go

# Bugs
Many tools do not exist and many features are not implemented.
I used "flag" to do the flag-parsing, it is a little strict (ie. no combining flags like -xyz instead of -x -y -z).

# Building
To keep the binary sizes small (around 20-30kb each, on my PC) I use the gccgo compiler.
You can build each tool like so:

    go build -compiler gccgo x.go
where x.go refers to whichever source file you want to compile (eg. cat.go or yes.go)
When it is done building the source file it will place a binary file in the same location.
The binary file will have the same name as the source file but without the .go (eg. cat.go becomes cat)

# Running
As long as you are in the same directory as the compiled program, you can run it like so:

    ./x
Where x refers to the name of the program

**WARNING:** Keep the binaries in a directory that is NOT on your $PATH or they may end up being run instead of your system's commands.

# License
I am releasing this under the terms of the GNU GPLv3 license with **ABSOLUTELY NO WARRANTY OR LIABILITY**.
I have included the license in the file LICENSE.md

This is just one of my coding experiments and should not be used to actually replace your system's tools.
I am not liable for whatever these programs might do. Use them at your own risk.
