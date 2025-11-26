import React from 'react'
import { useParams } from 'react-router-dom';

const InfoOfOneCustomer = () => {

const {id}=useParams()

  useEffect(() => {
    const fetchCustomer = async () => {
      try {
        const res = await fetch(`http://localhost:5000/customers/${id}`);
        const data = await res.json();
console.log(res)
        
      } catch (error) {
        console.error("Error fetching customer:", error);
      }
    };

    fetchCustomer();
  }, [id]);
  return (
    <div>
      {data.map((item, ind) => (
        <div key={ind}>
          <h1>{item.name}</h1>
          <h3>{item.email}</h3>
        </div>
      ))}
    </div>
  );
}

export default InfoOfOneCustomer