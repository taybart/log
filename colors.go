package log

// Colors
const (
	// Color escape
	ce           = "\033[0;"
	ceBold       = "\033[1;"
	ceItalic     = "\033[3;"
	ceUnderlined = "\033[4;"
	ceBlinking   = "\033[5;"

	// Normal Colors
	Gray   = ce + "37m"
	Purple = ce + "35m"
	Blue   = ce + "34m"
	Yellow = ce + "33m"
	Green  = ce + "32m"
	Red    = ce + "31m"

	// Bold Colors
	BoldGray   = ceBold + "37m"
	BoldPurple = ceBold + "35m"
	BoldBlue   = ceBold + "34m"
	BoldYellow = ceBold + "33m"
	BoldGreen  = ceBold + "32m"
	BoldRed    = ceBold + "31m"

	// Italic Colors
	ItalicGray   = ceItalic + "37m"
	ItalicPurple = ceItalic + "35m"
	ItalicBlue   = ceItalic + "34m"
	ItalicYellow = ceItalic + "33m"
	ItalicGreen  = ceItalic + "32m"
	ItalicRed    = ceItalic + "31m"

	// Underlined Colors
	UnderlinedGray   = ceUnderlined + "37m"
	UnderlinedPurple = ceUnderlined + "35m"
	UnderlinedBlue   = ceUnderlined + "34m"
	UnderlinedYellow = ceUnderlined + "33m"
	UnderlinedGreen  = ceUnderlined + "32m"
	UnderlinedRed    = ceUnderlined + "31m"

	// Blinking Colors
	BlinkingGray   = ceBlinking + "37m"
	BlinkingPurple = ceBlinking + "35m"
	BlinkingBlue   = ceBlinking + "34m"
	BlinkingYellow = ceBlinking + "33m"
	BlinkingGreen  = ceBlinking + "32m"
	BlinkingRed    = ceBlinking + "31m"

	Bold       = "\033[1;3m"
	Italic     = "\033[3;3m"
	Underlined = "\033[4;1m"
	Blinking   = "\033[5;1m"
	// Return to default
	Rtd   = ce + "0m"
	Reset = ce + "0m"
)
