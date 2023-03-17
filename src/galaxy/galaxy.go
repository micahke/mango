package galaxy

import (
	"github.com/aquilax/go-perlin"
)

type Galaxy struct {
	perlinGenerator *perlin.Perlin
}

var galaxy *Galaxy

var (
	GALAXY_ALPHA float32 = 2
	GALAXY_BETA  float32 = 2
	GALAXY_N     int32   = 3
	GALAXY_SEED  int32   = 100
	// GALAXY_ALPHA float32 = 0.166
	// GALAXY_BETA  float32 = 1.575
	// GALAXY_N     int32   = 3
	// GALAXY_SEED  int32   = 100
	GALAXY_FREQ float32 = 13.6 // GALAXY_FREQ float32 = 10
)

func Init() {

	galaxy = new(Galaxy)

	Rebuild()
}

func PerlinValueAtCoords(x, y int64, normalized bool) float64 {

	value := galaxy.perlinGenerator.Noise2D(float64(x)/float64(GALAXY_FREQ), float64(y)/float64(GALAXY_FREQ))

	if !normalized {
		return value
	}

	value = value / (float64(GALAXY_BETA) / float64(GALAXY_ALPHA))
	value = (value + 1) / 2
	return value

}

func Rebuild() {
	galaxy.perlinGenerator = perlin.NewPerlin(float64(GALAXY_ALPHA), float64(GALAXY_BETA), GALAXY_N, int64(GALAXY_SEED))
}
