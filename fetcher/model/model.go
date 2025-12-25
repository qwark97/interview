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

type IterData[T any] struct {
	Error error
	Data  []T
}
