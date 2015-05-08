package main

import "os"
import "fmt"
import "io/ioutil"
import "flag"

func safestring(bytes []byte) (string) {
    // This function replaces non printable characters with dots.
    s := make([]byte,len(bytes))
    for i:=0; i<len(bytes); i++ {
        if(bytes[i]>0x19 && bytes[i]<0x7F) {s[i]=bytes[i]} else {s[i]='.'}
    }
    return string(s)
}

func dump(bytes []byte,columns int) {
    if(len(bytes)%2==1) {bytes=append(bytes,0x00)} //Make sure bytes has an even amount of items
    w:=0         // Used to count how many words printed in each line
    space:=""    // Used to insert whitespace between words
    for i:=0; i<len(bytes); i++ {
        if(w>=columns) {
            fmt.Printf(" %s\n",safestring(bytes[i-(columns):i]));
            w=0;
        }
        if(w==0) { fmt.Printf("%07x: ",i) }
        if(i%2==1) {space=" "} else {space=""}          // Put a space every two bytes
        if(i==len(bytes)-1 && bytes[i]==0x00) {} else { // If the last byte is empty, don't print it
            fmt.Printf("%02x%s",bytes[i],space)
        }
        w++
    }
    if(w>0) { // If the last line is not full
        for i:=w; i<=columns; i++ {
            fmt.Printf("  ");
            if(i%2==1) {fmt.Printf(" ")}
        }
        fmt.Printf("  %s",safestring(bytes[len(bytes)-w:len(bytes)-1]))
    }
    fmt.Println()
}

func main() {
    cFlag := flag.Int("c",16,"columns")
    flag.Parse()
    if(len(flag.Args())==0) {
        bytes, _ := ioutil.ReadAll(os.Stdin)
        dump(bytes,*cFlag)
    } else {
        for j:=0; j<len(flag.Args()); j++ {
            bytes, _ := ioutil.ReadFile(flag.Arg(j))
            dump(bytes,*cFlag)
        }
    }
}
