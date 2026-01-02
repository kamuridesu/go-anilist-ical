package ical

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	EventStencil = `BEGIN:VEVENT
UID:%s
DTSTAMP:%s
DTSTART:%s
DTEND:%s
SUMMARY:%s
LOCATION:%s
DESCRIPTION:%s
LAST-MODIFIED:%s
END:VEVENT
`

	CalendarStencil = `BEGIN:VCALENDAR
PRODID:Kamu-anilist-scheduler
VERSION:%s
CALSCALE:%s
%s
END:VCALENDAR`
)

type VCalendar struct {
	version  string
	calscale string
	events   []*VEvent
}

type VEvent struct {
	uid          string
	dtStamp      string
	dtStart      string
	dtEnd        string
	summary      string
	location     string
	description  string
	lastModified string
}

func New() *VCalendar {
	return &VCalendar{
		version:  "2.0",
		calscale: "GREGORIAN",
	}
}

func (vc *VCalendar) addEvent(e *VEvent) {
	vc.events = append(vc.events, e)
}

func (vc *VCalendar) AddEvent(start, end time.Time, summary, description string) {
	ev := VEvent{
		uid:          uuid.NewString(),
		dtStart:      start.UTC().Format("20060102T150405Z"),
		dtEnd:        end.UTC().Format("20060102T150405Z"),
		dtStamp:      time.Now().UTC().Format("20060102T150405Z"),
		summary:      summary,
		description:  description,
		location:     "anilist",
		lastModified: time.Now().UTC().Format("20060102T150405Z"),
	}
	vc.addEvent(&ev)
}

func (vc *VCalendar) generateEvents() string {
	evString := ""
	for _, ev := range vc.events {
		evString += fmt.Sprintf(
			EventStencil,
			ev.uid,
			ev.dtStamp,
			ev.dtStart,
			ev.dtEnd,
			ev.summary,
			ev.location,
			ev.description,
			ev.lastModified,
		)
	}
	return strings.TrimSpace(evString)
}

func (vc *VCalendar) Build() string {
	return fmt.Sprintf(
		CalendarStencil,
		vc.version,
		vc.calscale,
		vc.generateEvents(),
	)
}
