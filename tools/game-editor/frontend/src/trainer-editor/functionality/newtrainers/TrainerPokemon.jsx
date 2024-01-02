import { useState } from "react"

import PokemonStats from "./pokemon/PokemonStats"
const TrainerPokemon = ({ setPokemons, pokemonSpeciesList, heldItemsList, pokemonIndex }) => {
    const [currentlySelectedPokemon, setCurrentlySelectedPokemon] = useState({})
    const SetData = (e) => {
        const pokemon = pokemonSpeciesList.find((pokemon) => pokemon.Name === e.target.value)
        console.log(pokemon)
        setCurrentlySelectedPokemon(pokemon)
        console.log(currentlySelectedPokemon)
    }
    return(
        <>
        <p>Pokemon #{pokemonIndex}</p>
        <select name="pokemons" onChange={(e) => SetData(e)} defaultValue={"placeholder"}>
            <option value={"placeholder"} disabled>Select a pokemon</option>
            {pokemonSpeciesList.map((pokemon) =>
                <option value={pokemon.Name} key={pokemon.Name}>{pokemon.Name}</option>
            )}
        </select>
        {currentlySelectedPokemon.length >= 1 ? <PokemonStats currentlySelectedPokemon={currentlySelectedPokemon}/> : <></>}
        </>
    )
}
export default TrainerPokemon