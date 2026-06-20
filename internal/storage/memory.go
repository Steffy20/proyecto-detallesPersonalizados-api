package storage

import "proyecto-detallesPersonalizados-api/internal/models"


//commit: implementar almacenamiento en memoria para pedidos
var Pedidos = []models.Pedido{}
var PedidoID = 1

var Personalizaciones = []models.Personalizacion{}
var PersonalizacionID = 1

var ProductosPersonalizados = []models.ProductoPersonalizado{}
var ProductoPersonalizadoID = 1