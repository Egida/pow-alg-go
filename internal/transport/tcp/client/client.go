package client

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

type Client struct {
	handler handler
	log     *slog.Logger
	conn    net.Conn
}

func NewClient(addr string, log *slog.Logger, handler handler) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create TCP connection: %w", err)
	}
	return &Client{conn: conn, log: log, handler: handler}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Run(ctx context.Context) error {
	return c.handler.Handle(ctx, c.conn, c.conn)
}
