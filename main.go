// Package to obtain data from various sources and formats

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	_ "github.com/go-sql-driver/mysql"
	//"log"
	"os"
)

type dataSource interface {
	get() (dataframe.DataFrame, error)
}

type CSV struct {
	name         string
	docLocal     string
	isThirdParty bool
}

// todo edit so that it also takes a query as a type or struct or func
type DB struct {
	dbType  string
	dbLogin string
}

//type logData struct {
//
//}

func (f CSV) get() (dataframe.DataFrame, error) {
	csvFile, err := os.Open(f.docLocal)
	reader := bufio.NewReader(csvFile)
	df := dataframe.ReadCSV(reader)
	return df, err
}

// Get function to obtain database info from a query
func (d DB) get() (*sql.Rows, error) {
	// https://github.com/go-sql-driver/mysql/wiki/Examples
	// https://github.com/golang/go/wiki/SQLDrivers
	db, err := sql.Open(d.dbType, d.dbLogin)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT country.Name FROM world.country")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	return rows, err
}

//func (l logData) get() (dataframe.DataFrame, error) {
//
//}

func getData(s dataSource) dataframe.DataFrame {
	df, _ := s.get()
	return df
}

func main() {
	//dt := CSV{name: "csv", docLocal: "df_test.csv", isThirdParty: false}
	//df, err := dt.get()
	//if err != nil {
	//	fmt.Println(err)
	//}

	dt := CSV{name: "csv", docLocal: "df_test.csv", isThirdParty: false}
	df := getData(dt)

	fmt.Println(df)
	fmt.Printf("%T \n", dt)
	fmt.Printf("%T \n", df)
	fmt.Println("-----------------------------------------------------------------------------------------------")

	dt2 := DB{dbType: "mysql", dbLogin: ""}
	df2, _ := dt2.get()
	for df2.Next() {
		var name string
		err := df2.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(name)
		fmt.Printf("%T \n", name)
		fmt.Println(&name)
		fmt.Printf("%T \n", &name)
	}

	fmt.Println(&df2)
	fmt.Println(df2)
	fmt.Printf("%T \n", dt2)
	fmt.Printf("%T \n", df2)
}
