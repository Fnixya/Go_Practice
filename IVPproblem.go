package main

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// IVPproblem is a struct that represents an initial value problem with function u' = f(u, t) and initial value u0 at time t0.
type IVPproblem struct {
	u0, t0, t1, k float64
	f             func(u, t float64) float64
}

func (ivp IVPproblem) forwardEuler() []float64 {
	N := int((ivp.t1-ivp.t0)/ivp.k) + 1
	// t1 = t0 + float64(N)*k
	y := make([]float64, N)

	y[0] = ivp.u0
	for i := 1; i < N; i++ {
		y[i] = y[i-1] + ivp.k*ivp.f(ivp.t0+float64(i-1)*ivp.k, y[i-1])
	}

	return y
}

func (ivp IVPproblem) backwardEuler() []float64 {
	return nil
}

func (ivp IVPproblem) plot(title, xlabel, ylabel string) {
	p := plot.New()

	p.Title.Text = title
	p.X.Label.Text = xlabel
	p.Y.Label.Text = ylabel

	points := make(plotter.XYs, int((ivp.t1-ivp.t0)/ivp.k)+1)
	vals := ivp.forwardEuler()
	for i := 0; i < len(points); i++ {
		points[i].X = ivp.t0 + float64(i)*ivp.k
		points[i].Y = vals[i]
	}

	// Add the data points to the plot
	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}
	line.Color = plotutil.Color(1)
	p.Add(line)

	fmt.Print("points: ", points)

	// Save the plot to a PNG file
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "line_plot.png"); err != nil {
		panic(err)
	}
}
