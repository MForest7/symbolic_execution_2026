package ssa

import (
	"fmt"
	"strings"

	"golang.org/x/tools/go/ssa"
)

func SprintSSA(fn *ssa.Function) string {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, "Функция %s\n", fn.Name())

	for _, block := range fn.Blocks {
		fmt.Fprintf(sb, "Блок %d: %s -> %d -> %s\n", block.Index, fmt.Sprint(block.Preds), block.Index, fmt.Sprint(block.Succs))

		for _, instr := range block.Instrs {
			s := instr.String()
			if v, ok := instr.(ssa.Value); ok {
				s = fmt.Sprintf("%s = %s", v.Name(), instr.String())
			}
			fmt.Fprintf(sb, "\t%s\n", s)
		}
	}

	return sb.String()
}
