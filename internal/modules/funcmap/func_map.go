package funcmap

import (
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/ast"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/token"
)

type Fns struct {
	fnmap map[string]*ast.Functionstatment
}

func (f *Fns) New() {
	f.fnmap = make(map[string]*ast.Functionstatment)
}
func (f *Fns) Add(n string, t token.Tokentype, fn *ast.Functionstatment) {
	str := n
	for _, i := range fn.Param {
		str += string(i.Type)
	}
	f.fnmap[str] = fn
}

func (f *Fns) Exists(n string, param []*ast.Identity) bool {
	str := n
	for _, i := range param {
		str += string(i.Type)
	}
	_, ok := f.fnmap[str]
	return ok
}

func (f *Fns) Getreturntype(n string, param []*ast.Identity) token.Tokentype {
	str := n
	for _, i := range param {
		str += string(i.Type)
	}
	i, _ := f.fnmap[str]
	return i.Tokentype()
}
