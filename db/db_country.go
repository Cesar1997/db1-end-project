package db

import (
	"database/sql"

	"github.com/Cesar1997/db1-end-project/structures"
)

func GetAllCountries() (countries []structures.Country, err error) {
	countries = make([]structures.Country, 0)
	sqlQuery := `SELECT * FROM pais`

	rows, err := db.Query(sqlQuery)

	if err == sql.ErrNoRows {
		return countries, nil
	}

	if err != nil {
		return countries, err
	}

	for rows.Next() {
		var country structures.Country
		err = rows.Scan(
			&country.ID,
			&country.Name,
			&country.Extension,
		)
		if err != nil {
			return countries, err
		}
		countries = append(countries, country)
	}
	return
}
