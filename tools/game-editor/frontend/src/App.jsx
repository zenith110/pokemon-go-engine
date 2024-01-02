
import NewTrainerCard from "./trainer-editor/functionality/newtrainers/NewTrainerCard"
import { useState } from "react"
function App() {
    const [newTrainer, setNewTrainer] = useState(false)
    const [editTrainers, setEditTrainers] = useState(false)
    return (
        <>
           <button onClick={() => setNewTrainer(true)}>Create new trainer</button>
           <button>Edit trainers</button>
           <br/>
           {
            newTrainer ? <NewTrainerCard setNewTrainer={setNewTrainer}/> : <></>
           }
        </>
    )
}

export default App
