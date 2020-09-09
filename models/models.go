package models

// {"author": "nama","title": "judul","comments": [{"message": "message1"},{"message": "message2"}]}

type (
	// Messages struct declare
	Messages struct {
		Comments []Message `json:"comments"`
	}

	// Author struct declare
	Author struct {
		Author string `json:"author"`
	}

	// Title struct declare
	Title struct {
		Title string `json:"title"`
	}

	// Message struct declare
	Message struct {
		Message string `json:"message"`
	}
)
