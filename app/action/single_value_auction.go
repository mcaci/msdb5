package action

func SingleValueAction(rq interface{ Value() string }, a interface{ ValueSet(string) }) {
	a.ValueSet(rq.Value())
}
