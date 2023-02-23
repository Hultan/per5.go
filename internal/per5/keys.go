package per5

func (p *Per5) SetKeyPressedFunc(f func(per5 *Per5)) {
	p.keyPressedFunc = f
}

func (p *Per5) KeyCode() uint {
	return p.keyCode
}
