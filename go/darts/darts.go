package darts

func Score(x, y float64) int {
	switch {
	case isOnCircle(x, y, 1):
		return 10
	case isOnCircle(x,y, 5):
		return 5
	case isOnCircle(x, y, 10): return 1
	default: return 0
	}
}

func isOnCircle(x, y, radius float64) bool {
	return x*x+y*y <= radius*radius
}
