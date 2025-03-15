package space

type Planet string

var rates = map[string]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	if rate, ok := rates[string(planet)]; ok {
		return seconds / (60 * 60 * 24 * 365.25 * rate)
	}

	return -1
}
