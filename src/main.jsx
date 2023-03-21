import React, { useState } from 'react'
import ReactDOM from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'

import './index.css'

import { AuthProvider, useAuth } from './routes/context/authContext.jsx'

// routes
import Root, { loader as rootLoader } from './routes/root.jsx'
import Login, { action as loginAction } from './routes/login.jsx'
import Friends from './routes/friends.jsx'
import Search from './routes/search.jsx'
import Home, {action as homeAction} from './routes/home.jsx'
import Profile from './routes/profile.jsx'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Root/>,
    loader: rootLoader,
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

function App() {
  const {user} = useAuth()

  if(!user) {
    return <Login/>
  }

  return <RouterProvider router={router} />
}

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <AuthProvider>
      <App/>
    </AuthProvider>
  </React.StrictMode>,
)
