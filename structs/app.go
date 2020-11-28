package structs

import (
	"cientosdeanuncios.com/backend/libs"
	"database/sql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type App struct {
	database Database
}

func (then *App) GetDb() *sql.DB {
	return then.database.Connection
}

func (then *App) Initialize(host string, userName string, password string, name string, port string) *App {
	then.database.New(host, userName, password, name, port)
	return then
}

func (then *App) LoadEnv() *App {
	_ = godotenv.Load()
	return then
}

func (then *App) DB() *App {
	then.database.Connect()
	return then
}

func (then *App) GRPC(port string, callback func(server *grpc.Server, db *sql.DB)) {
	server := libs.GrpcNewServer()
	callback(server, then.GetDb())
	listener := libs.GrpcGetNetListener(port)
	libs.GrpcListenerServer(server, listener)
	defer then.GetDb().Close()
}
