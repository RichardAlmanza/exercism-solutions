package resistorcolorduo

var colorCodes map[string]uint8 = map[string]uint8{
	"black": 0,
	"brown": 1,
	"red": 2,
	"orange": 3,
	"yellow": 4,
	"green": 5,
	"blue": 6,
	"violet": 7,
	"grey": 8,
	"white": 9,
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	var value int
	var multiplier int = 1

	for _, color := range colors[:2] {
		value *= multiplier
		value += int(colorCodes[color])
		
		multiplier *= 10
	}

	return value
}
