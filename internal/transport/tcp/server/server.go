package server

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net"
)

type handler interface {
	Handle(ctx context.Context, r io.Reader, w io.Writer) error
}

type Server struct {
	log      *slog.Logger
	listener net.Listener
	handler  handler
}

func NewServer(addr string, log *slog.Logger, h handler) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create TCP listener: %w", err)
	}
	return &Server{listener: l, handler: h, log: log}, nil
}

func (s *Server) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			conn, err := s.listener.Accept()
			if err != nil {
				return fmt.Errorf("failed accept connection: %w", err)
			}
			s.log.Info("handle connection")
			go s.handleConn(ctx, conn)
		}
	}
}

func (s *Server) handleConn(ctx context.Context, conn net.Conn) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			err := s.handler.Handle(ctx, conn, conn)
			if err != nil {
				s.log.Error("failed to handle conn", "err", err)
				conn.Close()
				return
			}
		}
	}
}

func (s *Server) Close() error {
	return s.listener.Close()
}
