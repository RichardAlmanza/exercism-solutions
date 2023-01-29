package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalCount int

	for i := 0; i < len(birdsPerDay); i++ {
		totalCount += birdsPerDay[i]
	}

	return totalCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var totalCount int
	var startingDay int
	var finishingDay int

	startingDay = (week - 1) * 7
	finishingDay = startingDay + 7

	for i := startingDay; i < finishingDay; i++ {
		totalCount += birdsPerDay[i]
	}

	return totalCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i % 2 == 0 {
			birdsPerDay[i]++
		}
	}

	return birdsPerDay
}
