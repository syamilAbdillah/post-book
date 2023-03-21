import { Outlet } from 'react-router-dom'
import Sidebar from './components/Sidebar.jsx'

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


