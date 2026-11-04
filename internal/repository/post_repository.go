package repository

import (
	"context"
	"fmt"
	"time"
	"website/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

type PostRepository struct {
	db DBTX
}

func NewPostRepository(db DBTX) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(ctx context.Context, title, content string) (int64, error) {
	query := `INSERT INTO posts (title, content, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int64
	now := time.Now()

	err := r.db.QueryRow(ctx, query, title, content, now, now).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Error on creating post: %w", err)
	}

	return id, nil
}

func (r *PostRepository) GetAll(ctx context.Context) ([]models.Post, error) {
	query := `SELECT * FROM posts ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving all posts: %w", err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post

		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("Error in reading a row: %w", err)
		}

		posts = append(posts, p)
	}

	return posts, nil
}
