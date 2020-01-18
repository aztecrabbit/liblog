package liblog

import (
	"fmt"
	"time"
	"strings"
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

func LogColor(message string, color string) {
	messages := strings.Split(message, "\n")

	for _, value := range messages {
		fmt.Printf("\r%s%s%s%s\n", "\033[K", color, value, Colors["CC"])
	}
}

func LogInfo(message string, info string, color string) {
	datetime := time.Now()
	LogColor(
		fmt.Sprintf("[%.2d:%.2d:%.2d] %[4]s::%[5]s %[6]s%[7]s %[4]s::%[5]s %[6]s%[8]s",
			datetime.Hour(), datetime.Minute(), datetime.Second(),
			Colors["P1"], Colors["CC"], color,
			info, message),
		color,
	)
}

func Log(message string) {
	LogColor(message, Colors["G1"])
}
