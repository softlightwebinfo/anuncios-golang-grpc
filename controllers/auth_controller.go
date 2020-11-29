package controllers

import (
	"cientosdeanuncios.com/backend/enums"
	"cientosdeanuncios.com/backend/libs"
	"cientosdeanuncios.com/backend/models"
	"cientosdeanuncios.com/backend/proto"
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type AuthController struct {
	DB   *sql.DB
	User int64
}

func (then AuthController) Auth(_ context.Context, rq *proto.AuthRequest) (*proto.AuthResponse, error) {
	if len(rq.GetUserName()) == 0 {
		return nil, status.Error(codes.NotFound, "El campo usuario esta vacio")
	}

	if len(rq.GetPassword()) == 0 {
		return nil, status.Error(codes.NotFound, "El campo contraseña esta vacio")
	}

	response := proto.AuthResponse{}
	user := proto.AuthUser{}
	var password string
	err := then.DB.QueryRow(
		"SELECT id, user_name, email, password, created_at, fk_role FROM users WHERE user_name=$1 or email=$1",
		rq.GetUserName(),
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&password,
		&user.CreatedAt,
		&user.Role,
	)

	if err != nil {
		return &response, status.Error(codes.NotFound, "El usuario no se ha encontrado")
	}

	if password != rq.GetPassword() {
		return &response, status.Error(codes.NotFound, "El usuario y/o contraseña son erroneos")
	}
	model := models.AuthModel{}
	token, err := model.CreateToken(models.User{Email: user.Email, Name: user.Name, ID: user.Id, Role: user.Role})

	if err != nil {
		return &response, status.Error(codes.Canceled, "No se ha podido generar el token a autenticación")
	}

	response.User = &user
	response.Token = token

	return &response, nil
}
func (then AuthController) Register(_ context.Context, rq *proto.AuthRegisterRequest) (*proto.AuthResponse, error) {
	if len(rq.GetName()) == 0 {
		return nil, status.Error(codes.NotFound, "El campo usuario esta vacio")
	}

	if len(rq.GetPassword()) == 0 {
		return nil, status.Error(codes.NotFound, "El campo contraseña esta vacio")
	}

	if len(rq.GetEmail()) == 0 {
		return nil, status.Error(codes.NotFound, "El campo email esta vacio")
	}

	response := proto.AuthResponse{}
	user := proto.AuthUser{}
	err := then.DB.QueryRow(
		"INSERT INTO users(user_name, email, password, fk_role) VALUES ($1, $2, $3, $4) returning id, user_name, email, fk_role, created_at",
		rq.GetName(),
		rq.GetEmail(),
		rq.GetPassword(),
		enums.User,
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return &response, status.Error(codes.NotFound, "No se ha podido crear el usuario")
	}

	model := models.AuthModel{}
	token, err := model.CreateToken(models.User{Email: user.Email, Name: user.Name, ID: user.Id, Role: user.Role})

	if err != nil {
		return &response, status.Error(codes.Canceled, "No se ha podido generar el token a autenticación")
	}

	response.User = &user
	response.Token = token

	return &response, nil
}
func (then AuthController) SendRecoveryAccount(_ context.Context, rq *proto.AuthRecoveryAccountRequest) (*proto.AuthRecoveryAccountResponse, error) {
	if len(rq.GetEmail()) == 0 {
		return nil, status.Error(codes.NotFound, "El campo email esta vacio")
	}

	response := proto.AuthRecoveryAccountResponse{}
	response.Success = false

	tokenGenerate := libs.Base64Encode(fmt.Sprintf("%s-%s", rq.GetEmail(), time.Now()))

	exec, err := then.DB.Exec(
		"UPDATE users set recovery_password=$1 where email=$2",
		tokenGenerate,
		rq.GetEmail(),
	)

	if err != nil {
		return &response, status.Error(codes.NotFound, "No se ha podido recuperar el usuario")
	}

	if row, err := exec.RowsAffected(); row == 0 || err != nil {
		return &response, status.Error(codes.NotFound, "El usuario no existe")
	}

	response.Success = true
	return &response, nil
}
func (then AuthController) RecoveryAccount(_ context.Context, rq *proto.AuthRecoveryAccountChangeRequest) (*proto.AuthRecoveryAccountResponse, error) {
	if len(rq.GetEmail()) == 0 {
		return nil, status.Error(codes.NotFound, "El campo email esta vacio")
	}

	if len(rq.GetToken()) == 0 {
		return nil, status.Error(codes.NotFound, "El campo token esta vacio")
	}

	if len(rq.GetPassword()) == 0 {
		return nil, status.Error(codes.NotFound, "El campo password esta vacio")
	}
	response := proto.AuthRecoveryAccountResponse{}
	response.Success = false

	// TODO: Mejor buscar el usuario si existe y despues actualizarlo? o hacemos que cuando de error el update...

	exec, err := then.DB.Exec(
		"UPDATE users set recovery_password=$1, password=$4 where email=$2 and recovery_password=$3",
		nil,
		rq.GetEmail(),
		rq.GetToken(),
		rq.GetPassword(),
	)

	if err != nil {
		return &response, status.Error(codes.NotFound, "No se ha podido recuperar el usuario")
	}

	if row, err := exec.RowsAffected(); row == 0 || err != nil {
		return &response, status.Error(codes.NotFound, "El usuario no existe")
	}

	response.Success = true
	return &response, nil
}
