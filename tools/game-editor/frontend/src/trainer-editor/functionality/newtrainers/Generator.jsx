import { useState } from "react";
import TrainerPokemon from "./TrainerPokemon"
const Generator = ({ setPokemons, pokemonSpeciesList, heldItemsList, pokemonsCount}) => {
    const [pokemonIndex, setPokemonIndex] = useState(0)
    const divStyle={
        overflowY: 'scroll',
        border:'1px solid red',
        width:'500px',
        float: 'center',
        height:'500px',
        position:'relative',
        textAlign: 'center',
        justifyContent: 'center'
      };
    return(
        <>
        <div style={divStyle}> 
            <TrainerPokemon setPokemons={setPokemons} pokemonSpeciesList={pokemonSpeciesList} heldItemsList={heldItemsList} key={x} pokemonIndex={pokemonIndex} setPokemonIndex={setPokemonIndex} />
        </div>
        <br/>
        <br/>
        </>
    )
}
export default Generator