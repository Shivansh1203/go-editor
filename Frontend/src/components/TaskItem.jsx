import React from "react";
import "./tasklist.css";

const TaskItem = ({id,
  icon,
  dname,
  description,
  date,
  openTask}) => {
  return (
    <div className="task-item" onClick={()=> openTask("http://localhost:3000/texteditor",id)}>
      <div className="task-info">
        <div className="task-icon task-info-item">
          <i className={icon} />
        </div>
        <div className="task-info-item task-name">{dname}</div>
        <div className="task-info-item task-date">{date}</div>
      </div>
      <div className="task-desc">
        <div className="task-desc-item">{description}</div>
      </div>

    </div>
  );
};

export default TaskItem;
