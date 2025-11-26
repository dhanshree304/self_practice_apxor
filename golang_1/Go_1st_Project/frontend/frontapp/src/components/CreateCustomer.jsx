import axios from "axios";
import React from "react";

const Create = () => {
  const [customerName, setCustomerName] = useState("");
  const [email, setEmail] = useState("");

  const fetchUsers = async () => {
    try {
      setLoading(true);
      const res = await axios.get("http://localhost:8080");
      setUsers(res.data);
      setLoading(false);
    } catch (err) {
      console.error(err);
      setLoading(false);
    }
  };

  
  
  const createCustomer = async () => {
    
    try {
      await axios.post(API_URL, {
        customerName,
        email,
      });
      setCustomerName("");
      setEmail("");
      fetchUsers()
      
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div>
      <input
        type="text"
        value={customerName}
        placeholder="enter name"
        onChange={(e) => setCustomerName(email.target.value)}
      />
      <input type="text" 
      value={email}
      placeholder="enter email"
      onChange={(e)=>setEmail(e.target.value)}
      />
      <button onClick={createCustomer}>Add Customer</button>
    </div>
  );
};

export default Create;
