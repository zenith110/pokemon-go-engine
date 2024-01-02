import TrainerPokemon from "./TrainerPokemon"
const Generator = ({ setPokemons, pokemonSpeciesList, heldItemsList, pokemonsCount}) => {
    const divStyle={
        overflowY: 'scroll',
        border:'1px solid red',
        width:'500px',
        float: 'center',
        height:'500px',
        position:'relative',
        textAlign: 'center'
      };
    return(
        <>
        <div style={divStyle}>
            {[...Array(pokemonsCount)].map((x) => 
            <TrainerPokemon setPokemons={setPokemons} pokemonSpeciesList={pokemonSpeciesList} heldItemsList={heldItemsList} key={x} pokemonIndex={x}/>
            )}
        </div>
        <br/>
        <br/>
        </>
    )
}
export default Generator