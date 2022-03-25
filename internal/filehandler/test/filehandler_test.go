package filehandler_test

import (
	"HCCProject/internal/filehandler"
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("FileHandler test opening of CSV file", func() {
	var (
		fileHandlerObj      filehandler.FileHandler
		visitData           filehandler.VisitData
		file                *os.File
		err                 error
		randomFileName      string
		correctFileName     string
		fileIncorrectFormat string
	)

	BeforeEach(func() {
		randomFileName = "xyzabc.csv"
		fileIncorrectFormat = "ports.json"
		correctFileName = "BATHROOM_VISIT_DATA.csv"
	})

	Context("We are opening a file", func() {
		Context("if file is not present", func() {
			BeforeEach(func() {
				file, err = filehandler.OpenFile(randomFileName)
			})
			It("file is not present", func() {
				Expect(file).To(BeNil())
				Expect(err).To(HaveOccurred())
			})
		})

		Context("if file is not the correct format", func() {
			BeforeEach(func() {
				file, err = filehandler.OpenFile(fileIncorrectFormat)
			})
			It("file has incorrect format", func() {
				Expect(file).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).Should(MatchError(errors.New("file format incorrect - expecting .csv")))
			})
		})

		Context("if file is opened successfully", func() {
			BeforeEach(func() {
				file, err = filehandler.OpenFile(correctFileName)
			})
			It("file opened successfully", func() {
				Expect(file).ToNot(BeNil())
				Expect(err).To(BeNil())
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})

	Context("we are populating data", func() {
		Context("populated data successfully", func() {
			BeforeEach(func() {
				visitData, err = filehandler.PopulateData()
			})
			It("populated data successfully", func() {
				Expect(err).To(BeNil())
				Expect(err).ToNot(HaveOccurred())
				Expect(visitData).ToNot(BeEmpty())
			})
		})
	})

	Context("we are calculating the trend line for each day given the amount of visits", func() {
		Context("calculated trend data successfully", func() {
			BeforeEach(func() {
				fileHandlerObj = filehandler.NewFileHandler()
				visitData, err = fileHandlerObj.CalculateTrendDataPoints()
			})
			It("calculated trend data & populated collection within filehandler", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(err).To(BeNil())
				Expect(visitData).ToNot(BeEmpty())
			})
		})
	})
})
