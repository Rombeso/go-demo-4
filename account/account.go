package account

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"net/url"
	"reflect"
	"time"
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) Output() {
	color.Cyan(acc.Login)
	color.Cyan(acc.Password)
	color.Cyan(acc.Url)
}

func (acc *Account) generatePassword(n int) {
	arr := make([]rune, n)
	for i := range arr {
		arr[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	acc.Password = string(arr)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:     login,
		Password:  password,
		Url:       urlString,
	}
	field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("Login")
	fmt.Println(string(field.Tag))
	fmt.Println(field)
	if newAcc.Password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz123456789-*!")
