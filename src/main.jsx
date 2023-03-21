import React, { useState } from 'react'
import ReactDOM from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'

import './index.css'

// routes
import Root from './routes/root.jsx'
import Login, { action as loginAction } from './routes/login.jsx'
import Friends from './routes/friends.jsx'
import Search from './routes/search.jsx'
import Home, {action as homeAction} from './routes/home.jsx'
import Profile from './routes/profile.jsx'

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
        path: 'profile',
        element: <Profile/>,
      },
      {
        index: true,
        element: <Home/>,
        action: homeAction,
      }
    ]
  },
  {
    path: '/login',
    element: <Login/>,
    action: loginAction,
  }
])

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
