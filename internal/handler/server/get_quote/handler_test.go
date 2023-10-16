package get_quote_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/ivasilkov/pow-alg-go/internal/handler/server/get_quote"
	"github.com/ivasilkov/pow-alg-go/internal/handler/server/get_quote/mocks"
	"github.com/ivasilkov/pow-alg-go/internal/pow/hashcash"
)

func TestServerGetQuoteHandler(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

type handlerSuite struct {
	suite.Suite
	ctx          context.Context
	quoteStorage *mocks.QuoteStorage
	hashStorage  *mocks.HashStorage
	verifier     *mocks.Verifier
	h            *get_quote.Handler
}

func (s *handlerSuite) SetupTest() {
	s.ctx = context.Background()
	s.quoteStorage = mocks.NewQuoteStorage(s.T())
	s.hashStorage = mocks.NewHashStorage(s.T())
	s.verifier = mocks.NewVerifier(s.T())
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))
	s.h = get_quote.NewHandler(l, s.quoteStorage, s.hashStorage, s.verifier)
}

func (s *handlerSuite) TestSuccess() {
	now := time.Unix(100500, 0)
	h := hashcash.NewHeader(20, "resource", now)

	in := bytes.NewBufferString(fmt.Sprintf("%s\n", h.String()))
	out := bytes.NewBuffer([]byte{})

	s.verifier.EXPECT().Verify(h).Return(nil).Once()
	s.hashStorage.EXPECT().AddUnique(s.ctx, h.Hash(), "0").Return(nil).Once()
	s.quoteStorage.EXPECT().GetRandomQuote(s.ctx).Return("quote", nil).Once()

	err := s.h.Handle(s.ctx, in, out)

	s.Require().NoError(err)
	s.Require().Equal("quote\n", out.String())
}

func (s *handlerSuite) TestInvalidInput() {
	in := bytes.NewBufferString("invalid\n")
	out := bytes.NewBuffer([]byte{})

	err := s.h.Handle(s.ctx, in, out)

	s.Require().Error(err)
	s.Require().ErrorContains(err, "failed to parse header")
}

func (s *handlerSuite) TestVerifyError() {
	now := time.Unix(100500, 0)
	h := hashcash.NewHeader(20, "resource", now)

	in := bytes.NewBufferString(fmt.Sprintf("%s\n", h.String()))
	out := bytes.NewBuffer([]byte{})

	verifierErr := errors.New("verifier error")
	s.verifier.EXPECT().Verify(h).Return(verifierErr).Once()

	err := s.h.Handle(s.ctx, in, out)

	s.Require().ErrorIs(err, verifierErr)
}

func (s *handlerSuite) TestHashStorageError() {
	now := time.Unix(100500, 0)
	h := hashcash.NewHeader(20, "resource", now)

	in := bytes.NewBufferString(fmt.Sprintf("%s\n", h.String()))
	out := bytes.NewBuffer([]byte{})

	storageErr := errors.New("hash storage error")
	s.verifier.EXPECT().Verify(h).Return(nil).Once()
	s.hashStorage.EXPECT().AddUnique(s.ctx, h.Hash(), "0").Return(storageErr).Once()

	err := s.h.Handle(s.ctx, in, out)

	s.Require().ErrorIs(err, storageErr)
}

func (s *handlerSuite) TestQuoteStorageError() {
	now := time.Unix(100500, 0)
	h := hashcash.NewHeader(20, "resource", now)

	in := bytes.NewBufferString(fmt.Sprintf("%s\n", h.String()))
	out := bytes.NewBuffer([]byte{})

	storageErr := errors.New("quote storage error")
	s.verifier.EXPECT().Verify(h).Return(nil).Once()
	s.hashStorage.EXPECT().AddUnique(s.ctx, h.Hash(), "0").Return(nil).Once()
	s.quoteStorage.EXPECT().GetRandomQuote(s.ctx).Return("", storageErr).Once()

	err := s.h.Handle(s.ctx, in, out)

	s.Require().ErrorIs(err, storageErr)
}
