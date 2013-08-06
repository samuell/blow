type Printer struct {
	flow.Component
	Line <-chan []byte // Input
}

// Prints a line when it gets it
func (p *Printer) OnLine(line []byte) {
	fmt.Println(string(line))
}