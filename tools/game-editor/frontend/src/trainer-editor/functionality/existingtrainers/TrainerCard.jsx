import PokemonSelection from "./PokemonSelection"
import PokemonStats from "./PokemonStats";
import { useState } from "react"
const TrainerCard = ({ trainerName, trainerId, pokemons, sprite, location, nextFightId }) => {
const [currentlySelectedPokemon, setCurrentlySelectedPokemon ] = useState(null);
const [visibleStats, setVisibleStats ] = useState(false)
const [pokemonData, setPokemonData ] = useState([])
return(
    <>
    <p>
    <label>Name: </label>
    <p>{trainerName}</p>
    <br/>
    <label>Id: </label>
    <p>{trainerId}</p>
    <img src={sprite} />
    {pokemons.maps((pokemon, index) => {
        <>
        <PokemonSelection pokemonData={pokemon} setCurrentlySelectedPokemon={setCurrentlySelectedPokemon} setVisible={setVisibleStats} index={index} setPokemonData={setPokemonData}/>
        </>
    })}
    <label>Found in: </label>
    <p>{location}</p>
    <br/>
    <label>Next Battle: </label>
    <p>{nextFightId}</p>
    </p>
    <PokemonStats pokemonData={pokemonData} />
    </>
)
}
export default TrainerCard