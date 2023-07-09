package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`       // ユーザーID
	Name     string `json:"name"`     // ユーザー名
	Email    string `json:"email"`    // ユーザーのメールアドレス
	Age      int    `json:"age"`      // ユーザーの年齢
	IsActive bool   `json:"isActive"` // ユーザーのアクティブ状態
}

var users []User // ユーザーリスト

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User

	// リクエストボディからユーザー情報をデコードする
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// エラーが発生した場合は、不正なリクエストとしてBadRequestステータスを返す
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ユーザーIDを設定する
	user.ID = len(users) + 1

	// ユーザーリストにユーザーを追加する
	users = append(users, user)

	// レスポンスのContent-Typeヘッダを設定する
	w.Header().Set("Content-Type", "application/json")

	// レスポンスのステータスコードをCreatedに設定する
	w.WriteHeader(http.StatusCreated)

	// レスポンスにユーザー情報をJSON形式でエンコードして返す
	json.NewEncoder(w).Encode(user)
}


func GetUser(w http.ResponseWriter, r *http.Request) {
	// パラメータを取得する
	params := mux.Vars(r)

	// パラメータからIDを取得する
	// Atoi関数を使って文字列を数値に変換する
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// エラーが発生した場合は、不正なリクエストとしてBadRequestステータスを返す
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ユーザーを検索する
	for _, user := range users {
		if user.ID == id {
			// ユーザーが見つかった場合は、レスポンスのContent-Typeヘッダを設定し、
			// ユーザー情報をJSON形式でエンコードして返す
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	// ユーザーが見つからなかった場合は、NotFoundステータスを返す
	w.WriteHeader(http.StatusNotFound)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// パスパラメータからIDを取得する
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// エラーが発生した場合は、不正なリクエストとしてBadRequestステータスを返す
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// リクエストボディから更新情報をデコードする
	var updatedUser User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		// エラーが発生した場合は、不正なリクエストとしてBadRequestステータスを返す
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ユーザーリストをループして対象のユーザーを更新する
	for i, user := range users {
		if user.ID == id {
			// ユーザーが見つかった場合は、更新情報でユーザーを更新する
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			users[i].Age = updatedUser.Age
			users[i].IsActive = updatedUser.IsActive

			// レスポンスのContent-Typeヘッダを設定する
			w.Header().Set("Content-Type", "application/json")

			// 更新されたユーザー情報をJSON形式でエンコードして返す
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}

	// ユーザーが見つからなかった場合は、NotFoundステータスを返す
	w.WriteHeader(http.StatusNotFound)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// パスパラメータからIDを取得する
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// エラーが発生した場合は、不正なリクエストとしてBadRequestステータスを返す
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ユーザーリストをループして対象のユーザーを検索し削除する
	for i, user := range users {
		if user.ID == id {
			// ユーザーが見つかった場合は、ユーザーリストから削除する
			users = append(users[:i], users[i+1:]...)

			// 削除が成功したことを示すためにNoContentステータスを返す
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	// ユーザーが見つからなかった場合は、NotFoundステータスを返す
	w.WriteHeader(http.StatusNotFound)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// レスポンスのContent-Typeヘッダを設定する
	w.Header().Set("Content-Type", "application/json")

	// ユーザーリストをJSON形式でレスポンスとしてエンコードし、レスポンスに書き込む
	json.NewEncoder(w).Encode(users)
}

func main() {
	// ルーターを作成します
	r := mux.NewRouter()

	// ルーターに各エンドポイントとハンドラ関数を登録します
	r.HandleFunc("/users", CreateUser).Methods("POST")    // POST /users で CreateUser ハンドラを呼び出す
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")    // GET /users/{id} で GetUser ハンドラを呼び出す
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT") // PUT /users/{id} で UpdateUser ハンドラを呼び出す
	r.HandleFunc("/users", DeleteUser).Methods("DELETE")   // DELETE /users で DeleteUser ハンドラを呼び出す
	r.HandleFunc("/users", GetUsers).Methods("GET")        // GET /users で GetUsers ハンドラを呼び出す

	// サーバーを起動します
	fmt.Println("サーバーをポート8080で起動します...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

