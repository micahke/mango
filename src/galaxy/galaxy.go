package galaxy

import (

	"github.com/aquilax/go-perlin"
)


type Galaxy struct {

  perlinGenerator *perlin.Perlin


}


var galaxy *Galaxy


var (
  GALAXY_ALPHA float32 = 3
  GALAXY_BETA float32 = 2.5
  GALAXY_N int32 = 3
  GALAXY_SEED int32 = 100
)


func Init() {

  galaxy = new(Galaxy) 

  Rebuild()
}


func PerlinValueAtCoords(x, y int64, normalized bool) float64 {


  value := galaxy.perlinGenerator.Noise2D(float64(x) / 10, float64(y) / 10)

  if !normalized {
    return value
  }

  value = (value + 1) / 2
  return value

}


func Rebuild() {
  galaxy.perlinGenerator = perlin.NewPerlin(float64(GALAXY_ALPHA), float64(GALAXY_BETA), GALAXY_N, int64(GALAXY_SEED))
}
