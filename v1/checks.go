package char

// IsDecimalDigit returns true if the character is in the set [0-9], else
// returns false.
func IsDecimalDigit(b byte) bool {
	return Zero <= b && b <= Nine
}

// IsLowerLetter returns true if the character is in the set [a-z], else returns
// false.
func IsLowerLetter(b byte) bool {
	return LowerA <= b && b <= LowerZ
}

// IsUpperLetter returns true if the character is in the set [A-Z], else returns
// false.
func IsUpperLetter(b byte) bool {
	return UpperA <= b && b <= UpperZ
}

// IsInlineWhitespace returns true if the character is in the set [ \t], else
// returns false.
func IsInlineWhitespace(b byte) bool {
	return b == Space || b == Tab
}

// IsLineBreak returns true if the character is in the set [\r\n], else returns
// false.
func IsLineBreak(b byte) bool {
	return b == LineFeed || b == CarriageReturn
}

// IsLetter returns true if the character is in the set [A-Za-z], else returns
// false.
func IsLetter(b byte) bool {
	return IsLowerLetter(b) || IsUpperLetter(b)
}

// IsOctalDigit returns true if the character is in the set [0-7], else returns
// false.
func IsOctalDigit(b byte) bool {
	return Zero <= b && b <= Seven
}

// IsHexDigit returns true if the character is in the set [0-9A-Fa-f], else
// returns false.
func IsHexDigit(b byte) bool {
	return IsDecimalDigit(b) ||
		(UpperA <= b && b <= UpperF) ||
		(LowerA <= b && b <= LowerF)
}

// IsWordCharacter returns true if the character is in the set [0-9A-Za-z_],
// else returns false.
func IsWordCharacter(b byte) bool {
	return IsLetter(b) ||
		IsDecimalDigit(b) ||
		b == Underscore
}
