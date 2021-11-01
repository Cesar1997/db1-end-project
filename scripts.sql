-- ABOUT THIS PROJECT
-- using sqlServer and go for create web services
ALTER TABLE dressMyTrip.dbo.usuario 
ALTER COLUMN password 
varchar(100) COLLATE 
SQL_Latin1_General_CP1_CI_AS 
NOT NULL;

ALTER TABLE dressMyTrip.dbo.usuario ALTER COLUMN password varchar(60) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL;



INSERT INTO pais (nombre, extension)
VALUES( N'Guatemala', N'502'),
	   ( 'El salvador',503),
	   ( 'Honduras',504),
       ( 'Nicaragua',505);


CREATE PROCEDURE sp_insert_foto(
	@Foto AS VARCHAR(100),
	@Formato AS VARCHAR(4),
	@IdUsuario AS INT
) AS
BEGIN
	DECLARE @IdFoto as INT
	SET NOCOUNT ON;
	INSERT INTO foto(
		foto,
		formato
	)
	VALUES(
		@Foto,
		@Formato
	)
	SET @IdFoto = SCOPE_IDENTITY();

	IF @IdFoto IS NOT NULL 
	BEGIN
		INSERT INTO usuario_foto(
			id_usuario,
			id_foto
		)
		VALUES(
			@IdUsuario,
			@IdFoto
		)
	END
END;

CREATE PROCEDURE sp_insert_telefono(
	@IdPais AS INT,
	@IdUsuario AS INT,
	@Telefono AS INT
) AS
BEGIN
	DECLARE @IdTelefono as INT
	SET NOCOUNT ON;
	INSERT INTO telefono(
		id_pais,
		telefono
	)
	VALUES(
		@IdPais,
		@Telefono
	)
	SET @IdTelefono = SCOPE_IDENTITY();

	IF @IdTelefono IS NOT NULL
	BEGIN
		INSERT INTO usuario_telefono(
			id_usuario,
			id_telefono,
			es_principal
		)
		VALUES(
			@IdUsuario,
			@IdTelefono,
			1
		)
	END
END;


CREATE PROCEDURE sp_insert_perfil_arrendador(
	@IdUsuario AS INT,
	@NombreNegocio AS VARCHAR(100),
	@Nit AS VARCHAR(12)
) AS
BEGIN
	SET NOCOUNT ON;
	INSERT INTO perfil_arrendador(
		id_usuario,
		nombre_negocio,
		nit
	)
	VALUES(
		@IdUsuario,
		@NombreNegocio,
		@Nit
	)
END;

USE dressMyTrip;
ALTER PROCEDURE sp_insert_usuario(
	@IdPais AS INT,
	@Email AS VARCHAR(50),
	@Password AS VARCHAR(60),
	@Nombres AS VARCHAR(60),
	@Apellidos AS VARCHAR(60),
	@Direccion AS VARCHAR(100),
	@Telefono AS INT,
	@Foto AS VARCHAR(100),
	@Formato AS VARCHAR(4),
	@EsArrendador AS TINYINT,
	@NombreNegocio AS VARCHAR(100),
	@Nit 		   AS VARCHAR(12)
) AS
BEGIN
	DECLARE @IdUsuario as INT
	SET NOCOUNT ON;
	INSERT INTO usuario(
		id_pais,
		email,
		password,
		nombres,
		apellidos,
		direccion
	)
	VALUES (
		@IdPais,
		@Email,
		@Password,
		@Nombres,
		@Apellidos,
		@Direccion
	)

	SET @IdUsuario = SCOPE_IDENTITY();
    -- crear perfil del arrendatario
	INSERT INTO perfil_arrendatario(id_usuario) VALUES (@IdUsuario)	

	IF @Foto IS NOT NULL
	BEGIN
		EXEC sp_insert_foto @Foto,@Formato,@IdUsuario
	END

	IF @Telefono IS NOT NULL
	BEGIN
		EXEC sp_insert_telefono @IdPais, @IdUsuario, @Telefono
	END

	IF @EsArrendador = 1
	BEGIN
		EXEC sp_insert_perfil_arrendador @IdUsuario, @NombreNegocio, @Nit
	END
END;






