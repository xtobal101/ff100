package main

import (
	"fmt"	
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
	
func main() {
		fmt.Println("hello world")

		a := app.New()
		w := a.NewWindow("Hello World")
	
		w.SetContent(widget.NewLabel("Hello World!"))
		w.ShowAndRun()
}


