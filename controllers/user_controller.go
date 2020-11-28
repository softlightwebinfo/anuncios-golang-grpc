package controllers

import (
	"cientosdeanuncios.com/backend/proto"
	"context"
	"database/sql"
)

type UserController struct {
	DB *sql.DB
}

func (then UserController) GetUsers(context.Context, *proto.GetUsersRequest) (*proto.GetUsersResponse, error) {
	print("HELLO, que tal")
	return nil, nil
}
