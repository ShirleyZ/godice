package dice

import "testing"

func TestRoll(t *testing.T) {
	_, _ = Roll("~!roll 1d20 to check for rad loot")
	_, _ = Roll("~!roll 5d10 to check for rad loot")
	_, _ = Roll("~!roll 7d8 to check for rad loot")
	_, _ = Roll("~!roll 1d100 to check for rad loot")
	_, _ = Roll("~!roll 3d4 to check for rad loot")
}
