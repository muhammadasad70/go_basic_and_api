package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

// we use the New() which automaticaly use for the making constructure

// this is for the creation of the structure
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid Input")
	}
	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil

}

// this is for the display the value of the structure
func (note Note) Display() {
	fmt.Printf("The Note title is: \n%v\nThe Note content is: \n%v\n", note.title, note.content)
}
