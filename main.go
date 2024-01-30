package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1749"
	dbname   = "postgres"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Создание таблицы, если она не существует
	createTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255),
		body TEXT
	);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/create-post", func(w http.ResponseWriter, r *http.Request) {
		var newPost Post
		err := json.NewDecoder(r.Body).Decode(&newPost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Вставляем данные в базу данных
		_, err = db.Exec("INSERT INTO posts (title, body) VALUES ($1, $2)", newPost.Title, newPost.Body)
		if err != nil {
			log.Println("Ошибка при создании поста:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Создан новый пост: ID=%d, Title=%s, Body=%s\n", newPost.ID, newPost.Title, newPost.Body)
		w.WriteHeader(http.StatusCreated)
	})

	http.HandleFunc("/get-posts", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM posts")
		if err != nil {
			log.Println("Ошибка при получении постов:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var posts []Post
		for rows.Next() {
			var post Post
			err := rows.Scan(&post.ID, &post.Title, &post.Body)
			if err != nil {
				log.Println("Ошибка при сканировании записей:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			posts = append(posts, post)
		}

		log.Println("Получены все посты.")
		json.NewEncoder(w).Encode(posts)
	})

	http.HandleFunc("/update-post", func(w http.ResponseWriter, r *http.Request) {
		var updatedPost Post
		err := json.NewDecoder(r.Body).Decode(&updatedPost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Обновляем данные в базе данных
		_, err = db.Exec("UPDATE posts SET title=$1, body=$2 WHERE id=$3", updatedPost.Title, updatedPost.Body, updatedPost.ID)
		if err != nil {
			log.Println("Ошибка при обновлении поста:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Пост обновлен: ID=%d, Title=%s, Body=%s\n", updatedPost.ID, updatedPost.Title, updatedPost.Body)
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/delete-post", func(w http.ResponseWriter, r *http.Request) {
		postID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		// Удаляем запись из базы данных
		_, err = db.Exec("DELETE FROM posts WHERE id=$1", postID)
		if err != nil {
			log.Println("Ошибка при удалении поста:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Пост удален: ID=%d\n", postID)
		w.WriteHeader(http.StatusOK)
	})

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
