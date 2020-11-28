package main

import (
	"cientosdeanuncios.com/backend/controllers"
	"cientosdeanuncios.com/backend/proto"
	"cientosdeanuncios.com/backend/structs"
	"database/sql"
	"google.golang.org/grpc"
	"os"
)

func main() {
	app := structs.App{}
	app.
		LoadEnv().
		Initialize(
			os.Getenv("APP_DB_HOST"),
			os.Getenv("APP_DB_USERNAME"),
			os.Getenv("APP_DB_PASSWORD"),
			os.Getenv("APP_DB_NAME"),
			os.Getenv("APP_DB_PORT"),
		).
		DB().
		GRPC(os.Getenv("GRPC_ADDR"), func(server *grpc.Server, db *sql.DB) {
			proto.RegisterUserServiceServer(server, controllers.UserController{DB: db})
			proto.RegisterArticleServiceServer(server, controllers.ArticleController{DB: db})
		})
}
