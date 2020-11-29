package controllers

import (
	"cientosdeanuncios.com/backend/proto"
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ContactController struct {
	DB   *sql.DB
	User int64
}

func (then ContactController) CreateContact(_ context.Context, rq *proto.CreateContactRequest) (*proto.ContactSuccessResponse, error) {
	response := proto.ContactSuccessResponse{}
	response.Success = false

	if len(rq.GetEmail()) == 0 {
		return &response, status.Error(codes.NotFound, "El campo email esta vacio")
	}

	if len(rq.GetName()) == 0 {
		return &response, status.Error(codes.NotFound, "El campo nombre esta vacio")
	}

	if len(rq.GetMessage()) == 0 {
		return &response, status.Error(codes.NotFound, "El campo mensaje esta vacio")
	}

	if len(rq.GetPhone()) == 0 {
		return &response, status.Error(codes.NotFound, "El campo telefono esta vacio")
	}

	_, err := then.DB.Exec(
		"INSERT INTO contacts(name, email, phone, message) VALUES ($1, $2, $3, $4)",
		rq.GetName(),
		rq.GetEmail(),
		rq.GetPhone(),
		rq.GetMessage(),
	)

	if err != nil {
		return &response, status.Error(codes.NotFound, "No se ha podido enviar el email")
	}

	response.Success = true

	return &response, nil
}
