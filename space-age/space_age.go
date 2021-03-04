package space

type Planet string

const secondsToYear = 31557600

//Planets conversion rates
var Planets = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

//Age function to calculate age on different planet
func Age(second float64, planet Planet) float64 {
	return second / Planets[planet]

}
