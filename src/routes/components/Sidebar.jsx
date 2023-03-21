import { IconSearch, IconUsers, IconHome2 } from '@tabler/icons-react'
import { NavLink } from 'react-router-dom'
import Avatar from './Avatar.jsx'

export default function Sidebar() {
	return (
		<Nav>
			<SidebarLink to="/" >
				<IconHome2 size={32} />
			</SidebarLink>
			<SidebarLink to="friends">
				<IconSearch size={32}/>
			</SidebarLink>
			<SidebarLink to="search">
				<IconUsers size={32} />
			</SidebarLink>
			<SidebarLink to="profile">
				<Avatar size="sm" src="https://avatars.dicebear.com/api/open-peeps/default-user.svg" alt="default-profile"/>
			</SidebarLink>
		</Nav>	
	)
}

function SidebarLink({children, ...rest}) {
	const base = 'flex justify-center items-center w-16 h-16  rounded-lg transition hover:bg-base-200'
	const active = `${base} text-primary bg-base-200`
	return <NavLink 
		className={({ isActive }) => isActive ? active: base}
		{...rest} 
	>
		{children}
	</NavLink>
}

const Nav = ({ children, ...rest }) => <nav className="fixed inset-y-0 left-0 z-50 flex flex-col items-center gap-2 py-8 w-20 border-r bg-base-100">
	{children}
</nav>