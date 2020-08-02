package routes

import "encoding/json"

// KenobiDist asdasd
type KenobiDist struct {
	Kenobi float64 `json:"kenobi"`
}

// KenobiMsg asdasdasd
type KenobiMsg struct {
	Kenobi []string `json:"kenobi"`
}

// KenobiTopSecMsg asdasd
type KenobiTopSecMsg struct {
	Distance KenobiDist `json:"distance"`
	Message  KenobiMsg  `json:"message"`
}

// SatoDist asdasd
type SatoDist struct {
	Sato float64 `json:"sato"`
}

// SatoMsg asdasdasd
type SatoMsg struct {
	Sato []string `json:"sato"`
}

// SatoTopSecMSg asdasd
type SatoTopSecMSg struct {
	Distance SatoDist `json:"distance"`
	Message  SatoMsg  `json:"message"`
}

// SkyWalkerDist asdasd
type SkyWalkerDist struct {
	Skywalker float64 `json:"skywalker"`
}

// SkywalkerMsg asdasdasd
type SkywalkerMsg struct {
	Skywalker []string `json:"skywalker"`
}

// SkywalkerTopSecMsg asdasd
type SkywalkerTopSecMsg struct {
	Distance SkyWalkerDist `json:"distance"`
	Message  SkywalkerMsg  `json:"message"`
}

func getType(myMap map[string]map[string]json.RawMessage, key1 string) bool {

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
