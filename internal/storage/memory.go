package storage

import "proyecto-detallesPersonalizados-api/internal/models"

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

type AlmacenMemoria struct{}

func NuevoAlmacenMemoria() *AlmacenMemoria {
	return &AlmacenMemoria{}
}

// =========================================================
// PEDIDOS
// =========================================================

func (a *AlmacenMemoria) ListarPedidos() []models.Pedido {
	return Pedidos
}

func (a *AlmacenMemoria) BuscarPedidoPorID(id int) (models.Pedido, bool) {
	for _, pedido := range Pedidos {
		if pedido.ID == id {
			return pedido, true
		}
	}
	return models.Pedido{}, false
}

func (a *AlmacenMemoria) CrearPedido(p models.Pedido) models.Pedido {
	p.ID = PedidoID
	PedidoID++
	Pedidos = append(Pedidos, p)
	return p
}

func (a *AlmacenMemoria) ActualizarPedido(id int, datos models.Pedido) (models.Pedido, bool) {
	for i, pedido := range Pedidos {
		if pedido.ID == id {
			datos.ID = id
			Pedidos[i] = datos
			return datos, true
		}
	}
	return models.Pedido{}, false
}

func (a *AlmacenMemoria) BorrarPedido(id int) bool {
	for i, pedido := range Pedidos {
		if pedido.ID == id {
			Pedidos = append(Pedidos[:i], Pedidos[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// PERSONALIZACIONES
// =========================================================

func (a *AlmacenMemoria) ListarPersonalizaciones() []models.Personalizacion {
	return Personalizaciones
}

func (a *AlmacenMemoria) BuscarPersonalizacionPorID(id int) (models.Personalizacion, bool) {
	for _, p := range Personalizaciones {
		if p.ID == id {
			return p, true
		}
	}
	return models.Personalizacion{}, false
}

func (a *AlmacenMemoria) CrearPersonalizacion(p models.Personalizacion) models.Personalizacion {
	p.ID = PersonalizacionID
	PersonalizacionID++
	Personalizaciones = append(Personalizaciones, p)
	return p
}

func (a *AlmacenMemoria) ActualizarPersonalizacion(id int, datos models.Personalizacion) (models.Personalizacion, bool) {
	for i, p := range Personalizaciones {
		if p.ID == id {
			datos.ID = id
			Personalizaciones[i] = datos
			return datos, true
		}
	}
	return models.Personalizacion{}, false
}

func (a *AlmacenMemoria) BorrarPersonalizacion(id int) bool {
	for i, p := range Personalizaciones {
		if p.ID == id {
			Personalizaciones = append(Personalizaciones[:i], Personalizaciones[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// PRODUCTOS PERSONALIZADOS
// =========================================================

func (a *AlmacenMemoria) ListarProductosPersonalizados() []models.ProductoPersonalizado {
	return ProductosPersonalizados
}

func (a *AlmacenMemoria) BuscarProductoPersonalizadoPorID(id int) (models.ProductoPersonalizado, bool) {
	for _, p := range ProductosPersonalizados {
		if p.ID == id {
			return p, true
		}
	}
	return models.ProductoPersonalizado{}, false
}

func (a *AlmacenMemoria) CrearProductoPersonalizado(p models.ProductoPersonalizado) models.ProductoPersonalizado {
	p.ID = ProductoPersonalizadoID
	ProductoPersonalizadoID++
	ProductosPersonalizados = append(ProductosPersonalizados, p)
	return p
}

func (a *AlmacenMemoria) ActualizarProductoPersonalizado(id int, datos models.ProductoPersonalizado) (models.ProductoPersonalizado, bool) {
	for i, p := range ProductosPersonalizados {
		if p.ID == id {
			datos.ID = id
			ProductosPersonalizados[i] = datos
			return datos, true
		}
	}
	return models.ProductoPersonalizado{}, false
}

func (a *AlmacenMemoria) BorrarProductoPersonalizado(id int) bool {
	for i, p := range ProductosPersonalizados {
		if p.ID == id {
			ProductosPersonalizados = append(ProductosPersonalizados[:i], ProductosPersonalizados[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// SOLICITUDES URGENTES
// =========================================================

func (a *AlmacenMemoria) ListarSolicitudesUrgentes() []models.SolicitudUrgente {
	return SolicitudesUrgentes
}

func (a *AlmacenMemoria) BuscarSolicitudUrgentePorID(id int) (models.SolicitudUrgente, bool) {
	for _, s := range SolicitudesUrgentes {
		if s.ID == id {
			return s, true
		}
	}
	return models.SolicitudUrgente{}, false
}

func (a *AlmacenMemoria) CrearSolicitudUrgente(s models.SolicitudUrgente) models.SolicitudUrgente {
	s.ID = SolicitudUrgenteID
	SolicitudUrgenteID++
	SolicitudesUrgentes = append(SolicitudesUrgentes, s)
	return s
}

func (a *AlmacenMemoria) ActualizarSolicitudUrgente(id int, datos models.SolicitudUrgente) (models.SolicitudUrgente, bool) {
	for i, s := range SolicitudesUrgentes {
		if s.ID == id {
			datos.ID = id
			SolicitudesUrgentes[i] = datos
			return datos, true
		}
	}
	return models.SolicitudUrgente{}, false
}

func (a *AlmacenMemoria) BorrarSolicitudUrgente(id int) bool {
	for i, s := range SolicitudesUrgentes {
		if s.ID == id {
			SolicitudesUrgentes = append(SolicitudesUrgentes[:i], SolicitudesUrgentes[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// AGENDAS DE PRODUCCIÓN
// =========================================================

func (a *AlmacenMemoria) ListarAgendasProduccion() []models.AgendaProduccion {
	return AgendasProduccion
}

func (a *AlmacenMemoria) BuscarAgendaProduccionPorID(id int) (models.AgendaProduccion, bool) {
	for _, ag := range AgendasProduccion {
		if ag.ID == id {
			return ag, true
		}
	}
	return models.AgendaProduccion{}, false
}

func (a *AlmacenMemoria) CrearAgendaProduccion(ag models.AgendaProduccion) models.AgendaProduccion {
	ag.ID = AgendaProduccionID
	AgendaProduccionID++
	AgendasProduccion = append(AgendasProduccion, ag)
	return ag
}

func (a *AlmacenMemoria) ActualizarAgendaProduccion(id int, datos models.AgendaProduccion) (models.AgendaProduccion, bool) {
	for i, ag := range AgendasProduccion {
		if ag.ID == id {
			datos.ID = id
			AgendasProduccion[i] = datos
			return datos, true
		}
	}
	return models.AgendaProduccion{}, false
}

func (a *AlmacenMemoria) BorrarAgendaProduccion(id int) bool {
	for i, ag := range AgendasProduccion {
		if ag.ID == id {
			AgendasProduccion = append(AgendasProduccion[:i], AgendasProduccion[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// SLOTS DE PRODUCCIÓN
// =========================================================

func (a *AlmacenMemoria) ListarSlotsProduccion() []models.SlotProduccion {
	return SlotsProduccion
}

func (a *AlmacenMemoria) BuscarSlotProduccionPorID(id int) (models.SlotProduccion, bool) {
	for _, s := range SlotsProduccion {
		if s.ID == id {
			return s, true
		}
	}
	return models.SlotProduccion{}, false
}

func (a *AlmacenMemoria) CrearSlotProduccion(s models.SlotProduccion) models.SlotProduccion {
	s.ID = SlotProduccionID
	SlotProduccionID++
	SlotsProduccion = append(SlotsProduccion, s)
	return s
}

func (a *AlmacenMemoria) ActualizarSlotProduccion(id int, datos models.SlotProduccion) (models.SlotProduccion, bool) {
	for i, s := range SlotsProduccion {
		if s.ID == id {
			datos.ID = id
			SlotsProduccion[i] = datos
			return datos, true
		}
	}
	return models.SlotProduccion{}, false
}

func (a *AlmacenMemoria) BorrarSlotProduccion(id int) bool {
	for i, s := range SlotsProduccion {
		if s.ID == id {
			SlotsProduccion = append(SlotsProduccion[:i], SlotsProduccion[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// CLIENTES
// =========================================================

func (a *AlmacenMemoria) ListarClientes() []models.Cliente {
	return Clientes
}

func (a *AlmacenMemoria) BuscarClientePorID(id int) (models.Cliente, bool) {
	for _, c := range Clientes {
		if c.ID == id {
			return c, true
		}
	}
	return models.Cliente{}, false
}

func (a *AlmacenMemoria) CrearCliente(c models.Cliente) models.Cliente {
	c.ID = ClienteID
	ClienteID++
	Clientes = append(Clientes, c)
	return c
}

func (a *AlmacenMemoria) ActualizarCliente(id int, datos models.Cliente) (models.Cliente, bool) {
	for i, c := range Clientes {
		if c.ID == id {
			datos.ID = id
			Clientes[i] = datos
			return datos, true
		}
	}
	return models.Cliente{}, false
}

func (a *AlmacenMemoria) BorrarCliente(id int) bool {
	for i, c := range Clientes {
		if c.ID == id {
			Clientes = append(Clientes[:i], Clientes[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// SEGUIMIENTOS PEDIDO
// =========================================================

func (a *AlmacenMemoria) ListarSeguimientosPedido() []models.SeguimientoPedido {
	return SeguimientosPedido
}

func (a *AlmacenMemoria) BuscarSeguimientoPedidoPorID(id int) (models.SeguimientoPedido, bool) {
	for _, s := range SeguimientosPedido {
		if s.ID == id {
			return s, true
		}
	}
	return models.SeguimientoPedido{}, false
}

func (a *AlmacenMemoria) CrearSeguimientoPedido(s models.SeguimientoPedido) models.SeguimientoPedido {
	s.ID = SeguimientoPedidoID
	SeguimientoPedidoID++
	SeguimientosPedido = append(SeguimientosPedido, s)
	return s
}

func (a *AlmacenMemoria) ActualizarSeguimientoPedido(id int, datos models.SeguimientoPedido) (models.SeguimientoPedido, bool) {
	for i, s := range SeguimientosPedido {
		if s.ID == id {
			datos.ID = id
			SeguimientosPedido[i] = datos
			return datos, true
		}
	}
	return models.SeguimientoPedido{}, false
}

func (a *AlmacenMemoria) BorrarSeguimientoPedido(id int) bool {
	for i, s := range SeguimientosPedido {
		if s.ID == id {
			SeguimientosPedido = append(SeguimientosPedido[:i], SeguimientosPedido[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// RECLAMOS
// =========================================================

func (a *AlmacenMemoria) ListarReclamos() []models.Reclamo {
	return Reclamos
}

func (a *AlmacenMemoria) BuscarReclamoPorID(id int) (models.Reclamo, bool) {
	for _, r := range Reclamos {
		if r.ID == id {
			return r, true
		}
	}
	return models.Reclamo{}, false
}

func (a *AlmacenMemoria) CrearReclamo(r models.Reclamo) models.Reclamo {
	r.ID = ReclamoID
	ReclamoID++
	Reclamos = append(Reclamos, r)
	return r
}

func (a *AlmacenMemoria) ActualizarReclamo(id int, datos models.Reclamo) (models.Reclamo, bool) {
	for i, r := range Reclamos {
		if r.ID == id {
			datos.ID = id
			Reclamos[i] = datos
			return datos, true
		}
	}
	return models.Reclamo{}, false
}

func (a *AlmacenMemoria) BorrarReclamo(id int) bool {
	for i, r := range Reclamos {
		if r.ID == id {
			Reclamos = append(Reclamos[:i], Reclamos[i+1:]...)
			return true
		}
	}
	return false
}

var _ Almacen = (*AlmacenMemoria)(nil)
