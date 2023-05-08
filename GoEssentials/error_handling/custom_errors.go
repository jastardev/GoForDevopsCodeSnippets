package error_handling

import (
	"errors"
	"fmt"
	"log"
)

const (
	UnknownCode     = 0
	UnreachableCode = 1
	AuthFailureCode = 2
)

type ErrNetwork struct {
	Code int
	Msg  string
}

func (e ErrNetwork) Error() string {
	return fmt.Sprintf("network error(%d): %s", e.Code, e.Msg)
}

func auth() (string, error) {
	return "failure", ErrNetwork{
		Code: AuthFailureCode,
		Msg:  "User unrecognized",
	}
}

func main() {
	var netErr ErrNetwork

	if errors.As(err, &netErr) {
		if netErr.Code == AuthFailureCode {
			log.Println("Failure due to authentication")
			break
		}
		log.Printf("Recoverable error: %s\n", netErr)
	}
	log.Printf("unrecoverable error: %s\n", err)
}
