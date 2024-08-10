import { ChangeEvent, useEffect, useState } from 'react';
import './App.css';
import { AddTask, LoadTasks, DeleteTask, CompleteTask} from "../wailsjs/go/main/App";

import { main } from "../wailsjs/go/models";

function App() {
    const [tasks, setTasks] = useState<main.Task[]>([]);
    const [newTaskName, setNewTaskName] = useState<string>("");


    // Загружаем задачи при запуске приложения
    useEffect(() => {
        LoadTasks().then((data: main.Task[]) => {
            setTasks(data);
        });
    }, []);

    // Функция для добавления задачи
    function handleAddTask() {
        if (newTaskName.trim() === "") {
            return;
        }
        if (newTaskName) {
            AddTask(newTaskName).then(() => {
                // После добавления задачи, перезагружаем список задач
                LoadTasks().then((data: main.Task[]) => {
                    setTasks(data);
                    setNewTaskName(""); // Очищаем поле ввода
                });
            });
        } else {
            alert("Please enter a task name");
        }
    }

    // Функция для удаления задачи
    function handleDeleteTask(taskName: string) {
        // Реализуй соответствующую функцию удаления в Go и вызови ее здесь
        DeleteTask(taskName).then(() => {
            LoadTasks().then((data: main.Task[]) => {
                setTasks(data);
            });
        });
    }

    // Функция для отметки задачи как выполненной
    function handleCompleteTask(taskName: string) {
        CompleteTask(taskName).then(() => {
            LoadTasks().then((data: main.Task[]) => {
                setTasks(data);
            });
        });
    }

    // Обработка изменений в поле ввода задачи
    function handleNewTaskNameChange(event: ChangeEvent<HTMLInputElement>) {
        setNewTaskName(event.target.value);
    }

    return (
        <div>
            <h1>Todo List</h1>

            <input
                type="text"
                value={newTaskName}
                onChange={handleNewTaskNameChange}
                placeholder="Enter task name"
            />
            <button onClick={handleAddTask}>Add Task</button>

            <ul>
                {tasks.map((task, index) => (
                    <li key={index}>
                        <span style={{ textDecoration: task.status === "completed" ? "line-through" : "none" }}>
                            {task.name} - {task.status}
                        </span>
                        <button onClick={() => handleCompleteTask(task.name)}>Complete</button>
                        <button onClick={() => handleDeleteTask(task.name)}>Delete</button>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default App;
