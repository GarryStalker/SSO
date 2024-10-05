package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {
	//иннициализация объекта конфига
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// TODO: инициализировать логгер

	// TODO: инициализировать приложение (арр)

	// TODO: запустить gRPC-сервер приложения
}
