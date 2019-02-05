package board

// Play func
func (b *Board) Play(number, seed, origin string) {
	p, _ := b.Players().Find(origin)
	c, _ := p.Play(number, seed)
	b.PlayedCards().Add(c)
}
