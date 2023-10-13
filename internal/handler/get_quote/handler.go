package get_quote

import (
	"bufio"
	"context"
	"io"
	"log/slog"

	"github.com/ivasilkov/pow-alg-go/internal/pow/hashcash"
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

func NewHandler(log *slog.Logger, storage storage) *Handler {
	return &Handler{log: log, storage: storage}
}

func (h *Handler) Handle(ctx context.Context, r io.Reader, w io.Writer) {
	in, err := bufio.NewReader(r).ReadString('\n')
	if err != nil {
		h.log.Error("failed to read data", "err", err.Error())
		return
	}

	header, err := hashcash.ParseHeader(in)
	if err != nil {
		h.log.Error("failed to verify header", "err", err.Error())
		return
	}

	err = h.verifier.Verify(header)
	if err != nil {
		h.log.Error("failed to verify header", "err", err.Error())
		return
	}

	quote, err := h.storage.GetRandomQuote(ctx)
	if err != nil {
		h.log.Error("failed to get random quote", "err", err.Error())
		return
	}

	_, err = w.Write([]byte(quote))
	if err != nil {
		h.log.Error("failed to write data", "err", err.Error())
		return
	}
	h.log.Info("successfully handled")
}
