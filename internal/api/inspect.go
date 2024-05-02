package api

func InspectPokemon(pokemon string) error {

	err := PokeStore.Inspect(pokemon)
	return err

}
