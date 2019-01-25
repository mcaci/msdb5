package player


// Players struct
type Players []*Player

// Add func
func (set *Players) Add(p Player) {
	*set = append(*set, &p)
}
