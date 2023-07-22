import React, { useState } from "react";
import defaultLogo from "../assets/logo.jpeg";

const NewsCard = ({ article,openLink, openArticle }) => {
  

 

  return (
    <div className="card" onClick={() => openArticle(openLink)}>
<div className="card-header"><img src={defaultLogo} alt="news-image" id="news-img" /></div>
      <div className="card-content">
        <h3 id="news-title">{article}</h3>
        <h6 className="news-source" id="news-source">
          CPP Code Editor
        </h6>
        <p className="news-desc" id="news-desc">
        preferable by most
        </p>
        <p className="news-source">SHIV App</p>
      </div>
    </div>
  );
};

export default NewsCard;