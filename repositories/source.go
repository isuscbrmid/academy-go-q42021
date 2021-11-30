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
	csvFile, csvOpenerErr := os.Open("./app/mock-data/users.csv")
	if csvOpenerErr != nil {
		return entities.User{}, common.OpenCsvError
	}
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

func SaveUsers(users []entities.User) ([]entities.User, common.AppError) {
	csvFile, csvOpenerErr := os.OpenFile("./app/mock-data/users.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if csvOpenerErr != nil {
		fmt.Println("asfas", csvOpenerErr)
		return users, common.OpenCsvError
	}

	// if csvReaderErr != nil {
	// 	fmt.Println(csvReaderErr)
	// 	fmt.Println("a1111", csvOpenerErr)
	// 	return users, common.CustomError
	// }

	csvLength := 200 // assing id properly
	fmt.Println("lenght of file", csvLength)

	csvwriter := csv.NewWriter(csvFile)

	for i, user := range users {
		fmt.Println(strconv.FormatInt(int64(i+csvLength), 10))
		var hdr = []string{strconv.FormatInt(int64(i+csvLength), 10), user.Name, user.Email}
		returnError := csvwriter.Write(hdr)
		if returnError != nil {
			return users, common.CustomError
		}
	}
	csvwriter.Flush()

	return users, common.AppError{}
}
