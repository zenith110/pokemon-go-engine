import { useState, useEffect } from "react"
import { CreateTrainerData, ParsePokemonData, ParseHeldItems} from "../../../../wailsjs/go/main/App"
import { v4 as uuidv4 } from 'uuid';
import TrainerPokemonsGenerator from "./TrainerPokemonsGenerator"
const NewTrainerCard = ({ setNewTrainer }) => {
    const [name, setName] = useState("")
    const [sprite, setSprite] = useState("")
    const [classTypes, setClassTypes] = useState([])
    const [classType, setClassType] = useState("")
    const [visibleClassTypesGroup, setVisibleClassTypesGroup] = useState(false)
    const [visibleSprites, setVisibleSprites] = useState(false)
    const [pokemonSpecies, setPokemonSpecies] = useState([])
    const [pokemonCount, setPokemonCount] = useState(0)
    const [pokemons, setPokemons] = useState([])
    const [heldItemsList, setHeldItemList] = useState([])
    useEffect(() => { 
        const fetchClassTypes = () => {

        }
        const fetchSprites = () => {

        }
        const fetchPokemonSpecies = async() => {
            let data = await ParsePokemonData()
            setPokemonSpecies(data)
        }
        const fetchHeldItemListData = async() => {
            let data = await ParseHeldItems()
            setHeldItemList(data)
        }
        fetchHeldItemListData()
        fetchClassTypes()
        fetchSprites()
        fetchPokemonSpecies()
    }, []) 
    const Submit = () => {
        let data = {
            "name": name,
            "sprite": "yo",
            "classType": "yo",
            "id": uuidv4(),
            "pokemons": pokemons
        }
        CreateTrainerData(data)
        setNewTrainer(false)
    }
    return(
        <>
        <label>Trainer name: </label>
        <input type="text" onChange={(event) => setName(event.target.value)}></input>
        <br/>
        <label>Total amount of pokemon: </label>
        <input type="number"  min="1" max="6" onChange={(event) => setPokemonCount(event.target.value)}/>
        <br/>
        {
            pokemonSpecies.length >= 1 && heldItemsList.length >= 1? <TrainerPokemonsGenerator setPokemons={setPokemons} pokemonSpeciesList={pokemonSpecies} heldItemsList={heldItemsList} pokemonsCount={pokemonCount}/> : <></>
        }
        <button onClick={() => Submit()}>Submit</button>
        <button onClick={() => setNewTrainer(false)}>Close</button>
        </>
    )
}
export default NewTrainerCard