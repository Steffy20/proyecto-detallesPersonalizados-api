package storage

import "proyecto-detallesPersonalizados-api/internal/models"

// commit: implementar almacenamiento en memoria para pedidos
var Pedidos = []models.Pedido{}
var PedidoID = 1

var Personalizaciones = []models.Personalizacion{}
var PersonalizacionID = 1

var ProductosPersonalizados = []models.ProductoPersonalizado{}
var ProductoPersonalizadoID = 1

//modulo de clientes
var Clientes = []models.Cliente{}
var ClienteID = 1

var SeguimientosPedido = []models.SeguimientoPedido{}
var SeguimientoPedidoID = 1

var Reclamos = []models.Reclamo{}
var ReclamoID = 1

<<<<<<< HEAD
//modulo de solicitud urgente
var SolicitudesUrgentes = []models.SolicitudUrgente{}
var SolicitudUrgenteID = 1
var AgendasProduccion = []models.AgendaProduccion{}
var AgendaProduccionID = 1
=======
var SolicitudesUrgentes = []models.SolicitudUrgente{}
var SolicitudUrgenteID = 1

var AgendasProduccion = []models.AgendaProduccion{}
var AgendaProduccionID = 1

>>>>>>> 38f5b01a7af5fa84b7b86cf0cb46324446b95096
var SlotsProduccion = []models.SlotProduccion{}
var SlotProduccionID = 1
