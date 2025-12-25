package model

import "errors"

var ErrDuplicate = errors.New("entity already exists")

type User struct {
	ID       string
	FullName string
}
