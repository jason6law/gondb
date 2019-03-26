package model

import (
	"testing"
	"reflect"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		user User
		expected error
	} {
		{User{1,"test_user",1},nil},
		{User{2,"test_user2",1},nil},
		{User{3,"test_user3",1},nil},
	}

	for _,tt := range tests {
		err := Add(tt.user)
		if err != tt.expected {
			t.Errorf("add error, get %v expected %v",err,tt.expected)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		id int
		expected error
	} {
		{1,nil},
	}

	for _,tt := range tests {
		err := Delete(tt.id)
		if err != tt.expected {
			t.Errorf("delete error, get %v expected %v",err,tt.expected)
		}
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		user User
		expected error
	} {
		{User{2,"test_user2",2},nil},
	}

	for _,tt := range tests {
		err := Update(tt.user)
		if err != tt.expected {
			t.Errorf("update error, get %v expected %v",err,tt.expected)
		}
	}
}

func TestGetOne(t *testing.T) {
	tests := []struct{
		id int
		expected_user User
		expected_err error
	} {
		{2,User{2,"test_user2",2},nil},
	}

	for _,tt := range tests {
		user,err := GetOne(tt.id)
		if err != tt.expected_err {
			t.Errorf("Getone error, get %v expected %v",err,tt.expected_err)
		}
		if !reflect.DeepEqual(user,tt.expected_user) {
			t.Errorf("Getone error, get %v expected %v",user,tt.expected_user)
		}
	}
}


func TestGetAll(t *testing.T) {
	tests := []struct{
		expected_users []User
		expected_err error
	} {
		{[]User{
			User{3,"test_user3",1},
			User{2,"test_user2",2},
		},nil},
	}

	for _,tt := range tests {
		users,err := GetAll()
		if err != tt.expected_err {
			t.Errorf("GetAll error, get %v expected %v",err,tt.expected_err)
		}
		if !reflect.DeepEqual(users,tt.expected_users) {
			t.Errorf("GetAll error, get %v expected %v",users,tt.expected_users)
		}
	}
}

func TestGetIdByName(t *testing.T) {
	tests := []struct{
		name string
		expected_user User
		expected_err error
	} {
		{"test_user2",User{2,"test_user2",2},nil},
		{"test_user3",User{3,"test_user3",1},nil},
	}
	for _,tt := range tests {
		user,err := GetOneByName(tt.name)
		if err != tt.expected_err {
			t.Errorf("GetIdByName error, get %v expected %v",err,tt.expected_err)
		}
		if !reflect.DeepEqual(user,tt.expected_user) {
			t.Errorf("GetIdByName error, get %v expected %v",user,tt.expected_user)
		}
	}
}