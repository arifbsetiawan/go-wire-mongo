package config

import "os"

type App struct {
	PORT     string
	MongoURI string
	MongoDB  string
}

func InitConfig() App {
	var app App
	app.PORT = os.Getenv("APP_PORT")
	app.MongoURI = os.Getenv("MONGO_URI")
	app.MongoDB = os.Getenv("MONGO_DB")

	return app
}
