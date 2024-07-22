package gojpcal

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type Group struct {
	Name   string   `xml:"title"`
	Events []*Event `xml:"entry"`
}

type Event struct {
	Group   *Group
	Title   string `xml:"title"`
	Summary string `xml:"summary"`
	Link    struct {
		URL string `xml:"href,attr"`
	} `xml:"link"`
	Start time.Time
	End   time.Time
}

// 複数日にわたるイベントはサポートしない
var eventTimePattern = regexp.MustCompile(`(\d{4}/\d{2}/\d{2} \d{2}:\d{2} ～ \d{2}:\d{2})`)

func FetchGroupEventFeed(connpassGroup string) ([]byte, error) {
	resp, err := http.Get("https://" + connpassGroup + ".connpass.com/ja.atom")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func ParseGroupEventFeed(rawFeed []byte) (*Group, error) {
	var group Group
	if err := xml.Unmarshal(rawFeed, &group); err != nil {
		return nil, err
	}
	group.Name = strings.TrimSuffix(group.Name, " グループの新着イベント")

	for _, event := range group.Events {
		event.Group = &group

		eventTimeStr := eventTimePattern.FindString(event.Summary)
		if eventTimeStr == "" {
			log.Printf("%q: skip サポートしていない形式の開催日時です", event.Title)
			continue
		}

		start, end, err := parseEventDate(eventTimeStr)
		if err != nil {
			return nil, err
		}
		event.Start = start
		event.End = end

		// URLからクエリパラメーターを削除する
		parsedURL, err := url.Parse(event.Link.URL)
		if err != nil {
			return nil, err
		}
		parsedURL.RawQuery = ""
		event.Link.URL = parsedURL.String()
	}

	return &group, nil
}

func parseEventDate(s string) (time.Time, time.Time, error) {
	parts := strings.Split(s, " ")
	layout := "2006/01/02 15:04"
	startTime, err := time.ParseInLocation(layout, parts[0]+" "+parts[1], tzAsiaTokyo)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	endTime, err := time.ParseInLocation(layout, parts[0]+" "+parts[3], tzAsiaTokyo)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return startTime, endTime, nil
}
