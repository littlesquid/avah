package usecase

import (
	remind "avah/usecase/reminder"
	"fmt"
	"net/http"
)

func Run(w http.ResponseWriter, r *http.Request) {
	fmt.Println("job triggered")

	remind.ExecuteDailyReminder()
}
