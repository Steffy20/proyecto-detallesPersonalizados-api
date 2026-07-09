package middleware

import (
	"context"
	"net/http"
	"strings"

	"proyecto-detallesPersonalizados-api/internal/service"
)

type claveContext string

const ClaveUsuarioID claveContext = "usuarioID"
const ClaveRolUsuario claveContext = "rolUsuario"

func Auth(auth *service.AuthService) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			encabezado := strings.TrimSpace(r.Header.Get("Authorization"))
			partes := strings.Fields(encabezado)

			if len(partes) != 2 || partes[0] != "Bearer" {
				responderNoAutorizado(w)
				return
			}

			claims, err := auth.ValidarTokenConClaims(partes[1])

			if err != nil {
				responderNoAutorizado(w)
				return
			}

			ctx := context.WithValue(r.Context(), ClaveUsuarioID, claims.UsuarioID)
			ctx = context.WithValue(ctx, ClaveRolUsuario, claims.Rol)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequiereRol(rolesPermitidos ...string) func(http.Handler) http.Handler {
	permitidos := make(map[string]bool, len(rolesPermitidos))
	for _, rol := range rolesPermitidos {
		permitidos[strings.ToLower(strings.TrimSpace(rol))] = true
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rol, _ := r.Context().Value(ClaveRolUsuario).(string)

			if !permitidos[strings.ToLower(strings.TrimSpace(rol))] {
				responderProhibido(w)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func responderProhibido(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)

	_, _ = w.Write([]byte(`{"error":"No tienes permisos para acceder a este recurso"}`))
}

func responderNoAutorizado(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	_, _ = w.Write([]byte(`{"error":"Token inexistente o inválido"}`))
}
