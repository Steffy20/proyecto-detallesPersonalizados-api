--MODULO 1
CREATE TABLE pedido (
    id INTEGER PRIMARY KEY,
    mensaje TEXT NOT NULL,
    estado TEXT NOT NULL
);

CREATE TABLE personalizacion (
    id INTEGER PRIMARY KEY,
    pedido_id INTEGER NOT NULL,
    mensaje TEXT NOT NULL,
    color TEXT NOT NULL
);

CREATE TABLE producto_personalizado (
    id INTEGER PRIMARY KEY,
    pedido_id INTEGER NOT NULL,
    nombre TEXT NOT NULL,
    cantidad INTEGER NOT NULL,
    precio REAL NOT NULL
);

--MODULO 2

CREATE TABLE SolicitudUrgente (
    id INTEGER PRIMARY KEY,
    cliente TEXT NOT NULL,
    descripcion TEXT NOT NULL,
    fecha_requerida TEXT NOT NULL,
    estado TEXT NOT NULL
);

CREATE TABLE AgendaProduccion (
    id INTEGER PRIMARY KEY,
    fecha TEXT NOT NULL,
    responsable TEXT NOT NULL,
    estado TEXT NOT NULL
);

CREATE TABLE slotProduccion (
    id INTEGER PRIMARY KEY,
    agenda_id INTEGER NOT NULL,
    capacidad_maxima INTEGER NOT NULL,
    pedidos_asignados INTEGER NOT NULL
);

--MODULO 3

CREATE TABLE cliente (
    id INTEGER PRIMARY KEY,
    nombre TEXT NOT NULL,
    telefono TEXT NOT NULL
);

CREATE TABLE Reclamo (
    id INTEGER PRIMARY KEY,
    cliente_id INTEGER NOT NULL,
    pedido_id INTEGER NOT NULL,
    descripcion TEXT NOT NULL,
    estado TEXT NOT NULL
);

CREATE TABLE SeguimientoPedido (
    id INTEGER PRIMARY KEY,
    pedido_id INTEGER NOT NULL,
    estado TEXT NOT NULL,
    fecha_estado TEXT NOT NULL
);

