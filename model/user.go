package model

import (
	"haologs.com/gondb/db"
	"fmt"
	"github.com/pingcap/goleveldb/leveldb/errors"
	"strconv"
	"strings"
)

// swagger:parameters User
type User struct {
	// in: body
	Id int `json:ID`
	Name string `json:Name`
	LoginTime int `json:LoginTime`
}

const (
	PREFIX = "::KEY::USER::"
	NAME_SUFFIX  = "::NAME::"
	LOGINTIME_SUFFIX = "::LOGINTIME::"
	INDEX_PERFIX = "::IDX::USER::"
	SEPARATOR = "::"
)

func checkUser(user User) error {
	if user.Id < 0 && user.Name == "" {
		return errors.New("parameters error")
	}
	return nil
}

func Add(user User) error {

	err := checkUser(user)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s%d%s",PREFIX,user.Id,NAME_SUFFIX)
	err = db.Client.Set(key,user.Name,0).Err()
	if err != nil {
		return err
	}

	key = fmt.Sprintf("%s%d%s",PREFIX,user.Id,LOGINTIME_SUFFIX)
	err = db.Client.Set(key,user.LoginTime,0).Err()
	if err != nil {
		return err
	}

	key = fmt.Sprintf("%s%s%s",INDEX_PERFIX,user.Name,NAME_SUFFIX)
	err = db.Client.Set(key,user.Id,0).Err()
	if err != nil {
		return err
	}

	return nil
}

func Delete(id int) error {
	key := fmt.Sprintf("%s%d%s",PREFIX,id,NAME_SUFFIX)
	name,err := db.Client.Get(key).Result()
	if err != nil {
		return err
	}
	err = db.Client.Del(key).Err()
	if err != nil {
		return err
	}
	key = fmt.Sprintf("%s%d%s",PREFIX,id,LOGINTIME_SUFFIX)
	err = db.Client.Del(key).Err()
	if err != nil {
		return err
	}
	key = fmt.Sprintf("%s%s%s",INDEX_PERFIX,name,NAME_SUFFIX)
	err = db.Client.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}

func Update(user User) error {

	return Add(user)
}

func GetOne(id int) (User,error) {

	key := fmt.Sprintf("%s%d%s",PREFIX,id,NAME_SUFFIX)
	name,err := db.Client.Get(key).Result()
	if err != nil {
		return User{},err
	}

	key = fmt.Sprintf("%s%d%s",PREFIX,id,LOGINTIME_SUFFIX)
	loginTime,err := db.Client.Get(key).Result()
	if err != nil {
		return User{},err
	}

	iloginTime,err := strconv.Atoi(loginTime)
	if err != nil {
		return User{},err
	}

	return User{id,name,iloginTime},nil
}

func GetAll() ([]User,error) {

	Users := []User{}
	cursor := uint64(0)
	match := fmt.Sprintf("%s*%s",PREFIX,NAME_SUFFIX)
	for {
		keys,cursor,err := db.Client.Scan(cursor,match,10).Result()
		if err != nil {
			return []User{},err
		}
		for _,key := range keys {
			if !strings.HasPrefix(key,PREFIX) || !strings.HasSuffix(key,NAME_SUFFIX) {
				continue
			}
			tmp := strings.Split(key,SEPARATOR)
			if len(tmp) != 6 {
				continue
			}
			id,err := strconv.Atoi(tmp[3])
			if err != nil {
				return []User{},err
			}
			user,err := GetOne(id)
			if err != nil {
				return []User{},err
			}
			Users = append(Users, user)
		}
		if cursor == 0 {
			break
		}
	}

	return Users,nil
}

func GetOneByName(name string) (User,error) {

	key := fmt.Sprintf("%s%s%s",INDEX_PERFIX,name,NAME_SUFFIX)
	result,err := db.Client.Get(key).Result()
	if err != nil {
		return User{},err
	}
	id,err := strconv.Atoi(result)
	if err != nil {
		return User{},err
	}

	return GetOne(id)
}