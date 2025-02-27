package utils

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

func readStdin(ctx context.Context) string {
	input := make(chan string)

	go func(lines chan string) {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		lines <- s.Text()
	}(input)

	defer close(input)

	for {
		select {
		case line := <-input:
			return line
		case <-ctx.Done():
			return ""
		}
	}
}

// ValidateFunc is a function that validates the input string.
// It returns a boolean indicating whether the input is valid, a message to display to the user, and an error.
type ValidateFunc func(string) (ok bool, message string, err error)

func Ask(ctx context.Context, question string, allowEmpty bool, validate ValidateFunc) (string, error) {
	fmt.Println("")

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}

		fmt.Print(question) //nolint:forbidigo

		result := readStdin(ctx)

		result = strings.TrimSpace(result)

		if allowEmpty && result == "" {
			return result, nil
		}

		if validate != nil {
			ok, message, err := validate(result)
			if err != nil {
				return "", err
			}
			if ok {
				return result, nil
			}

			fmt.Println(message)

			continue
		}

		if result != "" {
			return result, nil
		}
	}
}
