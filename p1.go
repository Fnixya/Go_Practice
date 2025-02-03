package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func forwardEuler(f func(float64, float64) float64, u0, t0, t1, k float64) []float64 {
	N := int((t1-t0)/k) + 1
	t1 = t0 + float64(N)*k

	y := make([]float64, N)
	y[0] = u0
	for i := 1; i < N; i++ {
		y[i] = y[i-1] + k*f(t0+float64(i-1)*k, y[i-1])
	}

	return y
}

func ivp(t, u float64) float64 {
	return u
}

func main() {
	// Create a new plot
	p := plot.New()

	p.Title.Text = "Simple Line Plot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Create some random data points
	t0, t1, k := 0.0, 10.0, 0.1
	points := make(plotter.XYs, int((t1-t0)/k)+1)
	vals := forwardEuler(ivp, 5, t0, t1, k)
	for i := 0; i < len(points); i++ {
		points[i].X = t0 + float64(i)*k
		points[i].Y = vals[i]
	}

	// Add the data points to the plot
	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}
	line.Color = plotutil.Color(0)
	p.Add(line)

	// Save the plot to a PNG file
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "line_plot.png"); err != nil {
		panic(err)
	}
}
