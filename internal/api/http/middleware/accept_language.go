package middleware

import (
	"context"
	"github.com/free-diagrams/sql-backend/internal/domain/entity"
	"net/http"
)

func (m *Middleware) AcceptLanguageMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		supportedLanguages := []string{"en-US", "ru-RU"}
		defaultLanguage := "en-US"
		acceptLang := r.Header.Get("Accept-Language")

		for _, supportedLanguage := range supportedLanguages {
			if supportedLanguage == acceptLang {
				r = r.WithContext(context.WithValue(r.Context(), entity.ContextLang, acceptLang))
				next.ServeHTTP(w, r)
				return
			}
		}

		r = r.WithContext(context.WithValue(r.Context(), entity.ContextLang, defaultLanguage))
		next.ServeHTTP(w, r)
		return
	})
}
