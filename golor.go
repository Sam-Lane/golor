package golor

import "fmt"

const (
	reset   = "\u001b[0m"
	red     = "\u001b[31m"
	black   = "\u001b[30m"
	green   = "\u001b[32m"
	yellow  = "\u001b[33m"
	blue    = "\u001b[34m"
	magenta = "\u001b[35m"
	cyan    = "\u001b[36m"
	white   = "\u001b[37m"

	brightBlack   = "\u001b[30;1m"
	brightRed     = "\u001b[31;1m"
	brightGreen   = "\u001b[32;1m"
	brightYellow  = "\u001b[33;1m"
	brightBlue    = "\u001b[34;1m"
	brightMagenta = "\u001b[35;1m"
	brightCyan    = "\u001b[36;1m"
	brightWhite   = "\u001b[37;1m"

	rgb = "\033[38;2;%d;%d;%dm%s\033[39;49m\n"
)

// wrap a string with a color code and reset code
func colorWrap(color string, message string) string {
	return color + message + reset
}

// RGB return a RGB colored string
func RGB(red int, green int, blue int, message string) string {
	return fmt.Sprintf(rgb, red, green, blue, message)
}

// Red return red colored string
func Red(message string) string {
	return colorWrap(red, message)
}

// Black return black colored string
func Black(message string) string {
	return colorWrap(black, message)
}

// Green return green colored string
func Green(message string) string {
	return colorWrap(green, message)
}

// Yellow return yellow colored string
func Yellow(message string) string {
	return colorWrap(yellow, message)
}

// Blue return blue colored string
func Blue(message string) string {
	return colorWrap(blue, message)
}

// Magenta return magenta colored string
func Magenta(message string) string {
	return colorWrap(magenta, message)
}

// Cyan return cyan colored string
func Cyan(message string) string {
	return colorWrap(cyan, message)
}

// White return white colored string
func White(message string) string {
	return colorWrap(white, message)
}
