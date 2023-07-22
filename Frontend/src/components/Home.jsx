import React, { useState, useEffect } from "react";
import Navbar from "./Navbar";
import NewsCard from "./Card";

const url = "https://api.newscatcherapi.com/v2/search?q=";

const App = () => {
  // const [articless, setArticles] = useState([]);
  // const [selectedNav, setSelectedNav] = useState(null);

  // useEffect(() => {
  //   fetchNews("India");
  // }, []);

  // const fetchNews = async (query) => {
  //   const res = await fetch(`${url}${query}`, {
  //     method: 'GET', 
  //     headers: {
  //       'Content-Type': 'application/json',
  //       'x-api-key': 'W2vWgxyuQTXYlx9vWB_c-_JE9t7ZMYl36fsop44bnxQ'
  //     }
  //   });

    
  //   const data = await res.json();
  //   if (data.status === 'ok') {
  //     bindData(data.articles);
  //   } else {
  //     console.error('Failed to fetch news:', data.message);
  //   }
  // };
   
  // const bindData=(articles) =>{
  //   setArticles(articles);
  // }
  // const reload = () => {
  //   window.location.reload();
  // };

  const openArticle = (url) => {
    window.open(url, "_blank");
  };

  // const onNavItemClick = (id) => {
  //   fetchNews(id);
  //   setSelectedNav(id);
  // };

  // const onSearchButtonClick = (query) => {
  //   if (!query) return;
  //   fetchNews(query);
  //   setSelectedNav(null);
  // };

  return (
    <>
      <Navbar/>
      <main>
        <div className="cards-container container flex">

  <NewsCard
    article="CodeEditor"
    openLink="http://localhost:3000/tasklist"
    openArticle={openArticle}
  />
  <NewsCard
    article="CodeEditor"
    openLink="http://localhost:3000/tasklist"
    openArticle={openArticle}
  />
        </div>
      </main>
    </>
  );
};

export default App;