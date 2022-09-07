package entity

type Car struct {
	Id          int
	Car_name    string
	Cc          int
	Capacity    int
	Total       int
	Id_car_type int
	Car_type    string
	Price       int
}

type Car_type struct {
	Id       int
	Car_type string
}
