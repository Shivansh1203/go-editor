import React, { useState } from "react";
import axios from "axios";
import Navv from "./Navv";

const Signup = () => {
  const [name, setName] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [role, setRole] = useState("Admin");
  const [parent, setParent] = useState("Parent");
  const [jgroup, setJgroup] = useState(["group1","2222"]);
  const [group, setGroup] = useState("");
  const [permission, setPermission] = useState("true");
  const [request, setRequest] = useState("");
  const [editPermision, setEditPermision] = useState(true);

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const data = {
        name,
        username,
        password,
        role,
        parent,
        jgroup,
        permission,
        request,
        editPermision,
      };
      await axios.post("http://localhost:8000/register/", data);
      alert("Registration Completed! Now login.");
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <>
      <Navv />

      <div
        className="auth-container"
        style={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          height: "100vh",
        }}
      >
        <form onSubmit={handleSubmit}>
          <h2>Register</h2>
          <div className="form-group">
            <label htmlFor="name">Name:</label>
            <input
              type="text"
              id="name"
              className="form-control"
              value={name}
              onChange={(event) => setName(event.target.value)}
            />
          </div>
          <div className="form-group">
            <label htmlFor="username">Email:</label>
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
          <div className="form-group">
            <label htmlFor="phoneNumber">Group name:</label>
            <input
              type="text"
              id="phoneNumber"
              className="form-control"
              value={group}
              placeholder=""
              onChange={(event) => setGroup(event.target.value)}
            />
          </div>
          <div className="form-group">
          <label htmlFor="permission">Can everyone add you in their group </label>
            <input
              type="checkbox"
              name="permission"
              checked={permission}
              onChange={(event) => setPermission(event.target.checked)}
            />
          </div>
          <button type="submit" className="btn btn-primary">
            Register
          </button>
        </form>
      </div>
    </>
  );
};

export default Signup;
