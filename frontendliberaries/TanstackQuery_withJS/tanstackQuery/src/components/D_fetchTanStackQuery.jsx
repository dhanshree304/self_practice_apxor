import { useQuery } from '@tanstack/react-query'
import axios from 'axios'
import React from 'react'

const D_fetchTanStackQuery = () => {

const {data,isLoading,isError,error}=useQuery({
    queryKey:["posts"],
    queryFn:()=>{
        return axios.get("http://localhost:3000/posts")
    }
})

if(isLoading){
    return <div>Page is loading...</div>
}

if(isError){
    return <div>{error.message}</div>
}

  return (
    <div
      style={{
        border: "5px solid gray",
        margin: "5px",
        padding: "5px",
        color: "red",
      }}
    >
      {data?.data.map((post) => (
        <div
          key={post.id}
          style={{
            border: "2px solid gray",
            margin: "5px",
            padding: "5px",
            color: "red",
          }}
        
          ><h1>{post.title}</h1>
          <h3>{post.body}</h3>
        </div>
      ))}
    </div>
  );
}

export default D_fetchTanStackQuery