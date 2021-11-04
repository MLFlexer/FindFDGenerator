package main

import (
	"fmt"
	"log"
	"os"
)

//See lecture 04/11 or HomeWork 3

func CreateSQL(table string, attribute1 string, attribute2 string) string {
	sqlScript := "SELECT '" + table + ": " + attribute1 + " --> " + attribute2 + "' AS FD,\n" +
		"CASE WHEN COUNT(*)=0 THEN 'MAY HOLD'\n" +
		"ELSE 'does not hold' END AS VALIDITY\n" +
		"FROM (\n" +
		"	SELECT T." + attribute1 + "\n" +
		"	FROM " + table + " T\n" +
		"	GROUP BY T." + attribute1 + "\n" +
		"	HAVING COUNT(DISTINCT T." + attribute2 + ") > 1\n" +
		") X;\n"
	return sqlScript
}

func main() {
	table := "rentals"
	attributeArr := [...]string{"pid", "hid", "pn", "s", "hs", "hz", "hc"}

	f, err := os.Create("FindFDSScript.sql")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for _, i := range attributeArr {
		for _, j := range attributeArr {
			if i != j {
				fmt.Println(CreateSQL(table, i, j))
				_, err2 := f.WriteString(CreateSQL(table, i, j))
				if err2 != nil {
					log.Fatal(err2)
				}
			}
		}
	}
}
