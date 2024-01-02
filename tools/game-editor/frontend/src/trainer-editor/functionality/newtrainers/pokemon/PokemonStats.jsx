const PokemonStats = ({ currentlySelectedPokemon }) => {
    return(
        <>
        <br/>
        <label>HP:</label>
        <input type="number" value={currentlySelectedPokemon.HP} max={300}></input>
        <br/>
        <br/>
        <label value={currentlySelectedPokemon.Attack} max={300}>Attack:</label>
        <input type="number"></input>
        <br/>
        <br/>
        <label>Defense:</label>
        <input type="number" value={currentlySelectedPokemon.Defense} max={300}></input>
        <br/>
        <br/>
        <label>SpecialAtk:</label>
        <input type="number" value={currentlySelectedPokemon.SpecialAttack} max={300}></input>
        <br/>
        <br/>
        <label>SpecialDef:</label>
        <input type="number" value={currentlySelectedPokemon.SpecialDefense} max={300}></input>
        <br/>
        <br/>
        <label>Speed:</label>
        <input type="number" value={currentlySelectedPokemon.Speed} max={300}></input>
        <br/>
        <br/>
        <label>Level:</label>
        <input type="number" max={100} min={1}></input>
        <br/>
        <br/>
        <label>Move1:</label>
        <select name="moves1" defaultValue={"placeholder"}>
        <option value={"placeholder"} disabled>Select a move</option>
        {/* {currentlySelectedPokemon.Moves.map((moves) => console.log(moves))} */}
        </select>
        </>
    )
}
export default PokemonStats