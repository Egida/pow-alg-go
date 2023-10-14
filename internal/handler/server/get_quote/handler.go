package get_quote

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/ivasilkov/pow-alg-go/internal/pow/hashcash"
	"github.com/ivasilkov/pow-alg-go/internal/transport/tcp"
)

type storage interface {
	GetRandomQuote(ctx context.Context) (string, error)
}

type verifier interface {
	Verify(h hashcash.Header) error
}

type Handler struct {
	log      *slog.Logger
	storage  storage
	verifier verifier
}

func NewHandler(log *slog.Logger, storage storage, verifier verifier) *Handler {
	return &Handler{log: log, storage: storage, verifier: verifier}
}

func (h *Handler) Handle(ctx context.Context, r io.Reader, w io.Writer) error {
	in, err := tcp.ReadString(r)
	if err != nil {
		return fmt.Errorf("failed to read data: %w", err)
	}

	header, err := hashcash.ParseHeader(in)
	if err != nil {
		return fmt.Errorf("failed to parse header: %w", err)
	}

	err = h.verifier.Verify(header)
	if err != nil {
		return fmt.Errorf("failed to verify header: %w", err)
	}

	quote, err := h.storage.GetRandomQuote(ctx)
	if err != nil {
		return fmt.Errorf("failed to get random quote: %w", err)
	}

	err = tcp.WriteString(w, quote)
	if err != nil {
		return fmt.Errorf("failed to write data: %w", err)
	}

	h.log.Info("successfully handled")
	return nil
}
