/*
 * @Author: dzw
 * @Date: 2020-03-01 10:57:15
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-01 11:17:07
 */

package main

import "fmt"

// People ...
// interface
type People interface {
	Speak(string) string
}

// Student ...
type Student struct{}

// Speak ...
func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

// ShowScore ...
func (stu Student) ShowScore() {
	fmt.Println("my score")
}

// Teacher ...
type Teacher struct{}

// Speak ...
func (t *Teacher) Speak(think string) (talk string) {
	if think == "as" {
		talk = "ahah"
	} else {
		talk = "we"
	}

	return
}

// Show ...
func Show(p People) {
	p.Speak("asd")
}

func main() {
	{
		var p People = &Student{}
		think := "asd"
		fmt.Println(p.Speak(think))

		var s Student = Student{}
		fmt.Println(s.Speak(think))
		s.ShowScore()

		var t Teacher = Teacher{}
		fmt.Println(t.Speak(think))
		var tp People = &t
		fmt.Println(tp.Speak(think))

		Show(&s)
	}

	{
		var p People
		var s = &Student{}
		p = s
	}
}
