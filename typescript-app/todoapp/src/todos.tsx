
import React, { useState } from 'react'
import { Button } from "@/components/ui/button";
import axios from 'axios';


const Todos = () => {
  const [task,setTask] = useState("")

    const addTask=()=>{
        return 
        axios.post(`http://localhost:3001/todos`,task)
        .then(()=>console.log("added in db"))
        .catch((e)=>console.log(e))
    } 

  return (
    <div className='flex-col justify-center '>
<input className='outline-1 w-[50%]' type="text" placeholder='enter todo' value={task} onChange={(e)=>setTask(e.target.value)}/>
<Button type="submit" onClick={addTask}>add</Button>

    </div>
  )
}

export default Todos
