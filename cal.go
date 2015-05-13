package main

import "fmt"
import flag "github.com/ogier/pflag"
import "strconv"
import "log"
import "time"

func leapyear(year int) (int) {
	//Return 1 if leapyear, 0 if not
	if year%4==0 && (year%100!=0 || year%400==0) {
		return 1
	}
	return 0
}

func monthlen(month int, year int) (int) {
	//Return length of month in days
	switch(month) {
		case  1: return 31
		case  2: return 28+leapyear(year)
		case  3: return 31
		case  4: return 30
		case  5: return 31
		case  6: return 30
		case  7: return 31
		case  8: return 31
		case  9: return 30
		case 10: return 31
		case 11: return 30
		case 12: return 31		
	}
	return 0
}

func calendar(month int, year int) {
	weekday := int(time.Date(year,time.Month(month),1,0,0,0,0,time.UTC).Weekday())
	fmt.Printf("%s %d\n",time.Month(month).String(),year)
	fmt.Printf("Su Mo Tu We Th Fr Sa\n")
	for i:=0; i<weekday; i++ {fmt.Printf("   ")}
	for day:=1; day<=monthlen(month,year); day++ {
		if weekday==6 {
			fmt.Printf("%2d\n",day)
			weekday=0
		} else {
			fmt.Printf("%2d ",day)
			weekday++
		}
	}
	if weekday!=6 {fmt.Printf("\n")}
}

func main() {
	flag.Parse()
	if len(flag.Args())==0 {
		year := int(time.Now().Year())
		month := int(time.Now().Month())
		calendar(month,year)
	} else if len(flag.Args())==1 {
		year, err := strconv.Atoi(flag.Arg(0))
		if err!=nil {log.Fatal(err)}
		for month:=1; month<=12; month++ {
			calendar(month,year)
		}
	} else if len(flag.Args())==2 {
		month, err := strconv.Atoi(flag.Arg(0))
		if err!=nil {log.Fatal(err)}
		year, err := strconv.Atoi(flag.Arg(1))
		if err!=nil {log.Fatal(err)}
		calendar(month,year)
		fmt.Println()
	}
}
