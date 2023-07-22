import './App.css';
import React from "react";
import { Routes, Route } from "react-router-dom";
import Home from './components/Home';
import Start from './components/Start';
import Login from './components/Login';
import Signup from './components/Signup';
import UpdateUser from './components/UpdateUser';
import TextEditor from './components/TextEditor';
import TaskList from './components/TaskList';
import 'bootstrap/dist/css/bootstrap.min.css';


function App() {
  return (
    <>
      <Routes>
        <Route exact path="/" element={<Start />} />

         <Route path="home" element={<Home />} />
         <Route path="updatation" element={<UpdateUser />} />
         <Route path="texteditor" element={<TextEditor />} />
         <Route path="tasklist" element={<TaskList />} />
        <Route exact path='/login' element={<Login />} />
        <Route exact path='/Signup' element={<Signup />} />

      </Routes>
    </>
  );
}

export default App;
