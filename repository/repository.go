package repository

import (
	"database/sql"
	"fmt"
	"sanbercode-mini-project/model"
)

func GetAllCategories(db *sql.DB) ([]model.Category, error) {
	sqlStatement := "SELECT id, name, created_at, modified_at FROM categories"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt, &c.ModifiedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func GetCategoryByID(db *sql.DB, id int) (model.Category, error) {
	var c model.Category
	sqlStatement := "SELECT id, name, created_at, modified_at FROM categories WHERE id = $1"
	err := db.QueryRow(sqlStatement, id).Scan(&c.ID, &c.Name, &c.CreatedAt, &c.ModifiedAt)
	if err != nil {
		return c, err
	}
	return c, nil
}

func InsertCategory(db *sql.DB, category *model.Category) error {
	err := db.QueryRow("INSERT INTO categories (name, created_at, modified_at) VALUES ($1, $2, $3) RETURNING id", 
		category.Name, category.CreatedAt, category.ModifiedAt).Scan(&category.ID)
	return err
}

func DeleteCategory(db *sql.DB, id int) error {
	sqlStatement := `DELETE FROM categories WHERE id = $1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return fmt.Errorf("id not found")
	}
	return nil
}

func UpdateCategory(db *sql.DB, id int, category *model.Category) error {
	sqlStatement := `UPDATE categories SET name = $1, modified_at = $2 WHERE id = $3`
	res, err := db.Exec(sqlStatement, category.Name, category.ModifiedAt, id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return fmt.Errorf("id not found")
	}
	category.ID = id
	return nil
}

func GetAllBooks(db *sql.DB) ([]model.Book, error) {
	sqlStatement := "SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, modified_at FROM books"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var b model.Book
		err = rows.Scan(&b.ID, &b.Title, &b.Description, &b.ImageUrl, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID, &b.CreatedAt, &b.ModifiedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func GetBooksByCategoryID(db *sql.DB, categoryID int) ([]model.Book, error) {
	sqlStatement := "SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, modified_at FROM books WHERE category_id = $1"
	rows, err := db.Query(sqlStatement, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var b model.Book
		err = rows.Scan(&b.ID, &b.Title, &b.Description, &b.ImageUrl, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID, &b.CreatedAt, &b.ModifiedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func InsertBook(db *sql.DB, book *model.Book) error {
	sqlStatement := `
	INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, modified_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	err := db.QueryRow(sqlStatement, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.CreatedAt, book.ModifiedAt).Scan(&book.ID)
	return err
}

func GetBookByID(db *sql.DB, id int) (model.Book, error) {
	var b model.Book
	sqlStatement := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, modified_at FROM books WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&b.ID, &b.Title, &b.Description, &b.ImageUrl, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID, &b.CreatedAt, &b.ModifiedAt)
	if err != nil {
		return model.Book{}, err
	}
	return b, nil
}

func DeleteBook(db *sql.DB, id int) error {
	sqlStatement := `DELETE FROM books WHERE id = $1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return fmt.Errorf("id not found")
	}
	return nil
}

func UpdateBook(db *sql.DB, id int, book *model.Book) error {
	sqlStatement := `UPDATE books SET title = $1, description = $2, image_url = $3, release_year = $4, 
		price = $5, total_page = $6, thickness = $7, category_id = $8, modified_at = $9 WHERE id = $10`
	res, err := db.Exec(sqlStatement, book.Title, book.Description, book.ImageUrl, book.ReleaseYear,
		book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.ModifiedAt, id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return fmt.Errorf("id not found")
	}
	book.ID = id
	return nil
}