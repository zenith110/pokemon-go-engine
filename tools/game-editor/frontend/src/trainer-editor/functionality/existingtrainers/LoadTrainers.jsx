import { useState, useEffect } from "react"

import { ParseTrainers } from "../../../../wailsjs/go/main/App"
import TrainerComboBox from "./TrainerComboBox"
const LoadTrainers = () => {
    const [trainers, setTrainers] = useState([])
    useEffect(() => {
        const FetchTrainers = async() =>{
            let data = await ParseTrainers()
            console.log(data.Trainers.length)
            setTrainers(trainers => [...trainers, data])
        }
        FetchTrainers()
    }, [])
    return(
        <>
        <TrainerComboBox trainers={trainers.Trainers}/>
        </>
    )
}
export default LoadTrainers