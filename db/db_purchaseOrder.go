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
