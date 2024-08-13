package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"strings"
)

var menu = map[string]func(db *account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	//files.Read()
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось найти ENV файл")
	}
	//res := os.Getenv("VAR")
	// fmt.Println(res)
	vault := account.NewVault(files.NewJsonDb("data.vault"), *encrypter.NewEncrypter())
	// vault := account.NewVault(cloud.NewCloudDb("www.data.js"))

	//for _, e := range os.Environ() {
	//	//fmt.Println(e)
	//	pair := strings.SplitN(e, "=", 2)
	//	fmt.Println(pair[0])
	//}
Menu:
	for {
		input := promptData(
			"1. Создать аккаунт",
			"2. Найти аккаунт по URL",
			"3. Найти аккаунт по логину",
			"4. Удалить аккаунт",
			"5. Выход",
			"Выберите вариант меню")
		menuFunc := menu[input]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
		//switch input {
		//case "1":
		//	createAccount(vault)
		//case "2":
		//	findAccount(vault)
		//case "3":
		//	deleteAccount(vault)
		//case "4":
		//	break Menu
		//default:
		//	println("Не выбран корректный пункт меню")
		//}
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Не верный формат URL или Login")
		return
	}

	vault.AddAccount(*myAccount)
}

func promptData(prompt ...any) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}

	var res string
	fmt.Scan(&res)
	return res
}

func findAccountByUrl(vault *account.VaultWithDb) {
	var str string
	var res []account.Account
	fmt.Print("Введите url для поиска:")
	fmt.Scan(&str)
	res = vault.FindAccounts(str, checkUrl)
	outputResults(&res)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	var str string
	var res []account.Account
	fmt.Print("Введите login для поиска:")
	fmt.Scan(&str)
	res = vault.FindAccounts(str, checkLogin)
	outputResults(&res)
}

func outputResults(res *[]account.Account) {
	if len(*res) > 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, acc := range *res {
		acc.Output()
	}
	fmt.Println(res)
}

func checkUrl(acc account.Account, str string) bool {
	return strings.Contains(acc.Url, str)
}
func checkLogin(acc account.Account, str string) bool {
	return strings.Contains(acc.Login, str)
}

func deleteAccount(vault *account.VaultWithDb) {
	var url string
	fmt.Print("Введите url для поиска:")
	fmt.Scan(&url)
	isDeleted := vault.DeleteAccountsByUrl(url)

	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}
