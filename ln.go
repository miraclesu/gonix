package main

import "os"
import flag "github.com/ogier/pflag"

func main() {
    sFlag := flag.BoolP("symbolic", "s", false, "symlink mode")
    flag.Parse();
    if(len(flag.Args())>1) {
        if(*sFlag) {
            os.Symlink(flag.Arg(0) , flag.Arg(1))
        } else {
            os.Link(flag.Arg(0) , flag.Arg(1))
        }
    }
}
