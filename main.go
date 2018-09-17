package main

import (
	"github.com/tfriedel6/canvas/sdlcanvas"
)

type point struct {
	x, y float64
}

func main() {
	wnd, cv, err := sdlcanvas.CreateWindow(1280, 720, "lines")
	if err != nil {
		println(err.Error())
		return
	}
	defer wnd.Destroy()

	pl := genPointList(100, 100, 300, 300)

	ps5 := genPointSeq(pl, []int{2, 3, 4})
	ps6 := genPointSeq(pl, []int{2, 1, 8})
	ps7 := genPointSeq(pl, []int{6, 5, 4})
	ps8 := genPointSeq(pl, []int{6, 7, 8})

	incr := 8
	segments := [][]point{ps5, ps6, ps7, ps8}
	var vertices [][]point
	for _, s := range segments {
		v := calcWaypoints(s, incr)
		vertices = append(vertices, v)
	}

	lenvert := len(vertices[0]) - 2
	framesCounter := 0
	t := 0
	wnd.MainLoop(func() {
		framesCounter++
		if framesCounter%2 == 0 {
			framesCounter = 0
	
			if t < lenvert {
				t++
			} else {
				w, h := float64(cv.Width()), float64(cv.Height())
				cv.SetFillStyle("#000")
				cv.FillRect(0, 0, w, h)
				t = 0
			}

		}

		for _, v := range vertices {
			cv.SetStrokeStyle("#1e90ff")
			cv.SetLineWidth(4.0)
			cv.BeginPath()
			cv.MoveTo(v[t].x, v[t].y)
			cv.LineTo(v[t+1].x, v[t+1].y)
			cv.Stroke()
		}
	})
}

func calcWaypoints(vertices []point, incr int) []point {
	waypoints := []point{}
	for i := 0; i < len(vertices)-1; i++ {
		pt0 := vertices[i]
		pt1 := vertices[i+1]
		switch {
		case pt1.x > pt0.x: /// px
			dx := pt1.x - pt0.x
			y := pt0.y
			for j := 0; j < incr; j++ {
				x := pt0.x + dx*float64(j)/float64(incr)
				waypoints = append(waypoints, point{x: x, y: y})
			}
		case pt1.x < pt0.x: /// nx
			dx := pt0.x - pt1.x
			y := pt0.y
			for j := 0; j < incr; j++ {
				x := pt0.x - dx*float64(j)/float64(incr)
				waypoints = append(waypoints, point{x: x, y: y})
			}
		case pt1.y > pt0.y: /// py
			dy := pt1.y - pt0.y
			x := pt0.x
			for j := 0; j < incr; j++ {
				y := pt0.y + dy*float64(j)/float64(incr)
				waypoints = append(waypoints, point{x: x, y: y})
			}
		case pt1.y < pt0.y: /// ny
			dy := pt0.y - pt1.y
			x := pt0.x
			for j := 0; j < incr; j++ {
				y := pt0.y - dy*float64(j)/float64(incr)
				waypoints = append(waypoints, point{x: x, y: y})
			}
		}
	}
	return waypoints
}

func genPointList(x, y, dx, dy float64) []point {
	pl := make([]point, 9)
	pl[1] = point{x, y}
	pl[2] = point{(dx-x)/2 + x, y}
	pl[3] = point{dx, y}
	pl[4] = point{dx, (dy-y)/2 + y}
	pl[5] = point{dx, dy}
	pl[6] = point{(dx-x)/2 + x, dy}
	pl[7] = point{x, dy}
	pl[8] = point{x, (dy-y)/2 + y}
	return pl
}

func genPointSeq(pl []point, seq []int) []point {
	var pseq []point
	for _, v := range seq {
		pseq = append(pseq, pl[v])
	}
	return pseq
}
