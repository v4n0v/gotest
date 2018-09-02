package main

import (
	"github.com/nu7hatch/gouuid"
	"fmt"
	"time"
)

var filelist = make(map[string]bool)

func main() {

	fmt.Println(formatDateOrNull("2018-09-21 21:00:00.000000"))
}
func formatDateOrNull(date string) string {
	if date != "null" {
		d, _ := time.Parse("2006-01-02", date)
		return d.String()
	}
	return date

}

func generateUUID() string {
	u, _ := uuid.NewV4()
	return u.String()
}
