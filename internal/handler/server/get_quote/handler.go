package get_quote

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/ivasilkov/pow-alg-go/internal/pow/hashcash"
	"github.com/ivasilkov/pow-alg-go/internal/transport/tcp"
)

type quoteStorage interface {
	GetRandomQuote(ctx context.Context) (string, error)
}

type hashStorage interface {
	AddUnique(ctx context.Context, key, value string) error
}

type verifier interface {
	Verify(h hashcash.Header) error
}

type Handler struct {
	log          *slog.Logger
	quoteStorage quoteStorage
	hashStorage  hashStorage
	verifier     verifier
}

func NewHandler(log *slog.Logger, storage quoteStorage, hashStorage hashStorage, verifier verifier) *Handler {
	return &Handler{log: log, quoteStorage: storage, hashStorage: hashStorage, verifier: verifier}
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

	err = h.hashStorage.AddUnique(ctx, header.Hash(), "0")
	if err != nil {
		return fmt.Errorf("failed to save hash: %w", err)
	}

	quote, err := h.quoteStorage.GetRandomQuote(ctx)
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
