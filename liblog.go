package liblog

import (
	"os"
	"fmt"
	"time"
	"strings"
	"golang.org/x/sys/unix"
)

var (
	Colors = map[string]string {
        "R1": "\033[31;1m", "R2": "\033[31;2m",
        "G1": "\033[32;1m", "G2": "\033[32;2m",
        "Y1": "\033[33;1m", "Y2": "\033[33;2m",
        "B1": "\033[34;1m", "B2": "\033[34;2m",
        "P1": "\033[35;1m", "P2": "\033[35;2m",
        "C1": "\033[36;1m", "C2": "\033[36;2m", "CC": "\033[0m",
	}
)

func GetTerminalSize() *unix.Winsize {
	terminal_size, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		// return nil, os.NewSyscallError("GetWinsize", err)
		panic(err)
	}

	return terminal_size
}

func LimitMessageLength(message string, slice int) string {
	terminal_size := GetTerminalSize()
	terminal_width := int(terminal_size.Col)

	if len(message) > terminal_width {
		message = message[:terminal_width - slice - 4] + "..."
	}

	return message
}

func Log(message string, color string, suffix string) {
	fmt.Printf("%s%s%s%s%s%s", "\r", "\033[K", color, message, Colors["CC"], suffix)
}

func LogColor(message string, color string) {
	messages := strings.Split(message, "\n")

	for _, value := range messages {
		Log(value, color, "\n")
	}
}

func Header(messages []string, color string) {
	LogColor("\033[2J" + "\033[H" + strings.Join(messages, "\n") + "\n", color)
}

func LogInfo(message string, info string, color string) {
	datetime := time.Now()
	LogColor(
		fmt.Sprintf("[%.2d:%.2d:%.2d]%[5]s %[4]s::%[5]s %[6]s%[7]s%[5]s %[4]s::%[5]s %[6]s%[8]s",
			datetime.Hour(), datetime.Minute(), datetime.Second(),
			Colors["P1"], Colors["CC"], color,
			info, message),
		color,
	)
}

func LogInfoSplit(message string, slice int, info string, color string) {
	terminal_size := GetTerminalSize()
	terminal_width := int(terminal_size.Col) - slice

	var data string

	for {
		if len(message) > terminal_width {
			data, message = message[:terminal_width], strings.TrimSpace(message[terminal_width:])
			LogInfo(data, info, color)
			continue
		}

		LogInfo(message, info, color)
		break
	}
}

func LogKeyboardInterrupt() {
    LogInfo(
    	"Keyboard Interrupt\n\n" +
    		"|   Ctrl-C again if not exiting automaticly\n" +
    		"|   Please wait...\n|\n",
    	"INFO", Colors["R1"],
    )
}

func LogException(err error, info string) {
	LogInfo(fmt.Sprintf("Exception:\n\n|   %v\n|\n", err), info, Colors["R1"])
}

func LogReplace(message string, color string) {
	Log(LimitMessageLength(message, 0), color, "\r")
}
