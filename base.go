package cli_base

import (
	"log/slog"
	"os"
)

func init() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)
}
