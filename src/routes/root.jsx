import { useEffect } from 'react'
import { Outlet, useLoaderData, useNavigate } from 'react-router-dom'
import Sidebar from './components/Sidebar.jsx'
import { delay, formatDate } from './components/utils'
import { useAuth } from './context/authContext.jsx'

export async function loader() {

	return null
}

export default function Root() {
	return <>
		<Sidebar/>
		<Main>
			<Outlet/>
		</Main>
	</>
}

const Main = ({ children, ...rest }) => <main className="min-h-screen ml-20">
	{children}
</main>


