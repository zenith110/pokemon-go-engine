import { useState } from "react"
const PokemonStats = ({ currentlySelectedPokemon, heldItemsList}) => {
    const [move1, setMove1] = useState("")
    const [move2, setMove2] = useState("")
    const [move3, setMove3] = useState("")
    const [move4, setMove4] = useState("") 
    const [heldItem, setHeldItem] = useState("")
    const [hp, setHp] = useState(0)
    const [attack, setAttack] = useState(0)
    const [defense, setDefense] = useState(0)
    const [specialAtk, setSpecialAtk] = useState(0)
    const [specialDef, setSpecialDef] = useState(0)
    const [speed, setSpeed] = useState(0)
    const [level, setLevel] = useState(0)
    return(
        <>
        <br/>
        <label>HP:</label>
        <input type="number" value={currentlySelectedPokemon.HP} max={300} onChange={(e) => setHp(e.target.value)}></input>
        <br/>
        <br/>
        <label max={300}>Attack:</label>
        <input value={currentlySelectedPokemon.Attack} type="number" onChange={(e) => setAttack(e.target.value)}></input>
        <br/>
        <br/>
        <label>Defense:</label>
        <input type="number" value={currentlySelectedPokemon.Defense} max={300} onChange={(e) => setDefense(e.target.value)}></input>
        <br/>
        <br/>
        <label>SpecialAtk:</label>
        <input type="number" value={currentlySelectedPokemon.SpecialAttack} max={300} onChange={(e) => setSpecialAtk(e.target.value)}></input>
        <br/>
        <br/>
        <label>SpecialDef:</label>
        <input type="number" value={currentlySelectedPokemon.SpecialDefense} max={300} onChange={(e) => setSpecialDef(e.target.value)}></input>
        <br/>
        <br/>
        <label>Speed:</label>
        <input type="number" value={currentlySelectedPokemon.Speed} max={300} onChange={(e) => setSpeed(e.target.value)}></input>
        <br/>
        <br/>
        <label>Level: </label>
        <input type="number" max={100} min={1} onChange={(e) => setLevel(e.target.value)}></input>
        <br/>
        <br/>
        <label>Move1:</label>
        <select name="moves1" defaultValue={"placeholder"} onChange={(e) => setMove1(e.target.value)}>
        <option value={"placeholder"} disabled>Select a move</option>
        {currentlySelectedPokemon.Moves.map((move) =>
            <option value={move.Name} key={move.Name}>{move.Name}</option> 
        )}
        </select>
        <br/>
        <br/>
        <label>Move2:</label>
        <select name="moves2" defaultValue={"placeholder"} onChange={(e) => setMove2(e.target.value)}>
        <option value={"placeholder"} disabled>Select a move</option>
        {currentlySelectedPokemon.Moves.map((move) =>
            <option value={move.Name} key={move.Name}>{move.Name}</option> 
        )}
        </select>
        <br/>
        <br/>
        <label>Move3:</label>
        <select name="moves3" defaultValue={"placeholder"} onChange={(e) => setMove3(e.target.value)}>
        <option value={"placeholder"} disabled>Select a move</option>
        {currentlySelectedPokemon.Moves.map((move) =>
            <option value={move.Name} key={move.Name}>{move.Name}</option> 
        )}
        </select>
        <br/>
        <br/>
        <label>Move4:</label>
        <select name="moves4" defaultValue={"placeholder"} onChange={(e) => setMove4(e.target.value)}>
        <option value={"placeholder"} disabled>Select a move</option>
        {currentlySelectedPokemon.Moves.map((move) =>
            <option value={move.Name} key={move.Name}>{move.Name}</option> 
        )}
        </select>
        <br/>
        <br/>
        <label>HeldItem:</label>
        <select name="heldItem" defaultValue={"placeholder"} onChange={(e) => setHeldItem(e.target.value)}>
        <option value={"placeholder"} disabled>Select held Item</option>
        {heldItemsList.map((item) =>
            <option value={item.Name} key={item.Name}>{item.Name}</option> 
        )}
        </select>
        </>
    )
}
export default PokemonStats