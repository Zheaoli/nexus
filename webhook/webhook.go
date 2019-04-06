package webhook

import "fmt"

type WBStrategy interface {
	ParseMessage() (WBMessage, error)
	fmt.Stringer
}

type WBMessage interface {
	Parse() string
	fmt.Stringer
}
