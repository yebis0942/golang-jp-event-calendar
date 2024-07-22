package gojpcal

import (
	"fmt"
	"time"

	ics "github.com/arran4/golang-ical"
)

func BuildICalendar(events []*Event) (string, error) {
	cal := ics.NewCalendar()
	cal.SetProductId(fmt.Sprintf("-//yebis0942.github.io//%s//JA", cmdName))
	cal.SetLastModified(time.Now())
	for _, event := range events {
		e := cal.AddEvent(fmt.Sprintf("golang-jp-event-calendar@%s", event.Link.URL))
		e.SetStartAt(event.Start)
		e.SetEndAt(event.End)
		e.SetSummary(event.Title)
		e.SetURL(event.Link.URL)
		e.SetDescription(fmt.Sprintf("%s\n%s", event.Group.Name, event.Link.URL))
	}
	return cal.Serialize(), nil
}
