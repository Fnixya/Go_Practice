package main

import {
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
}

type Method interface {
	Compute() []float64
	Plot() (*plotter.Line, error)
}

func (e Method) Plot() (*plotter.Line, error) {
	points := make(plotter.XYs, int((e.t1-e.t0)/e.k)+1)
	vals := e.Compute()
	for i := 0; i < len(points); i++ {
		points[i].X = e.t0 + float64(i)*e.k
		points[i].Y = vals[i]
	}
	return plotter.NewLine(points), nil
}

func (e EulerMethod) Compute() []float64 {
	return e.forwardEuler()
}


