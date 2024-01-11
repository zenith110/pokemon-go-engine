import { useState } from "react"
const TrainerComboBox = ({ trainers }) => {
    const [trainer, setTrainer] = useState([])
    return(
        <select name="trainers" onChange={(e) => console.log(e.target.value)} defaultValue={"placeholder"}>
            <option value={"placeholder"} disabled>Select a trainer</option>
            {trainers.map((trainer) =>
                <option value={trainer.Name} key={trainer.Name}>{trainer.Name}</option>
            )}
            </select>
    )
}
export default TrainerComboBox