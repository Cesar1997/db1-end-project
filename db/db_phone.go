package db

import (
	"database/sql"

	"github.com/Cesar1997/db1-end-project/structures"
)

func createPhone(phone structures.Phone) (err error) {
	sqlQuery := `
	INSERT INTO
		telefono
			(
				telefono
				id_pais
			)
	VALUES
		(
			@phoneNumber
            @countryId
		)
	`

	_, err = db.Exec(
		sqlQuery,
		sql.Named("phoneNumber", phone.PhoneNumber),
		sql.Named("countryId", phone.CountryID),
	)
	if err != nil {
		return err
	}

	return nil
}

func createUserPhone(userPhone structures.UserPhone) {

}
