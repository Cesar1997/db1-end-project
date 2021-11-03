package db

import (
	"database/sql"

	"github.com/Cesar1997/db1-end-project/structures"
)

func GetAllProductsFilteredByUser(profileID int) (listProducts []structures.ProductType, err error) {
	listProducts = make([]structures.ProductType, 0)
	sqlQuery := `
		SELECT
			prod.id_producto as id_producto,
			ar.nombre  as producto,
			ar.descripcion  as descripcion_prod,
			ar.precio_por_hora  as precio,
			col.nombre  as color,
			col.color_hexa  as color,
			tam.tamanio  as tamanio
		FROM
			producto prod
		INNER JOIN articulo ar ON ar.id_articulo  = prod.id_articulo
		INNER JOIN color col ON col.id_color  = prod.id_color
		INNER JOIN tamanio tam ON tam.id_tamanio  = prod.id_tamanio
		WHERE
			prod.id_perfil_arrendador  = @profileID

		`
	rows, err := db.Query(sqlQuery, sql.Named("profileID", profileID))
	if err == sql.ErrNoRows {
		return listProducts, nil
	}
	if err != nil {
		return listProducts, err
	}
	for rows.Next() {
		var product structures.ProductType
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.PriceXHour,
			&product.ColorInfo.Name,
			&product.ColorInfo.ColorHex,
			&product.SizeInfo.Size,
		)
		if err != nil {
			return listProducts, err
		}

		product.Photos, _ = GetPhotosByProductID(product.ID)

		listProducts = append(listProducts, product)
	}
	return
}

func GetProduct(productID int) (structures.ProductType, error) {
	sqlQuery := `
		SELECT
			prod.id_producto as id_producto,
			ar.nombre  as producto,
			ar.descripcion  as descripcion_prod,
			ar.precio_por_hora  as precio
		FROM
			producto prod
		INNER JOIN articulo ar ON ar.id_articulo  = prod.id_articulo
		WHERE
			prod.id_producto  = @productID

		`
	rows, err := db.Query(sqlQuery, sql.Named("productID", productID))
	if err == sql.ErrNoRows {
		return structures.ProductType{}, nil
	}
	if err != nil {
		return structures.ProductType{}, err
	}
	var product structures.ProductType
	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.PriceXHour,
		)
		if err != nil {
			return product, err
		}
		product.Photos, _ = GetPhotosByProductID(product.ID)
	}
	return product, nil
}

func GetPhotosByProductID(productID int) ([]structures.Photo, error) {
	listPhotos := make([]structures.Photo, 0)
	sqlQuery := `
		SELECT
			photo.foto ,
			photo.formato
		FROM
			producto_foto prodPhoto
		INNER JOIN foto photo ON photo.id_foto  = prodPhoto.id_foto 
		WHERE (photo.foto IS NOT NULL AND photo.foto  <> '') AND 
			 prodPhoto.id_producto = @IdProducto
	`
	rows, err := db.Query(sqlQuery, sql.Named("IdProducto", productID))
	if err == sql.ErrNoRows {
		return listPhotos, nil
	}
	if err != nil {
		return listPhotos, err
	}
	for rows.Next() {
		var photo structures.Photo
		err = rows.Scan(
			&photo.Photo,
			&photo.Extension,
		)
		if err != nil {
			return listPhotos, err
		}

		listPhotos = append(listPhotos, photo)
	}
	return listPhotos, nil
}

func GetAllSizes() (sizes []structures.Size, err error) {
	sizes = make([]structures.Size, 0)
	sqlQuery := `SELECT * FROM tamanio`

	rows, err := db.Query(sqlQuery)

	if err == sql.ErrNoRows {
		return sizes, nil
	}

	if err != nil {
		return sizes, err
	}

	for rows.Next() {
		var size structures.Size
		err = rows.Scan(
			&size.ID,
			&size.Size,
		)
		if err != nil {
			return sizes, err
		}
		sizes = append(sizes, size)
	}
	return
}

func GetAllColors() (colors []structures.Color, err error) {
	colors = make([]structures.Color, 0)
	sqlQuery := `SELECT * FROM color`

	rows, err := db.Query(sqlQuery)

	if err == sql.ErrNoRows {
		return colors, nil
	}

	if err != nil {
		return colors, err
	}

	for rows.Next() {
		var color structures.Color
		err = rows.Scan(
			&color.ID,
			&color.Name,
			&color.ColorHex,
		)
		if err != nil {
			return colors, err
		}
		colors = append(colors, color)
	}
	return
}

func GetAllTypeArticle() (typeArticles []structures.TypeArticle, err error) {
	typeArticles = make([]structures.TypeArticle, 0)
	sqlQuery := `SELECT * FROM tipo_articulo`

	rows, err := db.Query(sqlQuery)

	if err == sql.ErrNoRows {
		return typeArticles, nil
	}

	if err != nil {
		return typeArticles, err
	}

	for rows.Next() {
		var typeArticle structures.TypeArticle
		err = rows.Scan(
			&typeArticle.ID,
			&typeArticle.Description,
		)
		if err != nil {
			return typeArticles, err
		}
		typeArticles = append(typeArticles, typeArticle)
	}
	return
}

func CreateProduct(product structures.ProductType) (err error) {

	sqlQuery := `
		EXEC sp_insert_producto
			@IdPerfilArrendador,
			@IdColor,
			@IdTamanio,
			@NombreArticulo,
			@Descripcion,
			@PrecioPorHora,
			@IdTipoArticulo,
			@Foto
	`
	_, err = db.Exec(
		sqlQuery,
		sql.Named("IdPerfilArrendador", product.SellerProfileID),
		sql.Named("IdColor", product.ColorID),
		sql.Named("IdTamanio", product.SizeID),
		sql.Named("NombreArticulo", product.Name),
		sql.Named("Descripcion", product.Description),
		sql.Named("PrecioPorHora", product.PriceXHour),
		sql.Named("IdTipoArticulo", product.TypeArticleID),
		sql.Named("Foto", product.ImageURL),
	)
	if err != nil {
		return err
	}

	return nil
}
