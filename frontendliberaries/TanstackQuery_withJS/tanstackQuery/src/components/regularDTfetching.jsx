import React, { useEffect, useState } from "react";
import axios from "axios";
import PostCard from "./PostCard";

const RegularDTfetching = () => {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(false);

  const fetchPosts = async () => {
    try {
      setLoading(true);
      setError(false);

      const res = await axios.get("http://localhost:3000/posts");
      setData(res.data);
    } catch (err) {
      console.error("Error fetching posts:", err);
      setError(true);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchPosts();
  }, []);

  return (
    <div>
      {loading && <h2>Loading...</h2>}
      {error && <h2>Something went wrong</h2>}

      {data.map((post) => (
        <PostCard key={post.id} post={post} />
      ))}
    </div>
  );
};

export default RegularDTfetching;
