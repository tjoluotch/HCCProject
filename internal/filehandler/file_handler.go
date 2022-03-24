package filehandler

import (
	"gonum.org/v1/gonum/stat"
	"log"
)

type FileHandler struct {
	collection *visitData
}

// NewFileHandler returns an empty FileHandler type without any visits data
func NewFileHandler() FileHandler { return FileHandler{collection: nil} }

// CalculateTrendDataPoints runs a linear regression algorithm on bathroom visits per day and assigns the
// necessary data to the FileHandler collection field
func (fh FileHandler) CalculateTrendDataPoints() error {

	visitsData, err := populateData()
	if err != nil {
		return err
	}

	var visitsSlice []float64
	var dateSlice []float64

	for dateAsFloat, obj := range visitsData {
		// append to float slice for use with linear regression func
		dateSlice = append(dateSlice, dateAsFloat)
		visitsSlice = append(visitsSlice, obj.bathroomVisitCount)
		log.Println("date - ", obj.dateAsString, "visits -", obj.bathroomVisitCount)
	}

	// run linear regression algorithm from gonum package to get trend line equation
	// y = ax + b
	bValue, aValue := stat.LinearRegression(dateSlice, visitsSlice, nil, false)
	log.Printf("y = %v(x) + %v\n", aValue, bValue)

	// get y value (trend line data point) for each date
	for dateAsFloat, obj := range visitsData {
		ax := aValue * dateAsFloat
		y := ax + bValue
		// encode data structure with trend output point for particular date
		obj.linRegTrend = y
		// encode data structure with boolean evaluation: is bathroomVisitCount > than trend at given date
		obj.aboveTrend = y < obj.bathroomVisitCount
		log.Printf("%+v\n", obj)
	}

	fh.collection = &visitsData
	return nil
}
