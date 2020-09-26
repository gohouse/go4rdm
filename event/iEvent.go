package event

type IEvent interface {
	//Build() fyne.CanvasObject
	//Reset()
	//ReBuild()
	//Clear()
	Notify(arg EventObject)
}