--test without email
EXEC sp_insert_usuario 2, 'andres.2@gmail.com', '123456', 'cesar', 'apelidos', '4ta calle 1-12', NULL, NULL,'.jpeg';
EXEC sp_insert_usuario 2, 'user.4@gmail.com', '123456', 'user4', 'app4', '4ta calle 1-12', NULL, NULL,'.jpeg';


/* perfil arrendador */


/* Perfil arrendador */

/*  */
/* SCRIPTS PARA PRODUCTOS*/

-- articulos 

INSERT INTO tipo_articulo (descripcion) 
VALUES('Tienda de acampar'),
	  ('Zapatos'),
	  ('Linterna'),
	  ('Bloqueador solar'),
	  ('Tienda de acampar'),
	  ('Sleeping'),
	  ('Vaso'),
	  ('Automoviles'),
	  ('Camas');

INSERT INTO tamanio (tamanio)
VALUES('Small'),
 	  ('Medium'),
 	  ('Large');

INSERT INTO color (nombre, color_hexa)
VALUES('Rojo','#FF5733'),
 	  ('Naranja','#FFA533'),
 	  ('Amarillo','#FFF933'),
	  ('Verde','#A5FF33');


CREATE PROCEDURE sp_insert_articulo(
	@Nombre AS VARCHAR(100),
	@Descripcion AS VARCHAR(100),
	@PrecioPorHora AS DECIMAL(8,2),
	@IdTipoArticulo AS INT,
	@IdArticulo INT = NULL OUTPUT
) AS
BEGIN
	SET NOCOUNT ON;

	SET @IdArticulo  = (
		SELECT
			id_articulo
		FROM
			articulo
		WHERE
			nombre = @Nombre
	)

	IF @IdArticulo IS NULL
	BEGIN
		INSERT INTO articulo(
			nombre,
			descripcion,
			precio_por_hora,
			id_tipo_articulo
		)
		VALUES(
			@Nombre,
			@Descripcion,
			@PrecioPorHora,
			@IdTipoArticulo
		)
		SET @IdArticulo = SCOPE_IDENTITY();
	END
END;



CREATE PROCEDURE sp_insert_producto(
	@IdPerfilArrendador AS INT,
	@IdColor AS INT,
	@IdTamanio AS INT,
	@NombreArticulo AS VARCHAR(100),
	@Descripcion AS VARCHAR(100),
	@PrecioPorHora AS DECIMAL(8,2),
	@IdTipoArticulo AS INT
) AS
BEGIN
	DECLARE @IdArticulo as INT
	SET NOCOUNT ON;

	EXEC sp_insert_articulo @NombreArticulo, @Descripcion, @PrecioPorHora, @IdTipoArticulo, @IdArticulo = @IdArticulo OUTPUT;


	INSERT INTO producto(
		id_perfil_arrendador,
		id_color,
		id_tamanio,
		id_articulo,
		estado
	)
	VALUES(
		@IdPerfilArrendador,
		@IdColor,
		@IdTamanio,
		@IdArticulo,
		1
	)
END;

EXEC  sp_insert_producto 4,1,3,'Zapatos','Zapatos marca timberland',10,2
EXEC  sp_insert_producto 4,1,3,'Sweter','Zapatos marca timberland',10,2
/* END SCRIPTS PARA PRODUCTOS */

/* SCRIPTS PARA Anuncio */


INSERT INTO 
categoria (descripcion) 
VALUES('Deportes'),
	  ('Ropa y hogar'),
	  ('Verano'),
	  ('Invierno'),
	  ('Herramientas de caza');


/* 
CREATE PROCEDURE sp_insert_anuncio(
	@IdPerfilArrendador AS INT,
	@IdProducto AS  INT,
	@Titulo AS VARCHAR(50),
	@Descripcion AS VARCHAR(100) 
) AS
BEGIN
SET NOCOUNT ON;
	INSERT INTO anuncio(
		id_perfil_arrendador,
		id_producto,
		titulo,
		descripcion
	)
	VALUES(
		@IdPerfilArrendador,
		@IdProducto,
		@Titulo,
		@Descripcion
	)
END */





/* END SCRIPTS PARA ANUNCIO */


