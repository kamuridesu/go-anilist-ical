# go-anilist-ical

Generate ICalendar files for your Anilist current and planning lists

## Usage

Compile and run the program. It'll start to listen on http://localhost:8080, then just send a request like: http://localhost:8080/username.ics.

If you want to also get the planning schedule, just add `?iinclude_planning=true` to the end of your request.
