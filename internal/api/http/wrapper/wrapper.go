package wrapper

import (
	"bytes"
	"github.com/free-diagrams/sql-backend/internal/api/http/responsewriter"
	"github.com/free-diagrams/sql-backend/internal/domain/entity"
	"github.com/free-diagrams/sql-backend/pkg/logger"
	"github.com/go-chi/chi/middleware"
	"io"
	"net/http"
	"time"
)

type Wrapper struct {
	responseWriter *responsewriter.ResponseWriter
	log            *logger.Logger
}

func New(responseWriter *responsewriter.ResponseWriter, log *logger.Logger) *Wrapper {
	return &Wrapper{
		responseWriter: responseWriter,
		log:            log,
	}
}

func (w *Wrapper) Wrap(controllerFunction func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(writer, request.ProtoMajor)

		var requestBody bytes.Buffer
		if request.Body != nil {
			_, err := io.Copy(&requestBody, request.Body)
			if err != nil {
				w.log.Error().Err(err).Msg("failed to read request body")
				return
			}
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				w.log.Error().Err(err).Msg("failed to close request body")
			}
		}(request.Body)

		clonedRequest := request.Clone(request.Context())
		clonedRequest.Body = io.NopCloser(bytes.NewReader(requestBody.Bytes()))

		requester := request.Context().Value(entity.ContextUser)
		requesterUserID := ""
		requesterUsername := ""
		if requester != nil {
			if u, ok := requester.(*entity.User); ok {
				requesterUserID = *u.ID
				requesterUsername = u.Username
			}
		}

		ctx := request.Context()
		clonedRequest = clonedRequest.WithContext(ctx)

		controllerError := controllerFunction(ww, request)
		if controllerError != nil {
			w.responseWriter.WriteError(ctx, ww, controllerError)
		}

		duration := time.Since(start)
		var requestBodyJSON []byte
		if requestBody.Len() > 0 {
			requestBodyJSON = requestBody.Bytes()
		} else {
			requestBodyJSON = []byte("null")
		}

		if ww.Status() >= http.StatusInternalServerError {
			w.log.Error().
				Err(controllerError).
				Str("method", request.Method).
				Str("path", request.URL.String()).
				Str("duration", duration.String()).
				RawJSON("request_body", requestBodyJSON).
				Int("request_body_length_bytes", requestBody.Len()).
				Int("response_body_length_bytes", ww.BytesWritten()).
				Int("response_status", ww.Status()).
				Str("user_agent", request.UserAgent()).
				Str("user_id", requesterUserID).
				Str("user_username", requesterUsername).
				Str("source_ip", request.RemoteAddr).
				Msgf("request handled with unexpected error: %d %s %s", ww.Status(), request.Method, request.URL.Path)
		} else if ww.Status() >= http.StatusBadRequest {
			w.log.Warn().
				Err(controllerError).
				Str("method", request.Method).
				Str("path", request.URL.String()).
				Str("duration", duration.String()).
				RawJSON("request_body", requestBodyJSON).
				Int("request_body_length_bytes", requestBody.Len()).
				Int("response_body_length_bytes", ww.BytesWritten()).
				Int("response_status", ww.Status()).
				Str("user_agent", request.UserAgent()).
				Str("user_id", requesterUserID).
				Str("user_username", requesterUsername).
				Str("source_ip", request.RemoteAddr).
				Msgf("request handled with error: %d %s %s", ww.Status(), request.Method, request.URL.Path)
		} else {
			w.log.Info().
				Str("method", request.Method).
				Str("path", request.URL.String()).
				Str("duration", duration.String()).
				RawJSON("request_body", requestBodyJSON).
				Int("request_body_length_bytes", requestBody.Len()).
				Int("response_body_length_bytes", ww.BytesWritten()).
				Int("response_status", ww.Status()).
				Str("user_agent", request.UserAgent()).
				Str("user_id", requesterUserID).
				Str("user_username", requesterUsername).
				Str("source_ip", request.RemoteAddr).
				Msgf("request handled succussfully: %d %s %s", ww.Status(), request.Method, request.URL.Path)
		}
	}
}
