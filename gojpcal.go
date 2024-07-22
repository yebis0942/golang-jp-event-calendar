package gojpcal

import (
	"flag"
	"fmt"
	"log"
	"time"
)

const cmdName = "gojpcal"

// time.LoadLocation("Asia/Tokyo")
var tzAsiaTokyo *time.Location

func init() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	tzAsiaTokyo = loc
}

func Run(argv []string) error {
	config := LoadConfig()

	fs := flag.NewFlagSet(cmdName, flag.ContinueOnError)
	var sinceStr = fs.String("since", "", "2006-01-02")
	if err := fs.Parse(argv); err != nil {
		return err
	}
	since, err := time.ParseInLocation("2006-01-02", *sinceStr, tzAsiaTokyo)
	if err != nil {
		return err
	}

	events := []*Event{}
	for i, groupID := range config.ConnpassGroups {
		log.Printf("fetching: %s", groupID)

		feedBytes, err := FetchGroupEventFeed(groupID)
		if err != nil {
			return err
		}
		group, err := ParseGroupEventFeed(feedBytes)
		if err != nil {
			return err
		}
		for _, event := range group.Events {
			if since.After(event.Start) {
				continue
			}
			events = append(events, event)
		}

		if i != len(config.ConnpassGroups)-1 {
			time.Sleep(1 * time.Second)
		}
	}

	iCal, err := BuildICalendar(events)
	if err != nil {
		return err
	}
	fmt.Println(iCal)

	return nil
}
