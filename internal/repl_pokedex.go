package internal

var GlobalPokedex = &Pokedex{
	Data: make(map[string]Pokemon),
}

func (p *Pokedex) GetPokemon(name string) (Pokemon, bool) {
	pokemon, exists := p.Data[name]
	return pokemon, exists
}

func (p *Pokedex) AddPokemon(pokemon Pokemon) {
	p.Data[pokemon.Name] = pokemon
}
