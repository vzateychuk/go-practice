package main

type EnchantedSword struct {
	Sword // Embed the Sword, все методы Sword теперь в EnchantedSword
}

func (EnchantedSword) Damage() int {
	return 42 // Damage returns big damage dealt by the enchanted sword.
}
