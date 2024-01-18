package src

type Inctruct struct {
	code string
}

type tokens struct {
	stack []string
}

func (t *tokens) push(n string) {

	t.stack = append(t.stack, n)

}
func (t *tokens) pop() string {
	if len(t.stack)-1 != 0 {
		s := t.stack[len(t.stack)-1]
		t.stack = t.stack[:len(t.stack)-1]

		return s
	} else {
		return ""
	}
}
