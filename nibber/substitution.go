package nibber

import (
	"sort"
	"strings"
)

type Substitution struct {
	Origin      string
	Destination string
}

type OrderedSubstitution []Substitution

func (os OrderedSubstitution) Len() int {
	return len(os)
}

func (os OrderedSubstitution) Swap(i, j int) {
	os[i], os[j] = os[j], os[i]
}

// less is more
func (os OrderedSubstitution) Less(i, j int) bool {
	return len(os[i].Origin) > len(os[j].Origin)
}

func (os *OrderedSubstitution) Order() {
	sort.Sort(os)
}

func (os OrderedSubstitution) toReplacerArray() []string {
	res := []string{}
	for _, elem := range os {
		res = append(res, elem.Origin, elem.Destination)
	}
	return res
}

func mapToOrderedSubstitution(m map[string]string) OrderedSubstitution {
	var os OrderedSubstitution

	for key, value := range m {
		os = append(os, Substitution{strings.ToUpper(key), value})
	}

	return os
}
