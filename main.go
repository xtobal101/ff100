package main

// 16/03/23 Main goal for this file. To define a container canvas to deploy the questions and answers.

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type Question struct {
	pregunta   string //may I have default values??
	respuesta1 string
	respuesta2 string
	respuesta3 string
}

type GBL struct {
	strLabel1       binding.String
	strLabel2       binding.String
	currentQuestion int
	questionario    map[int]*Question
}

func (gbl *GBL) init() {

	gbl.currentQuestion = 0
	gbl.questionario = make(map[int]*Question)
	//Create 3 questions

	q0 := new(Question)
	q0.pregunta = "100 Invitados Dijerón"
	q0.respuesta1 = "1.-                "
	q0.respuesta2 = "2.-                "
	q0.respuesta3 = "3.-                "

	q1 := new(Question)
	q1.pregunta = "Animales con pelaje blanco"
	q1.respuesta1 = "Conejo"
	q1.respuesta2 = "Caballo"
	q1.respuesta3 = "Oso polar"

	q2 := new(Question)
	q2.pregunta = "Marcas de salchicha"
	q2.respuesta1 = "Swan"
	q2.respuesta2 = "Oscar Mayer"
	q2.respuesta3 = "Food"

	q3 := new(Question)
	q3.pregunta = "Tipo de baile"
	q3.respuesta1 = "Tango"
	q3.respuesta2 = "Salsa"
	q3.respuesta3 = "Cumbia"

	gbl.questionario[0] = q0
	gbl.questionario[1] = q1
	gbl.questionario[2] = q2
	gbl.questionario[3] = q3

	gbl.strLabel1 = binding.NewString()
	gbl.strLabel1.Set("Hola tio ...")

	gbl.strLabel2 = binding.NewString()
	gbl.strLabel2.Set("Current Question ?")

}

func (gbl *GBL) buttonUp(ev *fyne.KeyEvent) {
	fmt.Println("ButtonUp: ")
	fmt.Println(ev.Name)

	switch ev.Name {
	case "Down":

		fmt.Println("Down")
		gbl.currentQuestion -= 1
		if gbl.currentQuestion < 1 {
			gbl.currentQuestion = 1
		}

		gbl.strLabel2.Set(fmt.Sprint(gbl.currentQuestion))

	case "Up":
		fmt.Println("Up")
		gbl.currentQuestion += 1
		if gbl.currentQuestion > 3 {
			gbl.currentQuestion = 3
		}
		gbl.strLabel2.Set(fmt.Sprint(gbl.currentQuestion))

	case "Left":
		gbl.strLabel1.Set("Left")
	case "Right":
		gbl.strLabel1.Set("Right")
	case "Space":
		fmt.Println("Space")
		gbl.strLabel1.Set(gbl.questionario[gbl.currentQuestion].pregunta)
		// Place here the answers to the questions too

	case "Return":
		gbl.strLabel1.Set("Return")
	default:
		gbl.strLabel1.Set("Vacio")
	}
}
func main() {

	fmt.Println("Hola tio ...")

	myApp := app.New()
	myWindow := myApp.NewWindow("Hola tio ...")

	gbl := new(GBL) // this object stores binding.string to update labels based on keyboard events.
	gbl.init()

	myWindow.SetContent(widget.NewLabelWithData(gbl.strLabel2))

	if deskCanvas, ok := myWindow.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyUp(gbl.buttonUp)
	} // Validates the existence of a keyboard by checking if the app is running as desktop

	myWindow.ShowAndRun()

}
