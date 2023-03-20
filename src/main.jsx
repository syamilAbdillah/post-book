import React, { useState } from 'react'
import ReactDOM from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'

import './index.css'

// routes
import Root from './routes/root.jsx'
import Login from './routes/login.jsx'
import Friends from './routes/friends.jsx'
import Search from './routes/search.jsx'
import Home from './routes/home.jsx'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Root/>,
    children: [
      {
        path: 'friends',
        element: <Friends/>, 
      },
      {
        path: 'search',
        element: <Search/>,
      },
      {
        index: true,
        element: <Home/>
      }
    ]
  },
  {
    path: '/login',
    element: <Login/>,
  }
])

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
