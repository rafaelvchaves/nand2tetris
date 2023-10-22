package symbol

import "fmt"

type Table struct {
	nextSymbolAddress int16
	table             map[string]int16
}

func NewTable() *Table {
	table := make(map[string]int16)
	return &Table{
		nextSymbolAddress: 16,
		table:             table,
	}
}

func (t *Table) SymbolAddress(s string) string {
	v, ok := t.table[s]
	if ok {
		return fmt.Sprint(v)
	}
	t.table[s] = t.nextSymbolAddress
	t.nextSymbolAddress++
	return fmt.Sprint(t.table[s])
}

func (t *Table) Get(s string) (string, bool) {
	v, ok := t.table[s]
	if !ok {
		return "", false
	}
	return fmt.Sprint(v), true
}

func (t *Table) Put(s string, v int16) {
	t.table[s] = v
}
