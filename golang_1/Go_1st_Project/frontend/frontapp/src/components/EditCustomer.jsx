import axios from 'axios'
import React, { use, useEffect, useState } from 'react'
import { useParams, useNavigate } from "react-router-dom";

const EditCustomer = () => {

  const {id}=useParams()
  const navigate=useNavigate()
  const [custName,setCustName]=useState("")
  const [custEmail,setCustEmail]=useState("")


useEffect(()=>{
  const fetchCustomer=async()=>{
    try {
      const res=await axios.get("http://localhost:8080/customer/${id}")
      setCustName(res.name)
      setCustEmail(res.email)
    } catch (e) {
      console.log(e)
    }
  }
  fetchCustomer()
},[id])



  const handleSubmit=async(e)=>{
    e.preventDefault()
 const res = await axios.put(`http://localhost:8080/customer/${id}`,{name:custName,email:custEmail})
 .then(()=>navigate("/customers"))


  }

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="enter name"
          onChange={(e) => setCustName(e.target.value)}
          value={custName}
        />
        <input
          type="text"
          placeholder="enter email"
          onChange={(e) => setCustEmail(e.target.value)}
          value={custEmail}
        />
        <input type="submit" value="Edit" />
      </form>
    </div>
  );
}

export default EditCustomer