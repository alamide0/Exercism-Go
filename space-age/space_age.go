package space

type Planet string

func Age(seconds float64, planet Planet) float64 {

	r := seconds / 31557600
	var rate float64 = 0

	switch planet {
	case "Mercury":
		rate = 0.2408467 
	case "Venus":
		rate =  0.61519726
	case "Earth":
		rate = 1.0
	case "Mars":
		rate = 1.8808158
	case "Jupiter":
		rate = 11.862615
	case "Saturn":
		rate = 29.447498
	case "Uranus":
		rate = 84.016846
	case "Neptune":	
		rate = 164.79132
	}
	return r / rate
}

