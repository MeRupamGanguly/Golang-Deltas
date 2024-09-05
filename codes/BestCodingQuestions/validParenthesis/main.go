package main

import "fmt"

type stack struct {
	d []string
}

func NewStack(d []string) *stack {
	return &stack{
		d: d,
	}
}
func (s *stack) push(data string) {
	s.d = append(s.d, data)
}
func (s *stack) pop() string {
	if len(s.d) < 1 {
		return ""
	}
	t := s.d[len(s.d)-1]
	s.d = s.d[:len(s.d)-1]
	return t
}

func validParenthesesCheck(arr []string) bool {
	o := NewStack([]string{})
	m := make(map[string]string)
	m["{"] = "}"
	m["["] = "]"
	m["("] = ")"
	m["}"] = "{"
	m["]"] = "["
	m[")"] = "("
	for i := 0; i < len(arr); i++ {
		char := arr[i]
		if len(o.d) > 0 {
			oelem := o.pop()
			if m[char] != oelem {
				o.push(oelem)
				o.push(char)
			}
		} else {
			o.push(char)
		}
	}
	if len(o.d) == 0 {
		return true
	} else {
		return false
	}
}
func main() {
	example1 := []string{"(", "{", "[", "]", "}", ")"}
	example2 := []string{"(", "{", "[", "]", "}", "[", ")"}
	fmt.Println(validParenthesesCheck(example1))
	fmt.Println(validParenthesesCheck(example2))
}
