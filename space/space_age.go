//Package space implements a library containing interplanetary calculations
package space

type Planet string

const earthSeconds = 31557600 // 365 earth days in seconds

var orbitalPeriods = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

//Age returns someone's age on the given planet
func Age(seconds float64, planet Planet) float64 {

	spaceAge := seconds / earthSeconds / orbitalPeriods[planet]
	return spaceAge
}
