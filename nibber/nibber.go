package nibber

import "strings"

type Nibber struct {
	substitutions OrderedSubstitution
	replacer      strings.Replacer
}

func NewNibber(subsMap map[string]string) Nibber {
	subs := MapToOrderedSubstitution(subsMap)
	subs.Order()

	return Nibber{
		substitutions: subs,
		replacer:      *strings.NewReplacer(subs.ToReplacerArray()...),
	}
}

func (n Nibber) Nibbering(str string) string {
	return n.replacer.Replace(str)
}
