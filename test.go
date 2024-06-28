package tester

type Slice []Element
type Element int

func Sum(s []Element) (res Element) {
	for _, v := range s {
		res += v
	}
	return
}

func SliceFunction(s []Element, op func(Element) Element) {
	for i, v := range s {
		s[i] = op(v)
	}
}

func SliceOperation(s []Element, op func(Element, Element) Element, init Element) Element {
	res := op(init, s[0])
	for i := 1; i < len(s); i++ {
		res = op(res, s[i])
	}
	return res
}
