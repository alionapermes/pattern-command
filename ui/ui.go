package ui

import (
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	objSize  = 30
	uiHeight = 300
	uiWidth  = uiHeight * 2
)

var objColor = color.NRGBA{0, 255, 0, 0}

type UI struct {
	fyne.Window
	fyne.Container

	CurrentObj fyne.CanvasObject

	buttons []fyne.CanvasObject
}

func NewUI(app fyne.App, title string) *UI {
	return &UI{Window: app.NewWindow(title)}
}

func (self *UI) RegisterBtn(label string, onclick func()) {
	self.buttons = append(self.buttons, widget.NewButton(label, onclick))
}

func (self *UI) Setup() {
	size := fyne.NewSize(uiWidth, uiHeight)
	self.Canvas().Content().Resize(size)

	btnContainer := container.NewHBox(self.buttons...)
	content := container.NewVBox(
		btnContainer,
		widget.NewSeparator(),
		self.Canvas().Content(),
	)

	self.SetContent(content)
}

func (self *UI) Run() {
	self.Window.ShowAndRun()
}

func (self *UI) CreateObject(obj fyne.CanvasObject) {
	self.Container.Add(obj)
}

func (self *UI) CreateRandomObject() fyne.CanvasObject {
	obj := newObj()
	self.Container.Add(obj)

	return obj
}

func (self *UI) DeleteObject(obj fyne.CanvasObject) {
	self.Container.Remove(obj)
}

func (self *UI) MoveObject(obj fyne.CanvasObject, movement fyne.Vector2) {
	obj.Position().Add(movement)
}

func (self *UI) ScaleObject(obj fyne.CanvasObject, k float32) {
	oldWidth, oldHeight := obj.Size().Components()
	newWidth := oldWidth * k
	newHeight := oldHeight * k

	obj.Resize(fyne.NewSize(newWidth, newHeight))
}

func (self *UI) GetRandomObject() fyne.CanvasObject {
	id := rand.Intn(len(self.Container.Objects))
	return self.Container.Objects[id]
}

func newObj() fyne.CanvasObject {
	randX := rand.Float32() * uiWidth
	randY := rand.Float32() * uiHeight

	obj := canvas.NewCircle(objColor)
	obj.Size().AddWidthHeight(objSize, objSize)
	obj.Position().AddXY(randX, randY)

	return obj
}
