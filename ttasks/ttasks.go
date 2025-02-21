package ttasks

import (
	"time"

	"github.com/RichardKnop/machinery/v2/log"
)

// Add ...
func Add(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}

	//add delay of 5 senconds
	time.Sleep(5 * time.Second)

	log.INFO.Println("Sum of", args, "is", sum)
	return sum, nil
}
