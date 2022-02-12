package webhook

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Reply(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is reply function!")

	requestBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	fmt.Printf("got request: %s \n", string(requestBody[:]))
}
