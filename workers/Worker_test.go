package workers

import (
	"GenovaChallenge/models"
	"GenovaChallenge/satelites"
	"testing"
)

func TestGetLocation(t *testing.T) {
	sato := models.Sato{X: satelites.SatoPosX, Y: satelites.SatoPosY, Z: satelites.SatoPosZ}
	kenobi := models.Kenobi{X: satelites.KenobiPosX, Y: satelites.KenobiPosY, Z: satelites.KenobiPosZ}
	skywalker := models.Skywalker{X: satelites.SkywalkerPosX, Y: satelites.SkywalkerPosY, Z: satelites.SkywalkerPosZ}

	w := Worker{Kenobi: kenobi, Sato: sato, Skywalker: skywalker}
	_, _, _, err1 := w.GetLocation(5.2, 8.6, 12.5)

	if err1 != nil {
		t.Errorf("Error found trying to get position: %v \n", err1)
	}
	_, _, _, err2 := w.GetLocation(-1, -1, -1)

	if err2 == nil {
		t.Errorf("A calculation error should have happened\n")
	}
}

func TestGetMessage(t *testing.T) {
	sato := models.Sato{X: satelites.SatoPosX, Y: satelites.SatoPosY, Z: satelites.SatoPosZ}
	kenobi := models.Kenobi{X: satelites.KenobiPosX, Y: satelites.KenobiPosY, Z: satelites.KenobiPosZ}
	skywalker := models.Skywalker{X: satelites.SkywalkerPosX, Y: satelites.SkywalkerPosY, Z: satelites.SkywalkerPosZ}

	w := Worker{Kenobi: kenobi, Sato: sato, Skywalker: skywalker}
	_, err2 := w.GetMessage([]string{"This", "", "", "message"}, []string{"", "a", "new", "message"}, []string{"This", "", "", ""})

	if err2 != nil {
		t.Errorf("Error found trying to get message: %v \n", err2)
	}

	_, err3 := w.GetMessage([]string{"", "", "", "message"}, []string{"", "a", "new", "message"}, []string{"", "", "", ""})

	if err3 == nil {
		t.Errorf("A validation error should have happened (An empty word)\n")
	}

	_, err4 := w.GetMessage([]string{"This", "", "", "message"}, []string{"", "a", "new", "message"}, []string{"Not this", "", "", ""})

	if err4 == nil {
		t.Errorf("A validation error should have happened (Mismatch word between satelite 1 and 3)\n")
	}

	_, err5 := w.GetMessage([]string{"This", "", "", "message"}, []string{"", "a", "new", "NOT message"}, []string{"", "", "", ""})

	if err5 == nil {
		t.Errorf("A validation error should have happened (Mismatch word between satelite 1 and 2)\n")
	}

	_, err6 := w.GetMessage([]string{"This", "", "", "message"}, []string{"", "a", "new", "message"}, []string{"", "", "NOT new", ""})

	if err6 == nil {
		t.Errorf("A validation error should have happened (Mismatch word between satelite 2 and 3)\n")
	}

	_, err7 := w.GetMessage([]string{"", "", "", "message"}, []string{"", "a", "new", "message"}, []string{"This", "", ""})

	if err7 == nil {
		t.Errorf("A validation error should have happened (Mismatch msgs length between satelites, sat 3)\n")
	}

	_, err8 := w.GetMessage([]string{"", "", "message"}, []string{"", "a", "new", "message"}, []string{"This", "", "", ""})

	if err8 == nil {
		t.Errorf("A validation error should have happened (Mismatch msgs length between satelites, sat 1)\n")
	}

	_, err9 := w.GetMessage([]string{"", "", "", "message"}, []string{"", "a", "new"}, []string{"This", "", "", ""})

	if err9 == nil {
		t.Errorf("A validation error should have happened (Mismatch msgs length between satelites, sat 2)\n")
	}
}
