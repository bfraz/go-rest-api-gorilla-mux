package main

import (
	"fmt"
	"go-rest-api-course-v2/internal/comment"
	"go-rest-api-course-v2/internal/db"
	transportHttp "go-rest-api-course-v2/internal/transport/http"
)

// Run - is responsible for
// instantiation and startup
// of our go application
func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to db")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate db")
		return err
	}
	fmt.Println("successfully connected and pinged db")

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}
	/*
		fmt.Println(cmtService.PostComment(
			context.Background(),
			comment.Comment{
				Slug:   "manual-test",
				Author: "author",
				Body:   "Hello World",
			},
		))

		fmt.Println(cmtService.UpdateComment(
			context.Background(),
			"71c5d074-b6cf-11ec-b909-0242ac120002",
			comment.Comment{
				Slug:   "manual-test-updated",
				Author: "author-updated",
				Body:   "Hello World Updated",
			},
		))

		fmt.Println(cmtService.GetComment(
			context.Background(),
			"71c5d074-b6cf-11ec-b909-0242ac120002",
		))

		fmt.Println(cmtService.DeleteComment(
			context.Background(),
			"71c5d074-b6cf-11ec-b909-0242ac120002",
		))
	*/

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
