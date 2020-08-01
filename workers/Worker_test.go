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
}
