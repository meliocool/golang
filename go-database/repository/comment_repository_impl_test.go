package repository

import (
	"context"
	"fmt"
	go_database "go-database"
	"go-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repo@gmail.com",
		Comment: "Test Repo",
	}

	res, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 22)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}

func TestDelete(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.Delete(ctx, 22)
	if err != nil {
		panic(err)
	}
	if comments {
		fmt.Println("Comments with id", 22, "Deleted Successfully!")
	} else {
		fmt.Println("Deletion Failed!")
	}
}

func TestUpdate(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.Update(ctx, 23, "This is an updated Comment!")
	if err != nil {
		panic(err)
	}
	fmt.Println(comments)
}
