import { useState } from "react";
import TrainerPokemon from "./TrainerPokemon"
const Generator = ({ pokemonSpeciesList, heldItemsList, pokemonsCount, dictData, setDictData}) => {
    const [pokemonIndex, setPokemonIndex] = useState(1)
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
            <TrainerPokemon  pokemonSpeciesList={pokemonSpeciesList} heldItemsList={heldItemsList} key={pokemonIndex} pokemonIndex={pokemonIndex} setPokemonIndex={setPokemonIndex} pokemonsCount={pokemonsCount} dictData={dictData} setDictData={setDictData}/>
        </div>
        <br/>
        <br/>
        </>
    )
}
export default Generator