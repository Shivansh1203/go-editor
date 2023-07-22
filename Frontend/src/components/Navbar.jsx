import React, { useState, useEffect } from "react";
import logo from "../assets/logo.jpeg";
import { useNavigate, Link  } from 'react-router-dom'

const Navbar = () => {
  const navigate = useNavigate();
  async function signOut(event) {
    event.preventDefault();
    event.stopPropagation();
    localStorage.clear();
    navigate("/");
  }
  let [userInfo, setuserInfo] = useState("");
  console.log(localStorage.getItem("userID"))
  useEffect(() => {
    if (localStorage.getItem("userID") != null) {
      setuserInfo(localStorage.getItem("userID"));
    }
  }, []);
  console.log(userInfo)

  return (
    <nav>
      <div className="main-nav container flex">
        <a href="/home" className="company-logo">
          <img src={logo } alt="company logo" />
        </a>

        <div style={{ display: "flex", justifyContent: "center" }}>
          <div style={{marginRight:"10px"}}> <Link to="/updatation"> Shivansh Rastogi
          
          </Link></div>
          <div>
            <button
              id="search-button"
              className="search-button"
              onClick={signOut}
            >
              Logout
            </button>
          </div>
        </div>

      </div>

    </nav>
  );
};

export default Navbar;
