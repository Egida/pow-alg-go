package get_quote

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/ivasilkov/pow-alg-go/internal/pow/hashcash"
	"github.com/ivasilkov/pow-alg-go/internal/transport/tcp"
)

type pow interface {
	Compute(resource string) (hashcash.Header, error)
}

type Handler struct {
	log *slog.Logger
	pow pow
}

func NewHandler(log *slog.Logger, pow pow) *Handler {
	return &Handler{log: log, pow: pow}
}

func (h *Handler) Handle(ctx context.Context, r io.Reader, w io.Writer) error {
	header, err := h.pow.Compute(h.getResource())
	if err != nil {
		return fmt.Errorf("failed to compute header: %w", err)
	}

	err = tcp.WriteString(w, header.String())
	if err != nil {
		return fmt.Errorf("failed to write data: %w", err)
	}

	h.log.Info("successfully sent header", "str", header.String(), "hash", header.Hash())

	quote, err := tcp.ReadString(r)
	if err != nil {
		return fmt.Errorf("failed to read quote: %w", err)
	}

	h.log.Info("successfully get quote", "quote", quote)

	return nil
}

func (h *Handler) getResource() string {
	return "any-client-data"
}
