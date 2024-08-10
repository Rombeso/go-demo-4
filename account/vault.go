package account

import (
	"demo/password/files"
	"encoding/json"
	"github.com/fatih/color"
	"strings"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) FindAccountsByUrl(url string) []Account {
	var arr []Account
	for _, acc := range vault.Accounts {
		isMatched := strings.Contains(acc.Url, url)
		if isMatched {
			arr = append(arr, acc)
		}
	}
	return arr
}

func (vault *Vault) DeleteAccountsByUrl(url string) bool {
	var arr []Account
	isDeleted := false
	for _, acc := range vault.Accounts {
		isMatched := strings.Contains(acc.Url, url)
		if !isMatched {
			arr = append(arr, acc)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = arr
	vault.save()
	return isDeleted
}

func (vault *Vault) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать файл data.json")
	}
	files.WriteFile(data, "data.json")
}
