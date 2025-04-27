package models

type User struct {
    ID       string `json:"id" bson:"_id,omitempty"`
    Name     string `json:"name" bson:"name"`
    Email    string `json:"email" bson:"email"`
    Phone    string `json:"phone" bson:"phone"`
    Password string `json:"password" bson:"password"`
}
