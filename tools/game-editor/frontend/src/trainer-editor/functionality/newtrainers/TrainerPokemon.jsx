import { useState, useEffect} from "react"

import PokemonStats from "./pokemon/PokemonStats"
const TrainerPokemon = ({ setPokemons, pokemonSpeciesList, heldItemsList, pokemonIndex, setPokemonIndex }) => {
    const [currentlySelectedPokemon, setCurrentlySelectedPokemon] = useState({})
    const SetData = (e) => {
        const pokemon = pokemonSpeciesList.find((pokemon) => pokemon.Name === e.target.value)
        setCurrentlySelectedPokemon(pokemon)
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
        {Object.keys(currentlySelectedPokemon).length >= 1 ? <PokemonStats currentlySelectedPokemon={currentlySelectedPokemon} heldItemsList={heldItemsList}/> : <></>}
        </>
    )
}
export default TrainerPokemon