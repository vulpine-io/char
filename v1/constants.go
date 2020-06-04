package char

// Control Characters (0-31).
const (
	Null byte = iota
	StartHeading
	StartText
	EndText
	EndTransmission
	Enquiry
	Acknowledgment
	Bell
	Backspace
	HorizontalTab
	LineFeed
	VerticalTab
	FormFeed
	CarriageReturn
	ShiftOut
	ShiftIn
	DataLineEscape
	DeviceControl1
	DeviceControl2
	DeviceControl3
	DeviceControl4
	NegativeAcknowledgement
	SynchronousIdle
	EndTransmitBlock
	Cancel
	EndMedium
	Substitute
	Escape
	FileSeparator
	GroupSeparator
	RecordSeparator
	UnitSeparator

	Tab = HorizontalTab
)

// Number Characters (48-57).
const (
	Zero byte = '0' + iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
)

// Lowercase Letters.
const (
	LowerA byte = 'a' + iota
	LowerB
	LowerC
	LowerD
	LowerE
	LowerF
	LowerG
	LowerH
	LowerI
	LowerJ
	LowerK
	LowerL
	LowerM
	LowerN
	LowerO
	LowerP
	LowerQ
	LowerR
	LowerS
	LowerT
	LowerU
	LowerV
	LowerW
	LowerX
	LowerY
	LowerZ
)

// Uppercase Letters.
const (
	UpperA byte = 'A' + iota
	UpperB
	UpperC
	UpperD
	UpperE
	UpperF
	UpperG
	UpperH
	UpperI
	UpperJ
	UpperK
	UpperL
	UpperM
	UpperN
	UpperO
	UpperP
	UpperQ
	UpperR
	UpperS
	UpperT
	UpperU
	UpperV
	UpperW
	UpperX
	UpperY
	UpperZ
)

// ASCII Printable Symbols 1 (32-47).
const (
	Space byte = ' ' + iota
	ExclamationMark
	DoubleQuote
	Number
	Dollar
	Percent
	Ampersand
	SingleQuote
	ParenthesisOpen
	ParenthesisClose
	Asterisk
	Plus
	Comma
	Hyphen
	Period
	Slash

	Amp              = Ampersand
	Apos             = SingleQuote
	Bang             = ExclamationMark
	ForwardSlash     = Slash
	FullStop         = Period
	Hash             = Number
	ParenthesisLeft  = ParenthesisOpen
	ParenthesisRight = ParenthesisClose
	Pound            = Number
)

// ASCII Printable Symbols 2 (58-64).
const (
	Colon byte = ':' + iota
	Semicolon
	LessThan
	Equals
	GreaterThan
	QuestionMark
	AtSymbol

	AngledBracketOpen  = LessThan
	AngledBracketLeft  = LessThan
	AngledBracketClose = GreaterThan
	AngledBracketRight = GreaterThan
	Question           = QuestionMark
)

// ASCII Printable Symbols 3 (91-96).
const (
	SquareBracketOpen byte = '[' + iota
	Backslash
	SquareBracketClose
	Caret
	Underscore
	Grave

	Backtick           = Grave
	Circumflex         = Caret
	BracketOpen        = SquareBracketOpen
	BracketClose       = SquareBracketClose
	BracketLeft        = SquareBracketOpen
	BracketRight       = SquareBracketClose
	SquareBracketLeft  = SquareBracketOpen
	SquareBracketRight = SquareBracketClose
)

// ASCII Printable Symbols 4 (123-127).
const (
	CurlyBraceOpen byte = '{' + iota
	VerticalBar
	CurlyBraceClose
	Tilde
	Delete

	BraceOpen       = CurlyBraceOpen
	BraceClose      = CurlyBraceClose
	BraceLeft       = CurlyBraceOpen
	BraceRight      = CurlyBraceClose
	CurlyBraceLeft  = CurlyBraceOpen
	CurlyBraceRight = CurlyBraceClose
	Pipe            = VerticalBar
)
