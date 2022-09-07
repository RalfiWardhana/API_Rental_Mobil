package domain

type Transaction struct {
	Id          int
	Id_car      int
	Id_user     int
	Total_price int
	Username    string
	Email       string
	Duration    string
	Status      int
	Car_name    string
	Cc          int
	Capacity    int
	Total       int
	Car_type    string
}

type Transaction_get struct {
	Id          int
	Total_price int
	Username    string
	Email       string
	Duration    string
	Status      int
	Car_name    string
	Cc          int
}
