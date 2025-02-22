package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := GetNoteData()
	UserNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	UserNote.Display()
}
func GetNoteData() (string, string) {
	title := GetUSerData("Enter the Note title")
	content := GetUSerData("Enter thre content of the Note")
	return title, content
}
func GetUSerData(promt string) string {
	fmt.Println(promt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")

	return text
}

// some key point scanln() is only use for the single line input or the short input
//for the long input we use
// reader:=bufio.NewReader(os.Stdin)
// text,err:=reader.ReadString('\n')
// if err !=nil{
// 	return ""
// }
// text=strings.TrimSuffix(text,"\n")
// text=strings.TrimSuffix(text,"\n")

// return text
