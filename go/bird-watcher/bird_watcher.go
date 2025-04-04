package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	birdSum := 0
	for _, day := range birdsPerDay {
		birdSum += day
	}
	return birdSum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birdsInWeek := 0
	weekStart := 7 * (week - 1)
	weekEnd := weekStart + 7

	for i := weekStart; i < weekEnd && i < len(birdsPerDay); i++ {
		birdsInWeek += birdsPerDay[i]
	}
	return birdsInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}
