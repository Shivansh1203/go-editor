import React from "react";
import "./tasklist.css";
import TaskItem from "./TaskItem";

const TaskList = () => {
  const openTask = (url,id) => {
    window.open(url, "_blank");
  };


  const documents = [
    {
      id: 1,
      name: "Document 1",
      description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Fugit consequuntur repellendus blanditiis excepturi alias repudiandae unde, laborum vel commodi! Et ducimus blanditiis dicta deleniti nesciunt, praesentium adipisci eum aliquid quae amet ipsam incidunt debitis nobis, animi molestias veniam nam labore? Quam voluptatum veritatis sint doloremque. Exercitationem vero amet veniam ipsum. Sit velit in necessitatibus mollitia repudiandae? Nostrum porro laudantium, provident dolores non nihil nesciunt quos. Praesentium tempora facilis tenetur voluptate",
      icon: "fa-solid fa-file-code",
      dateModified: "2023-03-08T12:34:56Z",
    },
    {
      id: 2,
      name: "Document 2",
      description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Fugit consequuntur repellendus blanditiis excepturi alias repudiandae unde, laborum vel commodi! Et ducimus blanditiis dicta deleniti nesciunt, praesentium adipisci eum aliquid quae amet ipsam incidunt debitis nobis, animi molestias veniam nam labore? Quam voluptatum veritatis sint doloremque. Exercitationem vero amet veniam ipsum. Sit velit in necessitatibus mollitia repudiandae? Nostrum porro laudantium, provident dolores non nihil nesciunt quos. Praesentium tempora facilis tenetur voluptate.Lorem ipsum dolor sit amet consectetur adipisicing elit. Fugit consequuntur repellendus blanditiis excepturi alias repudiandae unde, laborum vel commodi! Et ducimus blanditiis dicta deleniti nesciunt, praesentium adipisci eum aliquid quae amet ipsam incidunt debitis nobis, animi molestias veniam nam labore? Quam voluptatum veritatis sint doloremque. Exercitationem vero amet veniam ipsum. Sit velit in necessitatibus mollitia repudiandae? Nostrum porro laudantium, provident dolores non nihil nesciunt quos. Praesentium tempora facilis tenetur voluptate.",
      icon: "fa-solid fa-file-code",
      dateModified: "2023-03-09T12:34:56Z",
    },
    {
      id: 3,
      name: "Document3",
      description: "This is the third document.",
      icon: "fa-solid fa-file-code",
      dateModified: "2023-03-10T12:34:56Z",
    },
  ];
  const documentItems = documents.map((document) => {
    return (
      <>
        <TaskItem
          key={document.id}
          id={document.id}
          icon={document.icon}
          dname={document.name}
          description={document.description}
          date={document.dateModified}
          openTask={openTask}
        />
        <hr />
      </>
    );
  });

  return (
    <div className="task-list-container">
        <h1>Task List</h1>
      <div className="task-list">
        {documentItems}
      </div>
    </div>
  );
};

export default TaskList;
