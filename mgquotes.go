package mgquotes

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Nice to meet you, %v. It's an honor to meet a living legend like yourself.", name)
	return message
}
