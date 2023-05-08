package error_handling

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func someFunc(data string) error {
	return ErrNetwork{
		Code: AuthFailureCode,
		Msg:  "This is a message",
	}
}

func restCall(data string) error {
	if err := someFunc(data); err != nil {
		return fmt.Errorf("restCall(%s) had an error: %w", data, err)
	}
	return nil
}

func main() {

	data := "This is some data."

	for {
		if err := restCall(data); err != nil {
			var netErr ErrNetwork
			if errors.As(err, &netErr) {
				log.Println("network error: ", err)
				time.Sleep(1 * time.Second)
				continue
			}
			log.Println("unrecoverable: ", err)
		}
	}
}
