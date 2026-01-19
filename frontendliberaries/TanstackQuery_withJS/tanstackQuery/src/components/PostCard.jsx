import React from 'react'

const PostCard = ({post}) => {
  return (
    <div style={{border:"5px solid gray",margin:"5px",padding:"5px",color:"white"}}>
<h2>Title: {post.title}</h2>
<p>Body: {post.body}</p>
    </div>
  )
}

export default PostCard