package main

//   100 Inviatados Dijerón.
//   Party trivia game based on the TV show of a similiar name.
//   Todo:
//   Add the question label centered on top.
// * Create a two by five table for the answers and corresponding percentages. Done

// * Change color to Dark Theme.
// * Resize the entire app.
// *

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {

	//App and main window definition
	myApp := app.New()
	myWindow := myApp.NewWindow("100 Invitados Dijerón")
	myWindow.Resize(fyne.NewSize(800, 600))

	// answersContainer elements definition
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	rect1 := canvas.NewRectangle(green)
	rect2 := canvas.NewRectangle(green)
	rect3 := canvas.NewRectangle(green)
	rect4 := canvas.NewRectangle(green)

	rect1.SetMinSize(fyne.NewSize(50, 50))
	rect2.SetMinSize(fyne.NewSize(50, 50))
	rect3.SetMinSize(fyne.NewSize(50, 50))
	rect4.SetMinSize(fyne.NewSize(50, 50))

	p1 := widget.NewLabel("Respuesta Uno    ...")
	p2 := widget.NewLabel("Respuesta Dos    ...")
	p3 := widget.NewLabel("Respuesta Tres   ...")
	p4 := widget.NewLabel("Respuesta Cuatro ...")
	p5 := widget.NewLabel("Respuesta Cinco  ...")

	answersContainer := container.NewMax(container.NewVBox(p1, p2, p3, p4, p5))

	answersContainer.Layout = layout.NewCenterLayout()

	//mainContainer elements definition ...

	red := color.NRGBA{R: 180, G: 0, B: 0, A: 255}
	yellow := color.NRGBA{R: 180, G: 180, B: 0, A: 255}
	blue := color.NRGBA{R: 180, G: 0, B: 180, A: 255}
	orange := color.NRGBA{R: 180, G: 180, B: 180, A: 255}

	top := canvas.NewRectangle(yellow)
	bottom := canvas.NewRectangle(orange)
	left := canvas.NewRectangle(red)
	right := canvas.NewRectangle(blue)

	top.SetMinSize(fyne.NewSize(50, 50))
	bottom.SetMinSize(fyne.NewSize(50, 50))
	left.SetMinSize(fyne.NewSize(50, 50))
	right.SetMinSize(fyne.NewSize(50, 50))

	mainContainer := container.NewBorder(top, bottom, left, right, answersContainer)

	//Add the main container to the window then show and run
	myWindow.SetContent(mainContainer)
	myWindow.ShowAndRun()

}
