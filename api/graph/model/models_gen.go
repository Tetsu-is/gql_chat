// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Message struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Body     string `json:"body"`
}

type Mutation struct {
}

type NewMessage struct {
	UserID string `json:"user_id"`
	Body   string `json:"body"`
}

type NewUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Query struct {
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
