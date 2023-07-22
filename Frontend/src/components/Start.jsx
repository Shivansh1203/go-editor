import React from 'react'
import '../App.css';
import vdo from "../assets/video.mp4";
import Navv from "./Navv";
const Start = () => {
   
    return (
        <>
        <Navv />
        <div className='hero-container'>
      <video src={vdo} autoPlay loop muted />
      <h1>HRS News</h1>
      <p>We are here to tell you the truth</p>
      <p>To make you aware </p>

    </div>
    
        </>
    )
}

export default Start