package controllers

import (
	"cientosdeanuncios.com/backend/proto"
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type ArticleController struct {
	DB *sql.DB
}

func (then ArticleController) GetArticles(context.Context, *proto.GetArticlesRequest) (*proto.GetArticlesResponse, error) {
	query, err := then.DB.Query("SELECT id, title, description, fk_user, created_at, updated_at, deleted_at from articles where deleted_at is null")
	if err != nil {
		return nil, err
	}
	response := proto.GetArticlesResponse{}
	for query.Next() {
		article := proto.Article{}
		err := query.Scan(
			&article.Id,
			&article.Title,
			&article.Description,
			&article.FkUser,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.DeletedAt,
		)
		if err != nil {
			fmt.Println("Error", err.Error())
		}
		response.Articles = append(response.Articles, &article)
	}
	return &response, nil
}

func (then ArticleController) GetArticle(_ context.Context, rq *proto.GetArticleRequest) (*proto.GetArticleResponse, error) {
	response := proto.GetArticleResponse{}
	article := proto.Article{}
	err := then.DB.QueryRow(
		"SELECT id, title, description, fk_user, created_at, updated_at, deleted_at from articles where id=$1 and deleted_at is null",
		rq.GetId(),
	).Scan(
		&article.Id,
		&article.Title,
		&article.Description,
		&article.FkUser,
		&article.CreatedAt,
		&article.UpdatedAt,
		&article.DeletedAt,
	)
	if err != nil { // scan will release the connection
		return nil, status.Error(codes.NotFound, "Ups, no se ha encontrado el anuncio")
	}
	response.Response = &article
	return &response, nil
}

func (then ArticleController) DeleteArticle(_ context.Context, rq *proto.DeleteArticleRequest) (*proto.SuccessResponse, error) {
	response := &proto.SuccessResponse{Success: false}
	exec, err := then.DB.Exec(
		"Update articles SET deleted_at=$1 where id=$2",
		time.Now(),
		rq.GetId(),
	)

	if err != nil {
		return response, status.Error(codes.Aborted, "Ups, no se ha podido modificar el anuncio")
	}

	if affect, err := exec.RowsAffected(); affect == 0 || err != nil {
		return response, status.Error(codes.NotFound, "Ups, no se ha encontrado el anuncio")
	}

	response.Success = true
	return response, nil
}

func (then ArticleController) SaveImagesArticle(_ context.Context, rq *proto.SaveImagesArticleRequest) (*proto.SuccessResponse, error) {
	response := &proto.SuccessResponse{Success: false}
	sqlStr := "INSERT INTO articles_images(fk_article, image) VALUES "
	var vals []interface{}
	count := 0

	for _, row := range rq.GetImages() {
		sqlStr += fmt.Sprintf("($%d, $%d),", count+1, count+2)
		vals = append(vals, rq.GetArticleId(), row)
		count += 2
	}

	stmt, _ := then.DB.Prepare(sqlStr[0 : len(sqlStr)-1])
	res, errStatement := stmt.Exec(vals...)

	if errStatement != nil {
		return response, status.Error(codes.NotFound, "Ups, no se ha encontrado el anuncio")
	}

	if affect, err := res.RowsAffected(); affect == 0 || err != nil {
		return response, status.Error(codes.NotFound, "Ups, no se ha encontrado el anuncio")
	}

	response.Success = true
	return response, nil
}
