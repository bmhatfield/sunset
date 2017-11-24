package main

import (
	"fmt"

	"github.com/bmhatfield/sunset/sunset"
)

func main() {
	s, err := sunset.Time()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Sunrise:", s.Sunrise.Local())
	fmt.Println("Sunset:", s.Sunset.Local())
	fmt.Println("Civil Twilight:", s.CivilTwilightEnd.Local())
	fmt.Println("Nautical Twilight:", s.NauticalTwilightEnd.Local())
}
