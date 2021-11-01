package db

import (
	"database/sql"

	"github.com/Cesar1997/db1-end-project/structures"
)

func GetAllAdds() (listAdds []structures.Adds, err error) {

	sqlQuery := `
		SELECT
			anun.id_anuncio,
			anun.id_perfil_arrendador,
			anun.id_producto,
			anun.titulo,
			anun.descripcion,
			ar.nombre  as producto,
			ar.descripcion  as descripcion_prod,
			ar.precio_por_hora  as precio,
			col.nombre  as color,
			col.color_hexa  as color,
			tam.tamanio  as tamanio
		FROM  anuncio anun
		INNER JOIN producto prod ON prod.id_producto  = anun.id_producto
		INNER JOIN articulo ar ON ar.id_articulo  = prod.id_articulo
		INNER JOIN color col ON col.id_color  = prod.id_color
		INNER JOIN tamanio tam ON tam.id_tamanio  = prod.id_tamanio
	`

	rows, err := db.Query(sqlQuery)

	if err == sql.ErrNoRows {
		return listAdds, nil
	}
	if err != nil {
		return listAdds, err
	}
	for rows.Next() {
		var add structures.Adds
		err = rows.Scan(
			&add.ID,
			&add.IDProfile,
			&add.Product.ID,
			&add.Title,
			&add.Description,
			&add.Product.Name,
			&add.Product.Description,
			&add.Product.PriceXHour,
			&add.Product.ColorInfo.Name,
			&add.Product.ColorInfo.ColorHex,
			&add.Product.SizeInfo.Size,
		)
		if err != nil {
			return listAdds, err
		}
		listAdds = append(listAdds, add)
	}
	return
}

func GetAdd(addID int) (structures.Adds, error) {

	sqlQuery := `
		SELECT
			anun.id_anuncio,
			anun.id_perfil_arrendador,
			anun.id_producto,
			anun.titulo,
			anun.descripcion,
			ar.nombre  as producto,
			ar.descripcion  as descripcion_prod,
			ar.precio_por_hora  as precio,
			col.nombre  as color,
			col.color_hexa  as color,
			tam.tamanio  as tamanio
		FROM  anuncio anun
		INNER JOIN producto prod ON prod.id_producto  = anun.id_producto
		INNER JOIN articulo ar ON ar.id_articulo  = prod.id_articulo
		INNER JOIN color col ON col.id_color  = prod.id_color
		INNER JOIN tamanio tam ON tam.id_tamanio  = prod.id_tamanio
		WHERE anun.id_anuncio = @addID
	`

	rows, err := db.Query(sqlQuery, sql.Named("addID", addID))

	if err == sql.ErrNoRows {
		return structures.Adds{}, nil
	}
	if err != nil {
		return structures.Adds{}, err
	}

	var add structures.Adds
	for rows.Next() {
		err = rows.Scan(
			&add.ID,
			&add.IDProfile,
			&add.Product.ID,
			&add.Title,
			&add.Description,
			&add.Product.Name,
			&add.Product.Description,
			&add.Product.PriceXHour,
			&add.Product.ColorInfo.Name,
			&add.Product.ColorInfo.ColorHex,
			&add.Product.SizeInfo.Size,
		)
		if err != nil {
			return structures.Adds{}, err
		}
	}
	return add, nil
}

func GetAllCategories() (categories []structures.Category, err error) {
	categories = make([]structures.Category, 0)
	sqlQuery := `SELECT  * FROM  categoria`

	rows, err := db.Query(sqlQuery)

	if err == sql.ErrNoRows {
		return categories, nil
	}

	if err != nil {
		return categories, err
	}

	for rows.Next() {
		var category structures.Category
		err = rows.Scan(
			&category.ID,
			&category.Description,
		)
		if err != nil {
			return categories, err
		}
		categories = append(categories, category)
	}
	return
}

func CreateAdd(add structures.Adds) error {
	sqlQuery := `
		INSERT INTO
			anuncio
				(
					id_perfil_arrendador,
					id_producto,
					titulo,
					descripcion
				)
		OUTPUT
			INSERTED.id_anuncio
		VALUES
			(
				@IdPerfilArrendador,
				@IdProducto,
				@Titulo,
				@Descripcion
			)
	`

	row := db.QueryRow(
		sqlQuery,
		sql.Named("IdPerfilArrendador", add.IDProfile),
		sql.Named("IdProducto", add.IDProduct),
		sql.Named("Titulo", add.Title),
		sql.Named("Descripcion", add.Description),
	)

	var insertedID sql.NullInt32

	err := row.Scan(
		&insertedID,
	)
	if err != nil {
		return err
	}

	if insertedID.Int32 != 0 {
		createCategoriesAdds(int(insertedID.Int32), add.Categories)
	}

	return nil
}

func createCategoriesAdds(addsID int, categoriesList []structures.Category) error {

	for _, value := range categoriesList {
		sqlQuery := `
			INSERT INTO
				anuncio_categoria(
					id_anuncio,
					id_categoria
				)
				VALUES(
					@IdAnuncio,
					@IdCategoria
				)
		`
		_, err := db.Exec(
			sqlQuery,
			sql.Named("IdAnuncio", addsID),
			sql.Named("IdCategoria", value.ID),
		)

		if err != nil {
			return err
		}

	}

	return nil
}
