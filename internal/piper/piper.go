package piper

import (
	"bufio"
	"io"
	"os"
	"tiddly-cli/internal/logger"
)

var (
	isPipeSet = false
	data      []rune
)

// Pipe is an interface for the stdin pipe
type Pipe struct {
	log logger.Logger
}

// New creates a new Pipe
func New(log logger.Logger) *Pipe {
	p := new(Pipe)
	p.log = log

	p.setup()
	return p
}

func (p *Pipe) setup() {
	f, err := os.Stdin.Stat()

	if err != nil {
		p.log.Fatal("Error reading the pipe from stdin")
	}

	if f.Mode()&os.ModeNamedPipe != 0 {
		isPipeSet = true
		p.readData()
	}
}

func (p *Pipe) readData() {
	r := bufio.NewReader(os.Stdin)

	for {
		input, _, err := r.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		data = append(data, input)
	}

	p.log.Info("Pipe read complete")
}

// IsPipeSet returns bool for checking if piping
func (p *Pipe) IsPipeSet() bool {
	return isPipeSet
}

// Get returns the string of stdin data
func (p *Pipe) Get() string {
	p.log.Info(data)
	return string(data)
}
