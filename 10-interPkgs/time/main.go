package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2019, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	then = then.Add(-time.Hour * 24) // Subtract a day
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	p("Then es antes de Now: ", then.Before(now))
	p("Then es después de Now: ", then.After(now))
	p("Then es igual a Now: ", then.Equal(now))

	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
}
