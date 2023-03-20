import { IconSearch, IconUsers, IconHome2 } from '@tabler/icons-react'
import { NavLink } from 'react-router-dom'

export default function Sidebar() {
	return (
		<nav className="navbar">
			<SidebarLink to="/" >
				<IconHome2 size={32} />
			</SidebarLink>
			<SidebarLink to="friends">
				<IconSearch size={32}/>
			</SidebarLink>
			<SidebarLink to="search">
				<IconUsers size={32} />
			</SidebarLink>
			<footer>
				<img className="avatar" src="https://avatars.dicebear.com/api/open-peeps/default-user.svg" alt="default-profile"/>
			</footer>
		</nav>	
	)
}

function SidebarLink({children, ...rest}) {
	return <NavLink 
		className={({ isActive }) => isActive ? "active": ""}
		{...rest} 
	>
		{children}
	</NavLink>
}