import React from "react";
import { useNavigate } from "react-router-dom";
import "./signup.css";

function UpdateUser() {
  const regexes = {
    name: /^[a-zA-Z ]+$/,
    email: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
    group_name: /^[a-zA-Z ]+$/,
  };

  // Making form input as controlled components by managing them in states
  const [formData, setFormData] = React.useState({
    name: "",
    email: "",
    password: "",
    group_name: "",
    permission: false,
  });

  // To link form data with state formData
  function updateFormData(event) {
    console.log(formData);
    if (event.target.name == "permission") {
      setFormData((prevFormData) => {
        return { ...prevFormData, [event.target.name]: event.target.checked };
      });
    } else {
      setFormData((prevFormData) => {
        return { ...prevFormData, [event.target.name]: event.target.value };
      });
    }
  }

  function toString(obj) {
    const keys = Object.keys(obj);
    const values = Object.values(obj);

    const stringifiedObject = keys.map((key, index) => {
      if (typeof values[index] === "string") {
        return `${key}: "${values[index]}"`;
      } else {
        return `${key}: ${values[index]}`;
      }
    });

    return stringifiedObject.join("\n");
  }

  // on form submit
  async function submitData(event) {
    event.preventDefault();
    event.stopPropagation();

    const errors = {};
    for (const key in regexes) {
      if (!regexes[key].test(formData[key])) {
        errors[key] = "Invalid input";
      }
    }

    console.log(JSON.stringify(errors));
    if (Object.keys(errors).length > 0) {
      alert(toString(errors));
    } else {
      // Submit Logic
      // let result = await fetch("http://localhost:5000/signup", {
      //   method: "post",
      //   body: JSON.stringify(formData),
      //   headers: {
      //     "Content-type": "application/json",
      //   },
      // });
      // result = await result.json();
      // store user info in local storage
      // localStorage.setItem("user", JSON.stringify(result.result));
      // // store auth token in local storage
      // localStorage.setItem("token", JSON.stringify(result.auth));
      // console.log(result);
    }
  }

  return (
    <div class="signup-container">
      <div class="signup">
        <h1>Update User info</h1>
        <form action="POST">
          <div class="form-ele">
            <label htmlFor="name">Name </label>
            <input
              type="text"
              placeholder="Enter name"
              onChange={updateFormData}
              name="name"
              value={formData.name}
            />
          </div>
          <div class="form-ele">
            <label htmlFor="group_name">Group Name </label>
            <input
              type="text"
              placeholder="Enter group name"
              onChange={updateFormData}
              name="group_name"
              value={formData.group_name}
            />
          </div>
          <div class="form-ele">
            <label htmlFor="permission">Permission </label>
            <input
              type="checkbox"
              name="permission"
              checked={formData.permission}
              onClick={updateFormData}
            />
          </div>
          <button type="submit" onClick={submitData} className="signup-btn">
            Submit
          </button>
        </form>
      </div>
    </div>
  );
}

export default UpdateUser;
