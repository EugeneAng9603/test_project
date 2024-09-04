package main

import (
	"fmt"

	"github.com/EugeneAng9603/test_project/csv"
	"github.com/EugeneAng9603/test_project/mongodb"
	"github.com/EugeneAng9603/test_project/mysql"
)

type DataSaver interface {
	Save(data string) error
}

func saveData(saver DataSaver, data string) {
	if err := saver.Save(data); err != nil {
		fmt.Println("Error saving data:", err)
	} else {
		fmt.Println("Data saved successfully")
	}
}

func main() {
	mysqlSaver, err := mysql.NewMySQLSaver("user:password@/dbname")
	if err != nil {
		fmt.Println("MySQL error:", err)
		return
	}
	saveData(mysqlSaver, "example data")

	// Example for MongoDB
	mongoSaver, err := mongodb.NewMongoDBSaver("mongodb://localhost:27017", "mydb", "mycoll")
	if err != nil {
		fmt.Println("MongoDB error:", err)
		return
	}
	saveData(mongoSaver, "example data")

	// Example for CSV
	csvSaver, err := csv.NewCSVFileSaver("data.csv")
	if err != nil {
		fmt.Println("CSV file error:", err)
		return
	}
	defer csvSaver.file.Close()
	defer csvSaver.writer.Flush()
	saveData(csvSaver, "example data")
}
