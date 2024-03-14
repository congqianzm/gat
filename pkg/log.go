package pkg

import (
	"io"
	"log/slog"
	"os"
)

func CreateLogger() *slog.Logger {
	f, err := os.OpenFile("logs/http.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil
	}
	defer func() {
		err := f.Close()
		if err != nil {
			return
		}
	}()

	multiWriter := io.MultiWriter(os.Stdout, f)

	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}
	return slog.New(slog.NewTextHandler(multiWriter, &opts))
}

var Logger = CreateLogger()
