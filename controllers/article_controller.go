package controllers

import (
	"cientosdeanuncios.com/backend/proto"
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
