package db

import (
	"database/sql"

	"github.com/Cesar1997/db1-end-project/structures"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user structures.User) (err error) {

	password, err := HashPassword(user.Password)

	if err != nil {
		return err
	}

	sqlQuery := `
		EXEC sp_insert_usuario
			@IdPais,
			@Email,
			@Password,
			@Nombres,
			@Apellidos,
			@Direccion,
			@Telefono,
			@Foto,
			@Formato,
			@EsArrendador,
			@NombreNegocio,
			@Nit
	`
	_, err = db.Exec(
		sqlQuery,
		sql.Named("IdPais", user.CountryID),
		sql.Named("Email", user.Email),
		sql.Named("Password", password),
		sql.Named("Nombres", user.Name),
		sql.Named("Apellidos", user.LastName),
		sql.Named("Direccion", user.Address),
		sql.Named("Telefono", user.InputNumberPhone),
		sql.Named("Foto", user.InputPhoto),
		sql.Named("Formato", user.InputExtension),
		sql.Named("EsArrendador", user.InputIsProfileLessor),
		sql.Named("NombreNegocio", user.InputNameCompany),
		sql.Named("Nit", user.InputNumberNit),
	)
	if err != nil {
		return err
	}

	return nil
}

func Login(user structures.User) (structures.User, error) {

	sqlQuery := `
	    SELECT
			id_usuario,
			email,
			password,
			nombres,
			apellidos,
			direccion
		FROM usuario
		WHERE email = @Email
	`
	row := db.QueryRow(
		sqlQuery,
		sql.Named("Email", user.Email),
	)

	var findedUser structures.User

	err := row.Scan(
		&findedUser.UserID,
		&findedUser.Email,
		&findedUser.Password,
		&findedUser.Name,
		&findedUser.LastName,
		&findedUser.Address,
	)

	if err == sql.ErrNoRows {
		return structures.User{}, nil
	}

	if err != nil {
		return structures.User{}, err // proper error handling instead of panic in your app
	}

	if !CheckPasswordHash(user.Password, findedUser.Password) {
		return structures.User{}, err
	}
	findedUser.Password = ""

	findedUser.SellerProfile, _ = getSellerProfile(findedUser.UserID)
	findedUser.ConsumerProfile, _ = getConsumerProfile(findedUser.UserID)
	return findedUser, nil
}

func getSellerProfile(userID int) (structures.SellerProfileType, error) {
	sqlQuery := `
	    SELECT
			id_perfil_arrendador,
			nombre_negocio,
			nit
		FROM perfil_arrendador
		WHERE id_usuario = @userID
	`
	row := db.QueryRow(
		sqlQuery,
		sql.Named("userID", userID),
	)

	var profile structures.SellerProfileType

	err := row.Scan(
		&profile.ID,
		&profile.BussinessName,
		&profile.NitNumber,
	)

	if err == sql.ErrNoRows {
		return profile, nil
	}

	if err != nil {
		return profile, err // proper error handling instead of panic in your app
	}

	return profile, nil
}

func getConsumerProfile(userID int) (structures.ConsumerProfileType, error) {
	sqlQuery := `
	    SELECT
			id_perfil_arrendatario
		FROM perfil_arrendatario
		WHERE id_usuario = @userID
	`
	row := db.QueryRow(
		sqlQuery,
		sql.Named("userID", userID),
	)

	var profile structures.ConsumerProfileType

	err := row.Scan(
		&profile.ID,
	)

	if err == sql.ErrNoRows {
		return profile, nil
	}

	if err != nil {
		return profile, err // proper error handling instead of panic in your app
	}

	return profile, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
