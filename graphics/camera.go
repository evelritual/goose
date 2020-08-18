package graphics

// Camera declares all methods needs to render the scene at a specific location
// and scale.
type Camera interface {
	SetPosition(x, y int32)
	SetScale(x, y float32)
	ScaleX() float32
	ScaleY() float32
	X() int32
	Y() int32
}
