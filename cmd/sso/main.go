package main

//go run cmd/sso/main.go --config=./config/local.yaml
import (
	"log/slog"
	"os"
	"sso/internal/app"
	"sso/internal/config"
	"sso/internal/lib/logger/handlers/slogpretty"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//иннициализация объекта конфига
	cfg := config.MustLoad()

	//инициализация логгера
	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("config", cfg))

	//инициализация приложения (арр)
	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	application.GRPCServer.MustRun()
	//application.gRPCServer.MustRun()

	// TODO: запустить gRPC-сервер приложения
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
