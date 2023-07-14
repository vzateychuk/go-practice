package main

import "testing"

func TestDamage(t *testing.T) {
	sword := Sword{}
	expectedDamage := 2

	gotDamage := sword.Damage()

	if gotDamage != expectedDamage {
		t.Errorf("Damage expected: %d, got: %d", expectedDamage, gotDamage)
	}
}
