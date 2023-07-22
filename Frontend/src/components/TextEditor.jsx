import React, { useEffect, useState } from "react";
import "./TextEditor.css";
import userEvent from "@testing-library/user-event";
import axios from "axios";
const TextInputContainer = () => {
  const [textValue, setTextValue] = useState("");
  const handleInputChange = (event) => {
    setTextValue(event.target.value);
  };

  const handleClearClick = () => {
    setTextValue("");
  };

  const userData = JSON.parse(localStorage.getItem("userData"));
  const [allow, setAllow] = useState(false);

  // Determine if the user is allowed to perform certain actions based on the role
  useEffect(() => {
    console.log("hello");
    if (userData.role === "watcher") {
      console.log("buh");
      setAllow(false);
    } else {
      console.log("huh");
      setAllow(true);
    }
    const taskData = JSON.parse(localStorage.getItem("ok"));
    setTextValue(taskData.jInput)
  }, []);

  const handleFetchClick = async(event) => {
   
    event.preventDefault();
    

    try {
      const data = {
        name: localStorage.getItem("task")
      };

      const response = await axios.post("http://localhost:8000/fetch/", data);
      localStorage.setItem("ok", JSON.stringify(response.data));
      alert("Fetching data: " + userData.role);
      window.location.reload()
    } catch (error) {
      if (error.response) {
        const { message } = error.response.data;
        alert(message);
      } else {
        console.error(error);
      }
    }
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const data = {
        name: localStorage.getItem("task"),
        input: textValue,
      };

      const response = await axios.post("http://localhost:8000/commit/", data);
      localStorage.setItem("userData", JSON.stringify(response.data));
      alert("committed");
    } catch (error) {
      if (error.response) {
        const { message } = error.response.data;
        alert(message);
      } else {
        console.error(error);
      }
    }
  };

  const handleSave = () => {
    // TODO: Implement the save functionality
  };

  return (
    <div className="text-editor-container">
      <textarea
        className="text-editor"
        value={textValue}
        onChange={handleInputChange}
        placeholder="Start typing here..."
      />
      <div className="button-container">
        {/* Disable Fetch button for watcher */}
        {allow ? (
          <>
            <button onClick={handleFetchClick}>Fetch</button>
            <button onClick={handleSubmit}>Commit</button>

            <button onClick={handleClearClick}>Clear</button>
          </>
        ) : (
          <>
            <button onClick={handleFetchClick} disabled>
              Fetch
            </button>
            <button onClick={handleSubmit} disabled>
              Commit
            </button>

            <button onClick={handleClearClick} disabled>
              Clear
            </button>
          </>
        )}
      </div>
    </div>
  );
};

export default TextInputContainer;
