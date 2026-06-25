package storage

import "proyecto-detallesPersonalizados-api/internal/models"

// commit: implementar almacenamiento en memoria para pedidos
var Pedidos = []models.Pedido{}
var PedidoID = 1

var Personalizaciones = []models.Personalizacion{}
var PersonalizacionID = 1

var ProductosPersonalizados = []models.ProductoPersonalizado{}
var ProductoPersonalizadoID = 1


var SolicitudesUrgentes = []models.SolicitudUrgente{}
var SolicitudUrgenteID = 1

var AgendasProduccion = []models.AgendaProduccion{}
var AgendaProduccionID = 1

var SlotsProduccion = []models.SlotProduccion{}
var SlotProduccionID = 1

var Clientes = []models.Cliente{}
var ClienteID = 1

var SeguimientosPedido = []models.SeguimientoPedido{}
var SeguimientoPedidoID = 1

var Reclamos = []models.Reclamo{}
var ReclamoID = 1