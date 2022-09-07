package entity

type User struct {
	Id           int
	Username     string
	Email        string
	Password     string
	Id_user_type int
	User_type    string
}

type User_type struct {
	Id        int
	User_type string
}
