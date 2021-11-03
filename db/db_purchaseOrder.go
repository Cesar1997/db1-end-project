package db

import (
	"database/sql"

	"github.com/Cesar1997/db1-end-project/structures"
)

func CreatePurchaseOrder(purchaseOrder structures.PurchaseOrderType) error {
	sqlQuery := `
		INSERT INTO orden_alquiler(
			id_perfil_arrendatario,
			fecha_creacion,
			fecha_pago,
			total_orden,
			estado
		)
		OUTPUT
			INSERTED.id_orden_alquiler
		VALUES (
			@IdPerfilArrendatario,
			GETDATE(),
			GETDATE(),
			@TotalOrden,
			1
		)`
	row := db.QueryRow(
		sqlQuery,
		sql.Named("IdPerfilArrendatario", purchaseOrder.ConsumerProfileID),
		sql.Named("TotalOrden", purchaseOrder.TotalOrder),
	)

	var insertedID sql.NullInt32

	err := row.Scan(
		&insertedID,
	)
	if err != nil {
		return err
	}

	if insertedID.Int32 != 0 {
		err := createDetailOrder(int(insertedID.Int32), purchaseOrder.Detail)
		if err != nil {
			return err
		}
	}
	return nil
}

func createDetailOrder(purchaseOrderID int, details []structures.DetailPurchaseOrderType) error {

	for _, value := range details {
		sqlQuery := `
			INSERT INTO
				detalle_orden(
					id_orden_alquiler,
					id_producto,
					fecha_entrega,
					cant_horas_arrendado,
					subtotal
				)
				VALUES(
					@IdOrdenAlquiler,
					@IdProducto,
					@FechaEntrega,
					@CantHorasArrendado,
					@Subtotal
				)
		`
		_, err := db.Exec(
			sqlQuery,
			sql.Named("IdOrdenAlquiler", purchaseOrderID),
			sql.Named("IdProducto", value.ProductID),
			sql.Named("FechaEntrega", value.DateTimeToDevolution),
			sql.Named("CantHorasArrendado", value.CantHours),
			sql.Named("Subtotal", value.SubTotal),
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func GetAllPurchaseOrder(consumerProfileID int) (reporteVentas []structures.ReporteVentas, err error) {
	reporteVentas = make([]structures.ReporteVentas, 0)
	sqlQuery := `
			SELECT
				orden.id_orden_alquiler,
			 	u.nombres,
				art.nombre,
			 	tipo.descripcion,
				col.nombre,
				tam.tamanio ,
				art.precio_por_hora,
				det.cant_horas_arrendado,
				det.subtotal
			FROM detalle_orden det
			INNER JOIN orden_alquiler orden ON orden.id_orden_alquiler  = det.id_orden_alquiler
			INNER JOIN producto prod ON prod.id_producto  = det.id_producto 
			INNER JOIN articulo art ON art.id_articulo  = prod.id_articulo 
			INNER JOIN tipo_articulo tipo ON tipo.id_tipo_articulo = art.id_tipo_articulo 
			INNER JOIN color col ON col.id_color  = prod.id_color 
			INNER JOIN tamanio tam ON tam.id_tamanio  = prod.id_tamanio 
			INNER JOIN perfil_arrendador pa ON pa.id_perfil_arrendador  = prod.id_perfil_arrendador 
			INNER JOIN usuario u ON u.id_usuario  = pa.id_usuario 
			WHERE orden.id_perfil_arrendatario  =  @profileID
		`
	rows, err := db.Query(sqlQuery, sql.Named("profileID", consumerProfileID))
	if err == sql.ErrNoRows {
		return reporteVentas, nil
	}
	if err != nil {
		return reporteVentas, err
	}
	for rows.Next() {
		var detail structures.ReporteVentas
		err = rows.Scan(
			&detail.IdOrden,
			&detail.Arrendatario,
			&detail.Articulo,
			&detail.Tipo,
			&detail.Color,
			&detail.Tamanio,
			&detail.Precio,
			&detail.CantHoras,
			&detail.Subtotal,
		)
		if err != nil {
			return reporteVentas, err
		}
		reporteVentas = append(reporteVentas, detail)
	}
	return
}
