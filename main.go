package main

import "fmt"

func main() {
	esword := EnchantedSword{
		Sword: Sword{Name: "EnchantedSword"},
	}
	fmt.Printf("Name: %s, Damage: %d\n", esword.Name, esword.Damage())
	fmt.Println(esword.String())
}
