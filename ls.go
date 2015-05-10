package main

import "os"
import "fmt"
import "io/ioutil"
import "strings"
import "os/exec"
import "math"
import "strconv"
import "runtime"
import "syscall"
import "text/tabwriter"
import "flag"
import "os/user"
import "time"

/*
#include <grp.h>
#include <unistd.h>
#include <stdlib.h>
*/
import "C"

const groupBuffer = iota

const listTemplate = "%s\u0020\u0020%s\u0020%s\u0020\u0020%s\u0020\u0020%s\u0020%s\u0020%s\u0020%s\u0020%s\n"

var tabLen = 8

var minItemLen = 4

var showAll = flag.Bool("a", false, "Show hidden files")
var formatAsListFlag = flag.Bool("l", false, "Format as list")

// used as a shortcut for `-l -a`
var laCombined = flag.Bool("la", false, "Format as list and show all")

func getTermWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	splitted := strings.Split(strings.TrimSpace(string(out)), " ")

	width, _ := strconv.Atoi(splitted[1])

	return width
}

func formatAsTable(files []os.FileInfo) {
	w := new(tabwriter.Writer)
	maxLineLen := getTermWidth()

	additional := 0

	var fnames []string

	for _, file := range files {
		fnames = append(fnames, file.Name())
	}

	for _, f := range fnames {
		if len(f) < minItemLen {
			additional += minItemLen - len(f)
		}
	}

	ratio := float64(len(strings.Join(fnames, ""))+(len(fnames)*tabLen)+additional) / float64(maxLineLen)

	lines := int(math.Ceil(ratio))

	var nameArrs [][]string
	nameArrs = make([][]string, lines)

	for i, f := range fnames {
		index := i % lines
		nameArrs[index] = append(nameArrs[index], f)
	}

	w.Init(os.Stdout, 0, tabLen, minItemLen, '\t', 0)
	for _, arr := range nameArrs {
		fmt.Fprintln(w, strings.Join(arr, "\t"))
	}

	w.Flush()
}

func lookupGroupId(gid uint32) (string, error) {
	var grp C.struct_group
	var result *C.struct_group
	var bufSize C.long
	if runtime.GOOS == "freebsd" {
		// FreeBSD doesn't have _SC_GETPW_R_SIZE_MAX
		// or SC_GETGR_R_SIZE_MAX and just returns -1.
		// So just use the same size that Linux returns
		bufSize = 1024
	} else {
		size := C.int(C._SC_GETGR_R_SIZE_MAX)
		constName := "_SC_GETGR_R_SIZE_MAX"
		bufSize = C.sysconf(size)
		if bufSize <= 0 || bufSize > 1<<20 {
			return "", fmt.Errorf("Failed to allocate memory for %s of %d", constName, bufSize)
		}
	}

	buf := C.malloc(C.size_t(bufSize))

	defer C.free(buf)

	rv := C.getgrgid_r(C.gid_t(gid),
		&grp,
		(*C.char)(buf),
		C.size_t(bufSize),
		&result)
	if rv != 0 {
		return "", fmt.Errorf("Group lookup failed %d", gid)
	}
	if result == nil {
		return "", fmt.Errorf("err")
	}
	return C.GoString(grp.gr_name), nil
}

func formatAsList(files []os.FileInfo) {
	largestFileSize := int64(0)
	mostLinksToFile := uint16(0)
	areThereDoubleDigitDays := false

	for _, file := range files {
		if file.Size() > largestFileSize {
			largestFileSize = file.Size()
		}
		stat := file.Sys().(*syscall.Stat_t)

		if stat.Nlink > mostLinksToFile {
			mostLinksToFile = stat.Nlink
		}

		if file.ModTime().Day() > 9 {
			areThereDoubleDigitDays = true
		}

	}

	fileSizeLen := len(fmt.Sprintf("%v", largestFileSize))
	fileLinkLen := len(fmt.Sprintf("%v", mostLinksToFile))

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 4, 4, 4, '\u0020', 0)
	for _, file := range files {

		fileSize := fmt.Sprintf("%v", file.Size())
		stat := file.Sys().(*syscall.Stat_t)
		numLinks := fmt.Sprintf("%v", stat.Nlink)
		user, _ := user.LookupId(fmt.Sprintf("%v", stat.Uid))
		username := user.Username
		group, _ := lookupGroupId(stat.Gid)
		day := fmt.Sprintf("%v", file.ModTime().Day())
		month := file.ModTime().Month().String()[:3]

		var hour = ""

		if file.ModTime().Hour() < 9 {
			hour = fmt.Sprintf("0%v", file.ModTime().Hour())
		} else {
			hour = fmt.Sprintf("%v", file.ModTime().Hour())
		}

		var minute = ""

		if file.ModTime().Minute() < 9 {
			minute = fmt.Sprintf("0%v", file.ModTime().Minute())
		} else {
			minute = fmt.Sprintf("%v", file.ModTime().Minute())
		}

		hours := fmt.Sprintf("%v:%v", hour, minute)
		year := file.ModTime().Year()
		var DateColumn3 = ""

		if len(fileSize) < fileSizeLen {
			fileSize = strings.Repeat(" ", fileSizeLen-len(fileSize)) + fileSize
		}

		if len(numLinks) < fileLinkLen {
			numLinks = strings.Repeat(" ", fileLinkLen-len(numLinks)) + numLinks
		}

		if areThereDoubleDigitDays && len(day) < 2 {
			day = " " + day
		}

		if year != time.Now().Year() {
			DateColumn3 = fmt.Sprintf(" %v", year)
		} else {
			DateColumn3 = hours
		}

		fmt.Fprintf(w, listTemplate, file.Mode(), numLinks, username, group, fileSize, month, day, DateColumn3, file.Name())
	}
	w.Flush()
}

func main() {
	flag.Parse()

	if *laCombined {
		*showAll = *laCombined
		*formatAsListFlag = *laCombined
	}

	args := flag.Args()
	var dir string
	var filteredFiles []os.FileInfo

	if len(args) < 1 {
		dir = "./"
	} else {
		dir = args[0]
	}

	files, _ := ioutil.ReadDir(dir)

	for _, f := range files {
		isHidden := strings.HasPrefix(f.Name(), ".")

		if (*showAll || !isHidden) && f.Name() != "" {
			filteredFiles = append(filteredFiles, f)
		}
	}

	if *showAll {
		parent, _ := os.Lstat(dir + "../")
		thisDir, _ := os.Lstat(dir + "./")
		filteredFiles = append([]os.FileInfo{thisDir, parent}, filteredFiles...)
	}

	if !*formatAsListFlag {
		formatAsTable(filteredFiles)
	} else {
		formatAsList(filteredFiles)
	}
}
