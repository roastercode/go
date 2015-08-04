package main

import (
	"fmt"
	"time"
)

const (
	EVERY_DAY    = time.Duration(24) * time.Hour
	EVERY_HOUR   = time.Duration(1) * time.Hour
	EVERY_MINUTE = time.Duration(1) * time.Minute
	EVERY_SECOND = time.Duration(1) * time.Second
	EVERY_MILLI  = time.Duration(1) * time.Millisecond
	ONCE         = 0
)

var Schedules []Schedule
var reloadMutex = make(chan bool)

func init() {
	go func() { <-reloadMutex }() // skip the first reload check
	AddSchedule(Schedule{"stats", EVERY_SECOND, doStats})
	AddSchedule(Schedule{"scan", ONCE, doScan})

	Reload()
}

type Schedule struct {
	Name  string
	Every time.Duration
	Exec  func()
}

func doScan()  { fmt.Println("doing scan") }
func doStats() { fmt.Println("doing stats") }

func AddSchedule(s Schedule) { Schedules = append(Schedules, s) }

func Reload() {
	reloadMutex <- true
	for _, x := range Schedules {
		fmt.Println("Loading", x.Name)
		go func() {
			select {
			case <-reloadMutex:
				fmt.Println("Reloading schedules")
				return

			case <-time.NewTicker(x.Every).C:
				x.Exec()
			}
		}()
	}
}

func main() {
	<-time.NewTicker(time.Second * time.Duration(5)).C
	//http.ListenAndServe("127.0.0.1:9988", nil)
}
