package raindrops

import (
	"strconv"
	"strings"
)

// Convert returns the corresponding sound for number of raindrops
func Convert(number int) string {
	// Yes, I'm going to kill a fly with a snipe
	var result string
	var stringBuilder strings.Builder
	factors := [3]int{3, 5, 7}
	sounds := [3]string{"Pling", "Plang", "Plong"}

	for i, factor := range factors {
		if number % factor == 0 {
			stringBuilder.WriteString(sounds[i])
		}
	}

	result = stringBuilder.String()

	if result == "" {
		result = strconv.Itoa(number)
	}

	return result
}
