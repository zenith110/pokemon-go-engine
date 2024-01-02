import Generator from "./Generator"
const TrainerPokemonsGenerator = ({ setPokemons, pokemonSpeciesList, heldItemsList, pokemonsCount }) => {
    return(
        <>
        { pokemonsCount >= 1 ? <Generator setPokemons={setPokemons} pokemonSpeciesList={pokemonSpeciesList} heldItemsList={heldItemsList} pokemonsCount={pokemonsCount}/>: <></>}
        </>
    )
}

export default TrainerPokemonsGenerator