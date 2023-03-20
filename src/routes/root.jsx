import { Outlet } from 'react-router-dom'
import Sidebar from './components/Sidebar.jsx'

export default function Root() {
	return <>
		<Sidebar/>
		<main className="main-content">
			<Outlet/>
		</main>
	</>
}


