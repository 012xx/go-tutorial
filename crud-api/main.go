package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// 映画の構造体
type Movie struct {
	ID       string    `json:"id"`       // 映画のID
	Isbn     string    `json:"isbn"`     // ISBN番号
	Title    string    `json:"title"`    // タイトル
	Director *Director `json:"director"` // 監督情報
}

// 監督の構造体
type Director struct {
	Firstname string `json:"firstname"` // 監督の名
	Lastname  string `json:"lastname"`  // 監督の姓
}

var movies []Movie // 映画のスライス

// 映画一覧を取得するハンドラ関数
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// 映画を削除するハンドラ関数
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// 特定の映画を取得するハンドラ関数
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

// 新しい映画を作成するハンドラ関数
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// 映画情報を更新するハンドラ関数
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var updatedMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
			updatedMovie.ID = params["id"]
			movies = append(movies, updatedMovie)
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

func main() {
	r := mux.NewRouter()

	// 映画データの初期化
	movies = append(movies, Movie{ID: "1", Isbn: "123456", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "123457", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	// ルーティングの設定とハンドラ関数の紐付け
	r.HandleFunc("/movies", getMovies).Methods("GET")           // 映画一覧を取得するエンドポイント
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")       // 特定の映画を取得するエンドポイント
	r.HandleFunc("/movies", createMovie).Methods("POST")        // 映画を作成するエンドポイント
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") // 映画を削除するエンドポイント
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")    // 映画情報を更新するエンドポイント

	fmt.Println("Server is running on port 8000\n") // サーバー起動メッセージ
	log.Fatal(http.ListenAndServe(":8000", r))      // サーバーの起動とエラーハンドリング
}
