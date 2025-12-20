package intermediate

import (
	"fmt"
	"time"
)

func main() {
	// Current local time
	fmt.Println(time.Now())

	// Specific time
	specificTime := time.Date(2025, time.December, 18, 20, 7, 0, 0, time.UTC)
	fmt.Println("Specific Time: ", specificTime)

	// Parse time
	parsedTime, _ := time.Parse("2006-01-02", "2025-12-18")
	fmt.Println("Parse Time: ", parsedTime)

	// Formatting time
	t := time.Now()
	fmt.Println("Formatted Time: ", t.Format("02-01-2006"))

	// Lets manipulate time
	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println("Added One Day To Today: ", oneDayLater)

	fmt.Println("Rounded Time: ", t.Round(time.Hour))

	loc, _ := time.LoadLocation("Asia/Kolkata")
	t = time.Now()

	// Convert this to the specific time
	tLocal := t.In(loc)

	// Perform rounding
	roundedTime := t.Round(time.Hour)
	roundedTimeLocal := roundedTime.In(loc)

	fmt.Println("Original Time (UTC):", t)
	fmt.Println("Original Time (Local):", tLocal)
	fmt.Println("Rounded Time (UTC):", roundedTime)
	fmt.Println("Rounded Time (Local):", roundedTimeLocal)

	// Truncate time always round down when given the time
	fmt.Println("Truncated Time: ", t.Truncate(time.Hour))

	// Lets make couple more time zone example
	loc, _ = time.LoadLocation("America/New_York")
	tInNy := time.Now().In(loc)
	fmt.Println("New York Time: ", tInNy)

	t1 := time.Now()
	turkeyLoc, _ := time.LoadLocation("Europe/Istanbul")
	t2 := time.Date(1983, time.February, 5, 3, 0, 0, 0, turkeyLoc)
	duration := t1.Sub(t2)

	fmt.Println("\n--- Duration Formatting ---")
	fmt.Println("Time Passed since my birth:")

	// Calculate total duration components
	days := int(duration.Hours() / 24)
	years := days / 365
	remainingDays := days % 365
	months := remainingDays / 30
	remainingDays = remainingDays % 30

	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	// Format and print
	fmt.Printf("Duration: %d years, %d months, %d days - %02d:%02d (hours:minutes)\n",
		years, months, remainingDays, hours, minutes)

	// Compare Times
	fmt.Println("t2 is after t1?", t2.After(t1))
	fmt.Println("t2 before t1? ", t2.Before(t1))
}
