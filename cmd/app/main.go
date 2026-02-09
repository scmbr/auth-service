package main

import "github.com/scmbr/auth-service/internal/app"

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
