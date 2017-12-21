package lcd2004

const (
	// commands
	LCDClearDisplay   = 0x01
	LCDReturnHome     = 0x02
	LCDEntryModeSet   = 0x04
	LCDDisplayControl = 0x08
	LCDCursorShift    = 0x10
	LCDFunctionSet    = 0x20
	LCDSetCgramAddr   = 0x40
	LCDSetDdramAddr   = 0x80

	// flags for display entry mode
	LCDEntryRight          = 0x00
	LCDEntryLeft           = 0x02
	LCDEntryShiftIncrement = 0x01
	LCDEntryShiftDecrement = 0x00

	// flags for display on/off control
	LCDDisplayOn  = 0x04
	LCDDisplayOff = 0x00
	LCDCursorOn   = 0x02
	LCDCursorOff  = 0x00
	LCDBlinkOn    = 0x01
	LCDBlinkOff   = 0x00

	// flags for display/cursor shift
	LCDDisplayMove = 0x08
	LCDCursorMove  = 0x00
	LCDMoveRight   = 0x04
	LCDMoveLeft    = 0x00

	// flags for function set
	LCD8BitMode = 0x10
	LCD4BitMode = 0x00
	LCD2Line    = 0x08
	LCD1Line    = 0x00
	LCD5x10Dots = 0x04
	LCD5x8Dots  = 0x00

	// flags for backlight control
	LCDBacklight   = 0x08
	LCDNoBacklight = 0x00
)
