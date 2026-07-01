-- ===================== PEDIDOS =====================

-- name: ListarPedidos :many
SELECT id, mensaje, estado
FROM pedido;

-- name: BuscarPedidoPorID :one
SELECT id, mensaje, estado
FROM pedido
WHERE id = ?;

-- name: CrearPedido :one
INSERT INTO pedido (mensaje, estado)
VALUES (?, ?)
RETURNING id, mensaje, estado;

-- name: ActualizarPedido :one
UPDATE pedido
SET mensaje = ?, estado = ?
WHERE id = ?
RETURNING id, mensaje, estado;

-- name: BorrarPedido :execrows
DELETE FROM pedido
WHERE id = ?;

-- ===================== PERSONALIZACIONES =====================

-- name: ListarPersonalizaciones :many
SELECT id, pedido_id, mensaje, color
FROM personalizacion;

-- name: BuscarPersonalizacionPorID :one
SELECT id, pedido_id, mensaje, color
FROM personalizacion
WHERE id = ?;

-- name: CrearPersonalizacion :one
INSERT INTO personalizacion (pedido_id, mensaje, color)
VALUES (?, ?, ?)
RETURNING id, pedido_id, mensaje, color;

-- name: ActualizarPersonalizacion :one
UPDATE personalizacion
SET pedido_id = ?, mensaje = ?, color = ?
WHERE id = ?
RETURNING id, pedido_id, mensaje, color;

-- name: BorrarPersonalizacion :execrows
DELETE FROM personalizacion
WHERE id = ?;

-- ===================== PRODUCTOS PERSONALIZADOS =====================

-- name: ListarProductosPersonalizados :many
SELECT id, pedido_id, nombre, cantidad, precio
FROM producto_personalizado;

-- name: BuscarProductoPersonalizadoPorID :one
SELECT id, pedido_id, nombre, cantidad, precio
FROM producto_personalizado     
WHERE id = ?;

-- name: CrearProductoPersonalizado :one
INSERT INTO producto_personalizado
(pedido_id, nombre, cantidad, precio)
VALUES (?, ?, ?, ?)
RETURNING id, pedido_id, nombre, cantidad, precio;

-- name: ActualizarProductoPersonalizado :one
UPDATE producto_personalizado
SET pedido_id = ?, nombre = ?, cantidad = ?, precio = ?
WHERE id = ?
RETURNING id, pedido_id, nombre, cantidad, precio;

-- name: BorrarProductoPersonalizado :execrows
DELETE FROM producto_personalizado
WHERE id = ?;




-- ===================== SOLICITUDES URGENTES =====================

-- name: ListarSolicitudesUrgentes :many
SELECT id, cliente, descripcion, fecha_requerida, estado
FROM SolicitudUrgente;

-- name: BuscarSolicitudUrgentePorID :one
SELECT id, cliente, descripcion, fecha_requerida, estado
FROM SolicitudUrgente
WHERE id = ?;

-- name: CrearSolicitudUrgente :one
INSERT INTO SolicitudUrgente
(cliente, descripcion, fecha_requerida, estado)
VALUES (?, ?, ?, ?)
RETURNING id, cliente, descripcion, fecha_requerida, estado;

-- name: ActualizarSolicitudUrgente :one
UPDATE SolicitudUrgente
SET cliente = ?, descripcion = ?, fecha_requerida = ?, estado = ?
WHERE id = ?
RETURNING id, cliente, descripcion, fecha_requerida, estado;

-- name: BorrarSolicitudUrgente :execrows
DELETE FROM SolicitudUrgente    
WHERE id = ?;

-- ===================== AGENDAS PRODUCCION =====================

-- name: ListarAgendasProduccion :many
SELECT id, fecha, responsable, estado
FROM AgendaProduccion;

-- name: BuscarAgendaProduccionPorID :one
SELECT id, fecha, responsable, estado
FROM AgendaProduccion
WHERE id = ?;

