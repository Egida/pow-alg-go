package get_quote

import (
	"bufio"
	"context"
	"io"
	"log/slog"

	"github.com/ivasilkov/pow-alg-go/internal/pow/hashcash"
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

func (h *Handler) Handle(ctx context.Context, r io.Reader, w io.Writer) {
	header, err := h.pow.Compute(h.getResource())
	if err != nil {
		h.log.Error("failed to compute header: %w", "err", err.Error())
		return
	}

	_, err = w.Write([]byte(header.String() + "\n"))
	if err != nil {
		h.log.Error("failed to write data", "err", err.Error())
		return
	}
	h.log.Info("successfully sent header", "str", header.String(), "hash", header.Hash())

	quote, err := bufio.NewReader(r).ReadString('\n')
	if err != nil {
		h.log.Error("failed to read quote", "err", err.Error(), "quote", quote)
		return
	}

	quote = quote[:len(quote)-1]
	h.log.Info("successfully get quote", "quote", quote)
}

func (h *Handler) getResource() string {
	return "any-client-data"
}
