package main

import (
	"exercism-go/src"
	"exercism-go/src/modules"
	"exercism-go/src/modules/accounts"
	"exercism-go/src/modules/mastertable"
)

func main() {
	src.Saludar("Gopher")
	modules.SaludoFormal("Gopher")
	accounts.Accounts("cuentas")
	mastertable.MasterTable("tblasglobales")
}
	