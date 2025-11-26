import { useState } from "react";
import { Route, Routes } from "react-router-dom";
import { useEffect } from "react";
import axios from "axios";
import EditCustomer from "./components/EditCustomer";
import InfoOfOneCustomer from "./components/InfoOfOneCustomer";
import Create from "./components/CreateCustomer";
import CustomerCard from "./components/CustomerCard";

function App() {
  <Routes>
    <Route path="/edit-customer/:id" element={<EditCustomer />} />
    <Route path="/single-customer/:id" element={<InfoOfOneCustomer />} />
    <Route path="/edit-customer" element={<Create />} />
  </Routes>;

  const [users, setUsers] = useState([]);
  //const [id, setId] = useState("");

  const API_URL = "http://localhost:8080/users";

  const fetchAllUsers = async () => {
    try {
      const res = await axios.get(API_URL);
      setUsers(res.data);
      console.log(res);
    } catch (e) {
      console.log(e);
    }
  };

  useEffect(() => {
    fetchAllUsers();
  }, []);

  return (
    <>
      {users.map((item, ind) => (
        <CustomerCard />
      ))}
    </>
  );
}

export default App;
