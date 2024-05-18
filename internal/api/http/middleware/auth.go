package middleware

import (
	"context"
	"github.com/free-diagrams/sql-backend/internal/domain/entity"
	"github.com/free-diagrams/sql-backend/internal/domain/errors"
	"github.com/go-chi/jwtauth"
	"net/http"
)

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		if token == "" {
			m.responseWriter.WriteError(r.Context(), w, errors.EmptyToken)
			return
		}

		user, err := m.userUseCase.GetUserByToken(r.Context(), token)
		if err != nil {
			m.responseWriter.WriteError(r.Context(), w, errors.InvalidToken)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), entity.ContextUser, user))

		next.ServeHTTP(w, r)
	})
}
