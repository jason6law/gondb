package model

// swagger:parameters User
type User struct {
	// in: body
	Id int `json:ID`
	Name string `json:Name`
	LoginTime int `json:LoginTime`
}

func Add(user User) error {
	return nil
}

func Delete(id int) error {
	return nil
}

func Update(user User) error {
	return nil
}

func GetOne(id int) (User,error) {
	return User{},nil
}

func GetAll() ([]User,error) {
	return []User{},nil
}

func GetIdByName(name string) (int,error) {
	return 0,nil
}