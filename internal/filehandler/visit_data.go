package filehandler

// VisitData maps a date in float64 format to a pointer of type data struct
type VisitData map[float64]*data

// data contains fields that are used to return output necessary for client application visualisation
type data struct {
	// date represented with "|" separator between year, month, day
	dateAsString string
	// total visits to the bathroom for a particular day
	bathroomVisitCount float64
	// trend value - y axis for given date
	linRegTrend float64
	// boolean indicating whether bathroomVisitCount is above or below trend for given date
	aboveTrend bool
}
