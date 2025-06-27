package db

import (
	"context"
	"database/sql"
	"fmt"
	"go-backend/internal/store"
	"math/rand"
)

func Seed(store store.Storage, db *sql.DB) error {
	ctx := context.Background()

	users := generateUsers(50)
	posts := generatePosts(100, users)
	comments := generateComments(200, users, posts)

	// Insert users
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			tx.Rollback()
			return fmt.Errorf("Error creating user: %v", err)
		}
	}
	tx.Commit()

	// Insert posts
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			return fmt.Errorf("Error creating post: %v", err)
		}
	}

	// Insert comments
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			return fmt.Errorf("Error creating comment: %v", err)
		}
	}

	fmt.Println("Seeding complete")
	return nil
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := range num {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			RoleID: 1,
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := range num {
		userID := rand.Intn(len(users)) + 1

		posts[i] = &store.Post{
			UserID:  int64(userID),
			Title:   titles[rand.Intn(len(titles))],
			Content: titles[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := range num {
		userID := rand.Intn(len(users)) + 1
		postID := rand.Intn(len(posts)) + 1
		
		cms[i] = &store.Comment{
			PostID:  int64(postID),
			UserID:  int64(userID),
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
