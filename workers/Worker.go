package workers

import (
	"ML/models"
	"fmt"
	"strings"

	"github.com/co60ca/trilateration"
)

// Worker asdasd
type Worker struct {
	Kenobi    models.Kenobi
	Skywalker models.Skywalker
	Sato      models.Sato
}

// GetLocation sads
// input: distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func (w Worker) GetLocation(distKenobi, distSkywalker, distSato float64) (x, y, z float64, err error) {
	kenobiPos := trilateration.Point3{w.Kenobi.X, w.Kenobi.Y, w.Kenobi.Z}
	skywalkerPos := trilateration.Point3{w.Skywalker.X, w.Skywalker.Y, w.Skywalker.Z}
	satoPos := trilateration.Point3{w.Sato.X, w.Sato.Y, w.Sato.Z}

	positions := []trilateration.Point3{kenobiPos, skywalkerPos, satoPos}
	distances := []float64{distKenobi, distSkywalker, distSato}

	params := trilateration.Parameters3{Loc: positions, Dis: distances}
	loc, err := params.SolveTrilat3()

	if err != nil {
		return 0.0, 0.0, 0.0, err
	}

	return loc[0], loc[1], loc[2], nil

}

// GetMessage sads
/**
* Input: El mensaje recibido en cada satelite tendra la misma cantidad de slots (palabras). Si asi no sucede, el mensaje sera rechazado de plano
* Output: El mensaje tal cual lo genera el emisor del mensaje
 */
func (w Worker) GetMessage(msgKenobi, msgSkywalker, msgSato []string) (message string, err error) {
	result := ""

	if len(msgKenobi) == len(msgSkywalker) && len(msgSkywalker) == len(msgSato) {
		for i, val := range msgKenobi {
			if val == "" && msgSkywalker[i] == "" && msgSato[i] == "" {
				return result, fmt.Errorf("Some word was not received in any satelite")
			}

			if len(val) > 0 {
				result += val + " "
			} else if len(msgSkywalker[i]) > 0 {
				result += msgSkywalker[i] + " "
			} else {
				result += msgSato[i] + " "
			}

		}

		return strings.Trim(result, " "), nil
	}

	return result, fmt.Errorf("The message lengths are not equal")

}
