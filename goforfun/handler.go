package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("Hello, again Go.Last try brahs. You said: %s", string(req))
}
