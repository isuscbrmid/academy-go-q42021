package repositories

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"go-capstone/common"
	"go-capstone/entities"
)

func GetEmployee(ID int64) (entities.User, common.AppError) {
	csvFile, csvOpenerErr := os.Open("./app/mock-data/data.csv")
	if csvOpenerErr != nil {
		return entities.User{}, common.OpenCsvError
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, csvReaderErr := csv.NewReader(csvFile).ReadAll()
	if csvReaderErr != nil {
		fmt.Println(csvReaderErr)
	}
	for _, line := range csvLines {
		if line[0] != "ID" {
			id, parseErr := strconv.ParseInt(line[0], 10, 64)
			if parseErr == nil && ID == id {
				user := entities.User{
					Id:    id,
					Name:  line[1],
					Email: line[2],
				}
				return user, common.AppError{}
			}
			if parseErr != nil {
				log.Fatalln(parseErr)
				return entities.User{}, common.CustomError
			}
		}
	}
	return entities.User{}, common.NotFoundError
}
