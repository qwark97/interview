package model

type Response struct {
	NextLink string
	Users    []User
}

type User struct {
	ID        string
	FirstName string
	LastName  string
}

type DataBatch struct {
	Users []User
	Error error
}
