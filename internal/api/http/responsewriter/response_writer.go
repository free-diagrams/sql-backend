package responsewriter

import (
	"context"
	"encoding/json"
	"github.com/free-diagrams/sql-backend/internal/domain/entity"
	"github.com/free-diagrams/sql-backend/internal/domain/errors"
	"github.com/free-diagrams/sql-backend/pkg/httpresp"
	"github.com/free-diagrams/sql-backend/pkg/loclzr"
	"github.com/free-diagrams/sql-backend/pkg/logger"
	"net/http"
)

type ResponseWriter struct {
	localizer *loclzr.Localizer
	log       *logger.Logger
}

func New(localizer *loclzr.Localizer, log *logger.Logger) *ResponseWriter {
	return &ResponseWriter{
		localizer: localizer,
		log:       log,
	}
}

func (w *ResponseWriter) Write(ctx context.Context, writer http.ResponseWriter, statusCode int, payload any) {
	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(statusCode)

	err := json.NewEncoder(writer).Encode(payload)
	if err != nil {
		w.log.Debug().Err(err).Msg("failed to encode response")
	}
}

func (w *ResponseWriter) WriteError(ctx context.Context, writer http.ResponseWriter, err error) {
	httpCode, messageID := errors.ErrorToHTTPStatusAndMessageID(err)

	w.Write(ctx, writer, httpCode, w.constructError(ctx, messageID, err))
}

func (w *ResponseWriter) constructError(ctx context.Context, messageID string, err error) *httpresp.Error {
	lang := ctx.Value(entity.ContextLang).(string)

	message, errLoc := w.localizer.Localize(lang, messageID)
	if errLoc != nil {
		w.log.Debug().Err(errLoc).Msg("failed to localize message")
		message = w.localizer.English(messageID)
	}

	return &httpresp.Error{
		Message: message,
		Error:   err.Error(),
	}
}
