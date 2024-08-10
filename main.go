package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	//files.ReadFile()
	vault := account.NewVault()
Menu:
	for {
		input := getMenu()
		switch input {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		case 4:
			break Menu
		default:
			println("Не выбран корректный пункт меню")
		}
	}
}

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Не верный формат URL или Login")
		return
	}

	vault.AddAccount(*myAccount)
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON ")
		return
	}
	files.WriteFile(data, "data.json")
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}

func getMenu() int {
	var newItem int
	fmt.Print(`Выберите вариант меню:
	1. Создать аккаунт
	2. Найти аккаунт
	3. Удалить аккаунт
	4. Выход
	`)
	fmt.Scan(&newItem)
	return newItem

}

func findAccount(vault *account.Vault) {
	var url string
	fmt.Print("Введите url для поиска:")
	fmt.Scan(&url)
	res := vault.FindAccountsByUrl(url)
	if len(res) > 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, acc := range res {
		acc.Output()
	}
	fmt.Println(res)
}

func deleteAccount(vault *account.Vault) {
	var url string
	fmt.Print("Введите url для поиска:")
	fmt.Scan(&url)
	isDeleted := vault.DeleteAccountsByUrl(url)

	if isDeleted {
		color.Green("Удалено")
	}
}
