package service

import "errors"

var (
	ErrPedidoNoEncontrado    = errors.New("pedido no encontrado")
	ErrClienteNoEncontrado   = errors.New("cliente no encontrado")
	ErrReclamoNoEncontrado   = errors.New("reclamo no encontrado")
	ErrCredencialesInvalidas = errors.New("credenciales inválidas")
	ErrEmailEnUso            = errors.New("el email ya está en uso")
)
