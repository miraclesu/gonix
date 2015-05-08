package main

import "os"
import "io"

func main() {
	//sf: source file, df: destination file
	sf, _ := os.Open(os.Args[1])
	sfinfo, _ := os.Stat(os.Args[1])
	sfmode := sfinfo.Mode()
	df, _ := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, sfmode)
	io.Copy(df, sf)
	sf.Close()
	df.Close()
}
