package filehandler

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"log"
)

// FileHandler type contains bathroom visits data
type FileHandler struct {
}

// NewFileHandler returns an empty FileHandler type without any visits data
func NewFileHandler() FileHandler { return FileHandler{} }

// CalculateTrendDataPoints runs a linear regression algorithm on bathroom visits per day and assigns the
// necessary data to the FileHandler collection field
func (fh FileHandler) CalculateTrendDataPoints() (VisitData, error) {

	visitsData, err := PopulateData()
	if err != nil {
		return nil, err
	}

	var visitsSlice []float64
	var dateSlice []float64

	for dateAsFloat, obj := range visitsData {
		// append to float slice for use with linear regression func
		dateSlice = append(dateSlice, dateAsFloat)
		visitsSlice = append(visitsSlice, obj.BathroomVisitCount)
		log.Println("date - ", obj.DateAsString, "visits -", obj.BathroomVisitCount)
	}

	// run linear regression algorithm from gonum package to get trend line equation
	// y = ax + b
	bValue, aValue := stat.LinearRegression(dateSlice, visitsSlice, nil, false)
	trendEquation := fmt.Sprintf("y = %v(x) + %v", aValue, bValue)
	log.Printf("%s\n", trendEquation)

	// get y value (trend line data point) for each date
	for dateAsFloat, obj := range visitsData {
		ax := aValue * dateAsFloat
		y := ax + bValue
		// encode data structure with trend output point for particular date
		obj.LinRegTrend = y
		// encode data structure with boolean evaluation: is bathroomVisitCount > than trend at given date
		obj.AboveTrend = y < obj.BathroomVisitCount
		// encode data structure with trend line - same for all Data objects see stat.LinearRegression in gonum 3rd party package
		obj.Equation = trendEquation
		log.Printf("%+v\n", obj)
	}

	return visitsData, nil
}
