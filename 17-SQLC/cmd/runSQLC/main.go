package main

import (
	"context"
	"database/sql"
	"hello_world/pos-go-expert/17-SQLC/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)
	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Backend courses", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "6756998a-5e8b-45ca-9315-aa6b85e0f981",
		Name:        "Backend updated",
		Description: sql.NullString{String: "Backend courses edit", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.ID, category.Name)
	}
}
