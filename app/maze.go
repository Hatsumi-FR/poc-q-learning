package app

import (
	"fyne.io/fyne/v2"
)

type Maze struct {
	Row        int
	Column     int
	ColumnSize int
	RowSize    int
	App        fyne.App
}

func (m Maze) Draw() {

}
