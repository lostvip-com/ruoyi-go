package color

import (
	"fmt"
)

const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

func Black(msg string) string {
	return SetColor(msg, 0, 0, TextBlack)
}

func Red(msg string) string {
	return SetColor(msg, 0, 0, TextRed)
}

func Green(msg string) string {
	return SetColor(msg, 0, 0, TextGreen)
}

func Yellow(msg string) string {
	return SetColor(msg, 0, 0, TextYellow)
}

func Blue(msg string) string {
	return SetColor(msg, 0, 0, TextBlue)
}

func Magenta(msg string) string {
	return SetColor(msg, 0, 0, TextMagenta)
}

func Cyan(msg string) string {
	return SetColor(msg, 0, 0, TextCyan)
}

func White(msg string) string {
	return SetColor(msg, 0, 0, TextWhite)
}

func SetColor(msg string, conf, bg, text int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
}

var LogoContent = "\n ____  ____                                     _                  __         _                 __  \n|_   ||   _|                                   (_)                [  |       (_)               |  ] \n  | |__| |  __   _   _ .--..--.   _ .--..--.   __   _ .--.   .--./)| |.--.   __   _ .--.   .--.| |  \n  |  __  | [  | | | [ `.-. .-. | [ `.-. .-. | [  | [ `.-. | / /'`\\;| '/'`\\ \\[  | [ `/'`\\]/ /'`\\' |  \n _| |  | |_ | \\_/ |, | | | | | |  | | | | | |  | |  | | | | \\ \\._//|  \\__/ | | |  | |    | \\__/  |  \n|____||____|'.__.'_/[___||__||__][___||__||__][___][___||__].',__`[__;.__.' [___][___]    '.__.;__] \n                                                           ( ( __))                                 \n"

var MQTTBrokerLogoContext = "\n ____                              \n/ ___| _   _  ___ ___ ___  ___ ___ \n\\___ \\| | | |/ __/ __/ _ \\/ __/ __|\n ___) | |_| | (_| (_|  __/\\__ \\__ \\\n|____/ \\__,_|\\___\\___\\___||___/___/\n                                   "
