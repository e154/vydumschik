package vydumschik

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var (
	Paragraphs = [...]string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam sit amet mi sed mi mattis suscipit. Integer lacinia lacinia erat. Etiam a arcu. Quisque et augue eget nisi luctus laoreet. Vestibulum sit amet quam eu nibh vehicula elementum. Vivamus urna pede, scelerisque eu, hendrerit euismod, vehicula vel, dolor.",
		"Quisque blandit. Nulla eget odio eu nisl sollicitudin tempus. Donec libero. Maecenas dui elit, venenatis eget, congue vel, semper sed, est. Nunc felis.",
		"Donec ante. Nullam rutrum, nibh eu tempor egestas, diam pede luctus ante, vitae fringilla ligula pede at turpis. Morbi sed metus quis odio luctus lacinia. In hac habitasse platea dictumst. Ut erat ipsum, faucibus sit amet, interdum in, tempor rutrum, metus.",
		"Maecenas ullamcorper leo et libero. Proin nec eros. In id lorem eu tellus tempus bibendum. Curabitur eu odio. Curabitur fermentum sem ut arcu. Ut eget leo ultricies mauris malesuada sollicitudin. Vivamus fringilla auctor nisl. Curabitur venenatis. Duis tempus, nunc quis convallis posuere, eros nibh lacinia nisl, non molestie enim lorem at nisi. Aenean dui. Cras elementum felis sit amet tellus consequat vehicula. Mauris rhoncus. Donec posuere mauris adipiscing tortor. Curabitur sagittis leo in magna.",
		"Cras quis libero ut urna pulvinar mattis. Vestibulum suscipit gravida nulla. Proin condimentum sapien ut augue. Cras turpis ligula, pharetra in, malesuada at, eleifend vel, pede. Phasellus turpis nisi, placerat in, semper sed, tincidunt a, tortor. Vivamus est nibh, mattis vitae, aliquet non, dignissim at, velit. Etiam elementum odio non erat. Donec varius felis sed nisl. Vestibulum imperdiet. Donec viverra pede ac tortor.",
	}
)

type Lorem struct {}

// слово
func (l *Lorem) Worlds(num int) string {

	if num == 0 {
		return ""
	}

	var worlds []string
	result := ""
	k := 0
	for i:=0;i<num;i++ {
		if len(worlds) == 0 || k >= len(worlds) {
			speech := l.Speech(1)
			worlds = strings.Split(speech[:utf8.RuneCountInString(speech)-1], " ")
			k=0
		}

		result += worlds[k]
		if i < num - 1 {
			result += " "
		}

		k++
	}

	if result[utf8.RuneCountInString(result)-1:] == "," {
		result = result[:utf8.RuneCountInString(result)-1]
	}

	return result + ". "
}

func (l *Lorem) Paragraphs(num int) string {

	if num == 0 {
		return ""
	}

	result := ""
	for i:=0;i<num;i++ {
		br := ""
		if i < num - 1 {
			br = "\n\n"
		} else {
			br = "\n"
		}

		result += fmt.Sprintf("%s%s", Paragraphs[random(0, len(Paragraphs))], br)
	}

	return result
}

func (l *Lorem) Bytes(num int) string {

	if num == 0 {
		return ""
	}

	result := ""
	speech := ""
	k := 0
	for i:=0;i<num;i++ {
		if speech == "" || k >= len(speech) - 2 {
			if speech != "" {
				result += ". "
			}
			speech = l.Speech(1)
			k=0
		}

		result += speech[k:k+1]

		if len(result) >= num {
			break
		}

		k++
	}

	return result
}

// предложение
func (l *Lorem) Speech(num int) string {

	if num == 0 {
		return ""
	}

	paragraph := strings.Split(strings.Replace(l.Paragraphs(1), ". ", ".", -1), ".")
	return paragraph[random(0, len(paragraph) - 1)] + "."
}