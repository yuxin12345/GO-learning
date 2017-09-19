package main
import (
	"errors"
	"fmt"
	"math"
)

// MySqrt -
func MySqrt(n float64) (sqrt float64, err error) {
	if n < 0 {
		err := errors.New("Invalid Number")

		return -1, err
	}

	return math.Sqrt(n), nil
}

func main() {
	sqrt, err := MySqrt(4)
	fmt.Println("MySqrt:", sqrt, "Error:", err)
	sqrt, err = MySqrt(-1)
	fmt.Println("MySqrt:", sqrt, "Error:", err)
}