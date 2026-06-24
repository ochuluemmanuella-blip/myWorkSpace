package main

import (
	"fmt"
	"strings"
)

type Report struct {
	lines []string
}

func NewReport() *Report {
	return &Report{
		lines: []string{},
	}
}

func (r *Report) AddLines(lines ...string) *Report {
	// append all lines, return r for chaining
	r.lines = append(r.lines, lines...)
	return r
}

func (r *Report) String() string {
	var sb strings.Builder
	for _, line := range r.lines {
		sb.WriteString(line)
		sb.WriteString("\n")
	}
	// use strings.Builder to join all lines with "\n"
	return sb.String()
}
func main() {
	report := NewReport().AddLines("Header").AddLines("Line 1", "Line 2").String()
	fmt.Println(report)
}
