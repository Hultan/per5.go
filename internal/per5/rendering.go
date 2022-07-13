package per5

func (p *Per5) CreateCanvas(width, height int) {
	p.da.SetSizeRequest(width, height)
	p.width, p.height = float64(width), float64(height)
}
