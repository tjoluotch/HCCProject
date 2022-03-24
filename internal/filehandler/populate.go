package filehandler

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

// populateData reads from the CSV file and populates the visitData type - see visit_data.go with any errors reporting
func populateData() (visitData, error) {
	dataV3 := make(visitData)

	visitsDataCSV, err := openFile()
	if err != nil {
		return nil, err
	}

	defer visitsDataCSV.Close()
	csvReader := csv.NewReader(visitsDataCSV)

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			log.Println("finished reading bathroom data csv")
			break
		}
		if err != nil {
			return nil, err
		}

		timeStamp := row[1]
		activity := row[3]

		// if activity for resident0001 type is a bathroom activity
		if strings.Contains(activity, bathroomValidator) {
			// parse timestamp to time.Time
			timeObject, err := time.Parse(time.RFC3339, timeStamp)
			if err != nil {
				return nil, err
			}

			// retreive date from timeObject
			year, month, day := timeObject.Date()
			dateAsString := strings.Join([]string{strconv.Itoa(day), strconv.Itoa(int(month)), strconv.Itoa(year)}, dateSeparator)

			// turn dateAsString into float64 type to use with LinearRegression func and key for map type
			dateArr := strings.Split(dateAsString, dateSeparator)
			dateStr := strings.Join(dateArr, "")
			dateAsFloat, err := strconv.ParseFloat(dateStr, 64)
			if err != nil {
				log.Fatal(err)
			}

			// if this is the 1st visit to the bathroom on a particular day, set the value to 1 and continue onto next iteration
			if dataV3[dateAsFloat] == nil {
				dataV3[dateAsFloat] = &data{}
				dataV3[dateAsFloat].bathroomVisitCount = 1
				dataV3[dateAsFloat].dateAsString = dateAsString
				continue
			}

			// bathroom visit has already happened on this day so iterate count and reassign new value
			totalVisitPerDay := dataV3[dateAsFloat].bathroomVisitCount
			totalVisitPerDay += 1
			dataV3[dateAsFloat].bathroomVisitCount = totalVisitPerDay
		}
	}
	return dataV3, nil
}
