package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
)

const (
	row    = 3
	column = 5
)

var (
	cellWidth        float32
	cellHeight       float32
	resolutionHeight float32
	resolutionWidth  float32
)

func createCircle(grid *fyne.Container, x, y int) error {
	circle := canvas.NewCircle(color.RGBA{
		R: 204,
		G: 33,
		B: 33,
		A: 250,
	})
	circle.Resize(fyne.NewSize(cellWidth, cellHeight))
	rect := canvas.NewRectangle(color.White)
	rect.SetMinSize(fyne.NewSize(cellWidth, cellHeight))
	cell := container.New(layout.NewMaxLayout())
	cell.Add(rect)
	cell.Add(circle)
	location := x*column + y
	grid.Objects[location] = cell
	return nil
}

func createGrid(cols, rows int) *fyne.Container {
	grid := container.New(layout.NewGridLayout(cols))

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			rect := canvas.NewRectangle(color.White)
			rect.SetMinSize(fyne.NewSize(cellWidth, cellHeight))
			rect.Move(fyne.NewPos(float32(x)*cellWidth, float32(y)*cellHeight))
			grid.Add(rect)
		}
	}

	return grid
}

func new() {
	cellWidth = resolutionWidth / float32(column)
	cellHeight = resolutionHeight / float32(row)
	resolutionWidth, resolutionHeight = 1920, 1080
}

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Quadrillage")

	grid := createGrid(column, row)
	createCircle(grid, 1, 2)
	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(resolutionWidth, resolutionHeight))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()

}
