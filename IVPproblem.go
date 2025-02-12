package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// IVPproblem is a struct that represents an initial value problem with function u' = f(u, t) and initial value u0 at time t0.

func (ivp IVPproblem) midpointRK2() []float64 {
	N := int((ivp.t1-ivp.t0)/ivp.k) + 1
	y := make([]float64, N)

	y[0] = ivp.u0
	for i := 1; i < N; i++ {
		yaux := y[i-1] + 0.5*ivp.k*ivp.f(y[i-1], ivp.t0+float64(i-1)*ivp.k)
		y[i] = y[i-1] + ivp.k*ivp.f(yaux, ivp.t0+(float64(i)-0.5)*ivp.k)
	}

	return y
}

func (ivp IVPproblem) trapezoidalRK2() []float64 {
	N := int((ivp.t1-ivp.t0)/ivp.k) + 1
	y := make([]float64, N)

	y[0] = ivp.u0
	for i := 1; i < N; i++ {
		yaux := y[i-1] + ivp.k*ivp.f(y[i-1], ivp.t0+float64(i-1)*ivp.k)
		y[i] = y[i-1] + 0.5*ivp.k*(ivp.f(y[i-1], ivp.t0+float64(i-1)*ivp.k)+ivp.f(yaux, ivp.t0+float64(i)*ivp.k))
	}

	return y
}

func (ivp IVPproblem) plotEuler() (*plotter.Line, error) {
	points := make(plotter.XYs, int((ivp.t1-ivp.t0)/ivp.k)+1)
	vals := ivp.forwardEuler()
	for i := 0; i < len(points); i++ {
		points[i].X = ivp.t0 + float64(i)*ivp.k
		points[i].Y = vals[i]
	}

	// fmt.Print("points: ", points)

	line, err := plotter.NewLine(points)
	return line, err
}

func (ivp IVPproblem) plotMidpoint() (*plotter.Line, error) {
	points := make(plotter.XYs, int((ivp.t1-ivp.t0)/ivp.k)+1)
	vals := ivp.midpointRK2()
	for i := 0; i < len(points); i++ {
		points[i].X = ivp.t0 + float64(i)*ivp.k
		points[i].Y = vals[i]
	}

	// fmt.Print("points: ", points)

	line, err := plotter.NewLine(points)
	return line, err
}

func (ivp IVPproblem) plotTrapezoidal() (*plotter.Line, error) {
	points := make(plotter.XYs, int((ivp.t1-ivp.t0)/ivp.k)+1)
	vals := ivp.trapezoidalRK2()
	for i := 0; i < len(points); i++ {
		points[i].X = ivp.t0 + float64(i)*ivp.k
		points[i].Y = vals[i]
	}

	// fmt.Print("points: ", points)

	line, err := plotter.NewLine(points)
	return line, err
}

func (ivp IVPproblem) plot(title, xlabel, ylabel string) {
	p := plot.New()

	p.Title.Text = title
	p.X.Label.Text = xlabel
	p.Y.Label.Text = ylabel

	var line *plotter.Line
	var err error

	// Add the data points to the plot
	line, err = ivp.plotEuler()
	if err != nil {
		panic(err)
	}
	line.Color = plotutil.Color(0)
	p.Add(line)

	line, err = ivp.plotMidpoint()
	if err != nil {
		panic(err)
	}
	line.Color = plotutil.Color(1)
	p.Add(line)

	line, err = ivp.plotTrapezoidal()
	if err != nil {
		panic(err)
	}
	line.Color = plotutil.Color(2)
	p.Add(line)

	// Save the plot to a PNG file
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "line_plot.png"); err != nil {
		panic(err)
	}
}
