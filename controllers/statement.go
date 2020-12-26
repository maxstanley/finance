package controllers

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/echo/v4"
)

type statement struct {
	AccountNumber int
	SortCode int
	Transactions []*transaction
}

type transaction struct {
	Date time.Time
	Description string
	Credit uint64
	Debit uint64
	Type string
	Essential bool
	PrimaryCategory string
	SecondaryCategory string
}

func RegisterStatementRoutes(group *echo.Group) {
	group.POST("/", createStatement)
}

func createStatement(c echo.Context) error {
	var err error
	// Identify the type of the file sent.
	accountNumber, err := strconv.Atoi(c.FormValue("accountNumber"))
	if err != nil {
		return err
	}

	sortCode, err := strconv.Atoi(c.FormValue("sortCode"))
	if err != nil {
		return err
	}

	file, err := c.FormFile("statement")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	s := statement{
		AccountNumber: accountNumber,
		SortCode: sortCode,
	}

	filename := strings.SplitAfter(file.Filename, ".")

	// Send the statement to the relevant parser.
	switch filename[len(filename)-1] {
	case "csv":
		err = parseLloydsStatement(&s, src)
	case "xlsx":
		err = parseSantanderStatement(&s, src)
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}

	// Return the parsed statement.
	return c.JSON(http.StatusOK, s)

	// FUTURE...

	// Save the parsed statements to the local Database.

	// Return Successful.
}

func parseLloydsStatement(s *statement, file io.Reader) error {
	transactions := []*transaction{}
	
	r := csv.NewReader(file)

	// Ignore the first line used for heading.
	r.Read()
	emptyLines := 0

	for {
		row, err := r.Read()

		if row == nil {
			emptyLines += 1
			if emptyLines > 3 {
				break
			}
			continue
		} else if emptyLines != 0 {
			emptyLines = 0
		}

		if err != nil {
			fmt.Println("Read Error")
			return err
		}

		t := transaction{}

		if t.Date, err = time.Parse("02/01/2006", row[0]); err != nil {
			fmt.Println("Date Parsing Error ", row[0])
			return err
		}

		t.Type = row[1]
		// row[2] SortCode
		// row[3] AccountNumber
		t.Description = row[4]

		if t.Debit, err = stringDecimalToInt(row[5]); err != nil {
			return err
		}
		if t.Credit, err = stringDecimalToInt(row[6]); err != nil {
			return err
		}

		// row[6] Balance

		transactions = append(transactions, &t)
	}

	s.Transactions = transactions

	return nil
}

func parseSantanderStatement(s *statement, file io.Reader) error {
	transactions := []*transaction{}
	
	fmt.Println("Open XLS File")
	xls, err := excelize.OpenReader(file)

	if err != nil {
		fmt.Println("Creating xls error")
		fmt.Println(err)
		return err
	}

	rows := xls.GetRows("Sheet1")

	index := 0

	for _, row := range rows {
		if index < 5 {
			index += 1
			continue
		}

		t := transaction{}

		// row[0] Blank
		if t.Date, err = time.Parse("02/01/2006", row[1]); err != nil {
			fmt.Println("Date Parsing Error ", row[1])
			return err
		}

		// row[2] Blank
		t.Description = row[3]
		// row[4] Blank

		if t.Credit, err = stringDecimalToInt(row[5]); err != nil {
			return err
		}
		if t.Debit, err = stringDecimalToInt(row[6]); err != nil {
			return err
		}

		// row[7] Balance

		transactions = append(transactions, &t)
	}

	s.Transactions = transactions

	return nil
}

func stringDecimalToInt(s string) (uint64, error) {
	var (
		i uint64
		err error
	)

	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, "£", "")

	if s == "" {
		return 0, nil
	}

	if i, err = strconv.ParseUint(s, 10, 64); err != nil {
		fmt.Println("Parsing Error", s)
		return 0, err
	} 

	return i, nil
}

