package lcd2004

import (
	"math/bits"
	"time"

	"periph.io/x/periph/conn"
)

const En = (1 << 2)

// Device defines an interface for an i2c device
type Device interface {
	Duplex() conn.Duplex
	String() string
	Tx(w, r []byte) error
	Write(b []byte) (int, error)
}

// LCDDisplay defines a new LCD display
type LCDDisplay struct {
	dev         Device
	currentLine int
}

// New crates a new LCDDisplay with the given connection
func New(d Device) *LCDDisplay {
	return &LCDDisplay{d, 0}
}

// Init initalizes the display
func (l *LCDDisplay) Init() {
	l.write(0x03)
	l.write(0x03)
	l.write(0x03)
	l.write(0x02)

	l.write(LCDFunctionSet | LCD2Line | LCD5x8Dots | LCD4BitMode)
	l.write(LCDDisplayControl | LCDDisplayOn)
	l.write(LCDClearDisplay)
	l.write(LCDEntryModeSet | LCDEntryLeft)

	time.Sleep(200 * time.Millisecond)
}

// Clear the display
func (l *LCDDisplay) Clear() {
	l.write(LCDClearDisplay)
	l.write(LCDReturnHome)
	l.currentLine = 0
}

func (l *LCDDisplay) write(d int) error {
	l.writeFourBits(0x00 | (d & 0xF0))
	l.writeFourBits(0x00 | ((d << 4) & 0xF0))

	return nil
}

// Println prints the given string to the display
func (l *LCDDisplay) Println(s string) {
	switch l.currentLine {
	case 0:
		l.write(0x80)
		l.currentLine++
	case 1:
		l.write(0xC0)
		l.currentLine++
	case 2:
		l.write(0x94)
		l.currentLine++
	case 3:
		l.write(0xD4)
		l.currentLine = 0
	}

	for _, c := range s {
		l.writeChar(int(c))
	}
}

// DisplayOff switches the display off
func (l *LCDDisplay) DisplayOff() {
	l.dev.Write([]byte{LCDNoBacklight})
}

func (l *LCDDisplay) writeChar(d int) error {
	l.writeFourBits(0x01 | (d & 0xF0))
	l.writeFourBits(0x01 | ((d << 4) & 0xF0))

	return nil
}

func (l *LCDDisplay) writeFourBits(b int) error {
	l.dev.Write([]byte{byte(b | LCDBacklight)})
	l.strobe(b)

	return nil
}

func (l *LCDDisplay) strobe(b int) {
	l.dev.Write([]byte{byte(b | En | LCDBacklight)})
	time.Sleep(500 * time.Microsecond)

	l.dev.Write([]byte{byte((b & int(bits.Reverse(En))) | LCDBacklight)})
	time.Sleep(100 * time.Microsecond)
}
