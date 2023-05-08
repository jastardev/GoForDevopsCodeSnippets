package error_handling

import "errors"

var (
	//ErrNetwork = errors.New("Network Error")
	ErrInput = errors.New("Input Error")
)

func main() {
	for {
		err := someFunc("data")
		if err == nil {
			break // request was successful
		}
		if errors.Is(err, ErrInput) {
			//this might be a recoverable input error, lets try again.
			continue
		}
		// A non-recoverable error occurred.
		break
	}
}