-- name: CrearAgendaProduccion :one
INSERT INTO AgendaProduccion
(fecha, responsable, estado)
VALUES (?, ?, ?)
RETURNING id, fecha, responsable, estado;

-- name: ActualizarAgendaProduccion :one
UPDATE AgendaProduccion
SET fecha = ?, responsable = ?, estado = ?
WHERE id = ?
RETURNING id, fecha, responsable, estado;

-- name: BorrarAgendaProduccion :execrows
DELETE FROM AgendaProduccion
WHERE id = ?;

-- ===================== SLOTS PRODUCCION =====================

-- name: ListarSlotsProduccion :many
SELECT id, agenda_id, capacidad_maxima, pedidos_asignados
FROM slotProduccion;

-- name: BuscarSlotProduccionPorID :one
SELECT id, agenda_id, capacidad_maxima, pedidos_asignados
FROM slotProduccion
WHERE id = ?;

-- name: CrearSlotProduccion :one
INSERT INTO slotProduccion
(agenda_id, capacidad_maxima, pedidos_asignados)
VALUES (?, ?, ?)
RETURNING id, agenda_id, capacidad_maxima, pedidos_asignados;

-- name: ActualizarSlotProduccion :one
UPDATE slotProduccion
SET agenda_id = ?, capacidad_maxima = ?, pedidos_asignados = ?
WHERE id = ?
RETURNING id, agenda_id, capacidad_maxima, pedidos_asignados;

-- name: BorrarSlotProduccion :execrows
DELETE FROM slotProduccion
WHERE id = ?;



-- ===================== CLIENTES =====================

-- name: ListarClientes :many
SELECT id, nombre, telefono
FROM cliente;

-- name: BuscarClientePorID :one
SELECT id, nombre, telefono
FROM cliente
WHERE id = ?;

-- name: CrearCliente :one
INSERT INTO cliente
(nombre, telefono)
VALUES (?, ?)
RETURNING id, nombre, telefono;

-- name: ActualizarCliente :one
UPDATE cliente
SET nombre = ?, telefono = ?
WHERE id = ?
RETURNING id, nombre, telefono;

-- name: BorrarCliente :execrows
DELETE FROM cliente
WHERE id = ?;

-- ===================== RECLAMOS =====================

-- name: ListarReclamos :many
SELECT id, cliente_id, pedido_id, descripcion, estado
FROM Reclamo;

-- name: BuscarReclamoPorID :one
SELECT id, cliente_id, pedido_id, descripcion, estado
FROM Reclamo
WHERE id = ?;

-- name: CrearReclamo :one
INSERT INTO Reclamo
(cliente_id, pedido_id, descripcion, estado)
VALUES (?, ?, ?, ?)
RETURNING id, cliente_id, pedido_id, descripcion, estado;

-- name: ActualizarReclamo :one
UPDATE Reclamo
SET cliente_id = ?, pedido_id = ?, descripcion = ?, estado = ?
WHERE id = ?
RETURNING id, cliente_id, pedido_id, descripcion, estado;

-- name: BorrarReclamo :execrows
DELETE FROM Reclamo
WHERE id = ?;

-- ===================== SEGUIMIENTO PEDIDOS =====================

-- name: ListarSeguimientosPedido :many
SELECT id, pedido_id, estado, fecha_estado
FROM SeguimientoPedido;

-- name: BuscarSeguimientoPedidoPorID :one
SELECT id, pedido_id, estado, fecha_estado
FROM SeguimientoPedido
WHERE id = ?;

-- name: CrearSeguimientoPedido :one
INSERT INTO SeguimientoPedido
(pedido_id, estado, fecha_estado)
VALUES (?, ?, ?)
RETURNING id, pedido_id, estado, fecha_estado;

-- name: ActualizarSeguimientoPedido :one
UPDATE SeguimientoPedido
SET pedido_id = ?, estado = ?, fecha_estado = ?
WHERE id = ?
RETURNING id, pedido_id, estado, fecha_estado;

-- name: BorrarSeguimientoPedido :execrows
DELETE FROM SeguimientoPedido
WHERE id = ?;

