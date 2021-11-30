package controllers

import "errors"

var (
	ErrorPersonNameEmpty = errors.New("person name must exist")
)

type Person struct {
	Name    string
	Friends []Person
}

func (p *Person) ChangeName(name string) error {
	if name == "" {
		return ErrorPersonNameEmpty
	}

	p.Name = name

	return nil
}

func (p *Person) AddNewFriend(friend Person) {
	p.Friends = append(p.Friends, friend)
}