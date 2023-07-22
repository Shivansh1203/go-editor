import axios from 'axios';
import React, { useState } from 'react'

import {useNavigate} from 'react-router-dom'
import Navv from "./Navv";
const Login = () => {


  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault();
  
    try {
      const data = {
        username: username,
        password: password,
      };
    const response = await axios.post("http://localhost:8000/login/", data);
      localStorage.setItem("userData", JSON.stringify(response.data));
      localStorage.setItem("task", "threeSum");
      alert("Login completed");
      navigate("/home");
    } catch (error) {
      if (error.response) {
        const { message } = error.response.data;
        alert(message);
      } else {
        console.error(error);
      }
    }
  };
  return (
    <>
<Navv />
   <div className="auth-container d-flex align-items-center justify-content-center" style={{ minHeight: '100vh' }}>
  <form onSubmit={handleSubmit} className="p-4 border rounded">
    <h2>Login</h2>
    <div className="form-group">
      <label htmlFor="username">Username:</label>
      <input
        type="text"
        id="username"
        className="form-control"
        value={username}
        onChange={(event) => setUsername(event.target.value)}
      />
    </div>
    <div className="form-group">
      <label htmlFor="password">Password:</label>
      <input
        type="password"
        id="password"
        className="form-control"
        value={password}
        onChange={(event) => setPassword(event.target.value)}
      />
    </div>
    <button type="submit" className="btn btn-primary">Login</button>
  </form>
</div>

    </>
  );
};


export default Login