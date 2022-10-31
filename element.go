package gui

// Element represents all drawable objects. Other libraries call these "widgets"
type Element interface {
	Draw()
	Trigger()
}
