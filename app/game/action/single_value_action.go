package action

func singleValueAction(rq interface{ Value() string }, a interface{ valueSet(string) }) {
	a.valueSet(rq.Value())
}
