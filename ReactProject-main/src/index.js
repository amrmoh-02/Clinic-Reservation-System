import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';


import Home from './pages/home';
import Doctor from './pages/doctor';
import Patient from './pages/patient';



import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home/>,
    errorElement: <h1>erorr..................</h1> ,
  },
  {
    path: "/doctor",
    element: <Doctor/>,
  },
  {
    path: "/patient",
    element: <Patient/>,
  },
]);





const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
     <RouterProvider router={router} />
  </React.StrictMode>
);
