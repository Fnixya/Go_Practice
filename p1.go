package main

func main() {
	ode1 := IVPproblem{1.0, 0.0, 1.0, 0.25, func(u, t float64) float64 {
		return u
	}}
	ode1.plot("Plot of u'(t) = u", "t", "u(t)")
}
