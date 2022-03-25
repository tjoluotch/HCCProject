package filehandler

// VisitData maps a date in float64 format to a pointer of type data struct
type VisitData map[float64]*Data

// data contains fields that are used to return output necessary for client application visualisation
type Data struct {
	// date represented with "|" separator between year, month, day
	DateAsString string `json:"date_as_string"`
	// total visits to the bathroom for a particular day
	BathroomVisitCount float64 `json:"bathroom_visit_count"`
	// trend value - y axis for given date
	LinRegTrend float64 `json:"lin_reg_trend"`
	// boolean indicating whether bathroomVisitCount is above or below trend for given date
	AboveTrend bool `json:"above_trend"`
	// string representing equation that calculated trend line - same for all Data objects. format: y = a(x) + b
	Equation string `json:"equation"`
}
