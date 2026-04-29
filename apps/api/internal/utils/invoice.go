package utils

import (
	"fmt"
	"time"
)

func GenerateInvoiceNumber() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")

	return fmt.Sprintf(
		"INV-%s",
		time.Now().
			In(loc).
			Format("20060102150405"),
	)
}
