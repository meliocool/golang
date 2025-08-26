package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-database/entity"
	"strconv"
	"strings"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlScript := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	result, err := repository.DB.ExecContext(ctx, sqlScript, comment.Email, comment.Comment)

	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	sqlScript := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, sqlScript, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id" + strconv.Itoa(int(id)) + "Not Found!")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	sqlScript := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, sqlScript)
	var comments []entity.Comment
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repository *commentRepositoryImpl) Delete(ctx context.Context, id int32) (bool, error) {
	sqlScript := "DELETE FROM comments WHERE id = ?"
	res, err := repository.DB.ExecContext(ctx, sqlScript, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}

func (repository *commentRepositoryImpl) Update(ctx context.Context, id int32, newComment string) (entity.Comment, error) {
	if id <= 0 {
		return entity.Comment{}, errors.New("Id cannot be 0!")
	}
	updateComment := strings.TrimSpace(newComment)
	if updateComment == "" {
		return entity.Comment{}, errors.New("Comment cannot be empty!")
	}
	if len(updateComment) > 255 {
		return entity.Comment{}, errors.New("Comment is too long")
	}

	sqlScript := "UPDATE comments SET comment = ? WHERE id = ?"
	res, err := repository.DB.ExecContext(ctx, sqlScript, updateComment, id)

	if err != nil {
		return entity.Comment{}, fmt.Errorf("update comment id=%d: %w", id, err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return entity.Comment{}, fmt.Errorf("rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return entity.Comment{}, sql.ErrNoRows
	}

	sqlSelect := "SELECT id, email, comment FROM comments WHERE id = ?"
	var out entity.Comment
	if err := repository.DB.QueryRowContext(ctx, sqlSelect, id).Scan(&out.Id, &out.Email, &out.Comment); err != nil {
		return entity.Comment{}, fmt.Errorf("select updated comment id=%d: %w", id, err)
	}
	return out, nil

}
