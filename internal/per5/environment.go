package per5

func (p *Per5) Width() float64 {
	return p.width
}

func (p *Per5) Height() float64 {
	return p.height
}

// FrameRate sets the frame rate. Default is 60 frames per second.
// FrameRate must be set int Setup().
func (p *Per5) FrameRate(f int) {
	p.frameRate = float64(f)
}

// FrameCount return the number of frame updates that has been made.
// FrameCount is 0 in Setup(), 1 the first time Draw() is called, etc.
func (p *Per5) FrameCount() int {
	return p.frameCount
}
