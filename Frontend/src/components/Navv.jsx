import React from 'react'
import logo from "../assets/logo.jpeg";
import '../App.css';
import { useNavigate } from "react-router-dom";
const Navv = () => {
    const navigate = useNavigate();
    const onLogin = async (event) => {
        event.preventDefault();


        navigate("/login");

    };
    const onSignup = async (event) => {
        event.preventDefault();
        navigate("/signup");

    };
  return (
    <>
            <nav>
            <div className="main-nav container flex">
                <a href="/" className="company-logo">
                    <img src={logo} alt="company logo" />
                </a>
                <div className="search-bar flex">

                    <button
                        id="signup-button"
                        className="btn btn-primary mx-2"
                        onClick={onSignup}
                    >
                        Sign Up
                    </button>
                    <button
                        id="login-button"
                        className="btn btn-outline-primary"
                        onClick={onLogin}
                    >
                        Login
                    </button>

                </div>
            </div>
        </nav>
    </>
  )
}

export default Navv