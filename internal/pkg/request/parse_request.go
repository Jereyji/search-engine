package request

import (
	"fmt"
	"strings"
)

func ParseRequest(input string) (*Request, error) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil, fmt.Errorf("zero command")
	}

	var (
		command = parts[0]
		args    = parts[1:]
	)

	req := NewRequest(command)
	for _, arg := range args {
		keyValue := strings.Split(arg, "=")
		if len(keyValue) != 2 {
			return nil, fmt.Errorf("incorrect flag: %s", arg)
		}

		values := strings.Split(keyValue[1], ",")
		req.SetFlag(keyValue[0], values[0])
	}

	return req, nil
}
