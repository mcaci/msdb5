package action

type Join func(string)

func (j Join) ValueSet(val string) {
	j(val)
}
