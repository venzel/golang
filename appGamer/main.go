package main

import (
	aGamer "appGamer/gamer"
	"fmt"
	"time"
)

func main() {
	address1Tiago := aGamer.Address{
		Street: "13 de maio",
		Number: 13,
	}

	address2Tiago := aGamer.Address{
		Street: "14 de maio",
		Number: 14,
	}

	listAddressTiago := make([]aGamer.Address, 0)
	listAddressTiago = append(listAddressTiago, address1Tiago)
	listAddressTiago = append(listAddressTiago, address2Tiago)

	address1Cintia := aGamer.Address{
		Street: "15 de maio",
		Number: 15,
	}

	address2Cintia := aGamer.Address{
		Street: "16 de maio",
		Number: 16,
	}

	listAddressCintia := make([]aGamer.Address, 0)
	listAddressCintia = append(listAddressCintia, address1Cintia)
	listAddressCintia = append(listAddressCintia, address2Cintia)

	tiago := aGamer.Account{
		Name:     "Tiago",
		Address:  listAddressTiago,
		Birthday: time.Date(1983, 12, 28, 0, 0, 0, 0, time.Local),
	}

	cintia := aGamer.Account{
		Name:     "Cintia",
		Address:  listAddressCintia,
		Birthday: time.Date(1983, 12, 28, 0, 0, 0, 0, time.Local),
	}

	listAccounts := make([]aGamer.Account, 0)
	listAccounts = append(listAccounts, tiago)
	listAccounts = append(listAccounts, cintia)

	for i := range listAccounts {
		fmt.Println(listAccounts[i].Name)

		adresses := listAccounts[i].Address

		for k := range adresses {
			fmt.Println(adresses[k].Street)
		}
	}

}
