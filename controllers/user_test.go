package controllers_test

import (
	"github.com/stretchr/testify/assert"
	"go-testing-11/controllers"
	"reflect"
	"testing"
)

func TestPerson_ChangeName(t *testing.T) {
	testCase := []struct{
		testName		string
		newName			string
		expectedName	string
		expectedError	error
	}{
		{
			testName: "change name to Ali",
			newName: "Ali",
			expectedName: "Ali",
			expectedError: nil,
		},
		{
			testName: "change name with empty string",
			newName: "",
			expectedName: "Melati",
			expectedError: controllers.ErrorPersonNameEmpty,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.testName, func(t *testing.T) {
			p := &controllers.Person{
				Name:    "Melati",
			}

			err := p.ChangeName(tc.newName)

			assert.Equal(t, tc.expectedError, err, "error not match")
			assert.Equal(t, tc.expectedName, p.Name, "name not match")

		})
	}
}

func TestPerson_AddNewFriend(t *testing.T) {
	testCase := []struct{
		testName		string
		newFriends		[]controllers.Person
	}{
		{
			testName: "Add new friend",
			newFriends: []controllers.Person{
				{
					Name: "B",
				},
				{
					Name: "C",
				},
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.testName, func(t *testing.T) {
			p := &controllers.Person{
				Name:    "A",
			}

			for _, friend := range tc.newFriends {
				p.AddNewFriend(friend)
			}

			if !reflect.DeepEqual(p.Friends, tc.newFriends) {
				t.Errorf("Add friend fail")
			}
		})
	}
}
