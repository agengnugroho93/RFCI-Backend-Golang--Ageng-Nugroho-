package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rcfi/models"
	"rcfi/repository"
)

const (
	indexPage = "Selamat Datang"
)

// deklarasi dan handling home ("/")
func homeHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, indexPage)
}

func blogsHandler(response http.ResponseWriter, request *http.Request) {
	status := "Success:"
	scheme := "POST"
	baseurl := "http://127.0.0.0:8080/blogs"
	payload := request.FormValue("body")

	var messages models.Messages
	var author models.Author
	var title models.Title

	json.Unmarshal([]byte(payload), &messages)
	json.Unmarshal([]byte(payload), &author)
	json.Unmarshal([]byte(payload), &title)
	messagesjson, _ := json.Marshal(messages)
	authorjson, _ := json.Marshal(author)
	titlejson, _ := json.Marshal(title)

	authorstr := (string(authorjson))
	titlestr := (string(titlejson))
	messagestr := (string(messagesjson))

	err := repository.Logger(authorstr, titlestr, messagestr, status, scheme, baseurl)
	if err != nil {
		println("Failed to logging")
	}

	fmt.Fprintln(response, 201)
}
