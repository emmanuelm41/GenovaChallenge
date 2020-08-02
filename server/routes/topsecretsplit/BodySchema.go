package topsecretsplit

import "encoding/json"

// KenobiDist Distancia del emisor al satelite Kenobi
type KenobiDist struct {
	Kenobi float64 `json:"kenobi"`
}

// KenobiMsg Mensajes recibidos por el satelite Kenobi
type KenobiMsg struct {
	Kenobi []string `json:"kenobi"`
}

// KenobiTopSecMsg Distancia y mensajes recibidos por el satelite Kenobi
type KenobiTopSecMsg struct {
	Distance KenobiDist `json:"distance"`
	Message  KenobiMsg  `json:"message"`
}

// SatoDist Distancia del emisor al satelite Sato
type SatoDist struct {
	Sato float64 `json:"sato"`
}

// SatoMsg Mensajes recibidos por el satelite Sato
type SatoMsg struct {
	Sato []string `json:"sato"`
}

// SatoTopSecMSg Distancia y mensajes recibidos por el satelite Sato
type SatoTopSecMSg struct {
	Distance SatoDist `json:"distance"`
	Message  SatoMsg  `json:"message"`
}

// SkyWalkerDist Distancia del emisor al satelite SkyWalker
type SkyWalkerDist struct {
	Skywalker float64 `json:"skywalker"`
}

// SkywalkerMsg Mensajes recibidos por el satelite SkyWalker
type SkywalkerMsg struct {
	Skywalker []string `json:"skywalker"`
}

// SkywalkerTopSecMsg Distancia y mensajes recibidos por el satelite SkyWalker
type SkywalkerTopSecMsg struct {
	Distance SkyWalkerDist `json:"distance"`
	Message  SkywalkerMsg  `json:"message"`
}

// Permite identificar si un mensaje recibido pertenece a un satelite dado - Input: Mapa generado a partir del JSON recibido - Output: Indica si el mensaje recibido corresponde al satelite en cuestion
func getSateliteFromMsg(myMap map[string]map[string]json.RawMessage, key1 string) bool {

	found1 := false
	found2 := false

	for k := range myMap {
		if k == "distance" {

			for k1 := range myMap[k] {
				if k1 == key1 {
					found1 = true
				} else {
					return false
				}
			}
		} else if k == "message" {

			for k1 := range myMap[k] {
				if k1 == key1 {
					found2 = true
				} else {
					return false
				}
			}
		} else {
			return false
		}
	}

	return found1 && found2

}
