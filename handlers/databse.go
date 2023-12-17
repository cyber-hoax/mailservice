package handlers

import (
	models "blog_service/models"
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func init() {
	// Replace with your MySQL database connection details
	dbURI := "username:password@tcp(localhost:3306)/blog_db"
	var err error
	db, err = sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Check if the connection is alive
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	fmt.Println("Connected to MySQL database")
}

// CreateBlogPost creates a new blog post in the database.
func CreateBlogPost(post models.BlogPost) (int64, error) {
	result, err := db.Exec("INSERT INTO blog_posts (title, content) VALUES (?, ?)", post.Title, post.Content)
	if err != nil {
		log.Println("Error creating blog post:", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return 0, err
	}

	return id, nil
}

// GetBlogPosts retrieves all blog posts from the database.
func GetBlogPosts() ([]models.BlogPost, error) {
	rows, err := db.Query("SELECT id, title, content FROM blog_posts")
	if err != nil {
		log.Println("Error querying blog posts:", err)
		return nil, err
	}
	defer rows.Close()

	var posts []models.BlogPost

	for rows.Next() {
		var post models.BlogPost
		if err := rows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
			log.Println("Error scanning blog post:", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// GetBlogPostByID retrieves a specific blog post by ID from the database.
func GetBlogPostByID(id int) (models.BlogPost, error) {
	var post models.BlogPost
	err := db.QueryRow("SELECT id, title, content FROM blog_posts WHERE id = ?", id).
		Scan(&post.ID, &post.Title, &post.Content)

	if err != nil {
		log.Println("Error getting blog post by ID:", err)
		return models.BlogPost{}, err
	}

	return post, nil
}

// UpdateBlogPost updates an existing blog post in the database.
func UpdateBlogPost(post models.BlogPost) error {
	_, err := db.Exec("UPDATE blog_posts SET title = ?, content = ? WHERE id = ?", post.Title, post.Content, post.ID)
	if err != nil {
		log.Println("Error updating blog post:", err)
		return err
	}

	return nil
}

// DeleteBlogPost deletes a blog post from the database by ID.
func DeleteBlogPost(id int) error {
	_, err := db.Exec("DELETE FROM blog_posts WHERE id = ?", id)
	if err != nil {
		log.Println("Error deleting blog post:", err)
		return err
	}

	return nil
}
