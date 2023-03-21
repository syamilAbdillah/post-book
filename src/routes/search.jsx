const users = [
    {
        id: Math.floor(Math.random() * 10e9),
        username: 'John Doe',
        avatar: 'https://avatars.dicebear.com/api/open-peeps/default-user.svg',
    },
    {
        id: Math.floor(Math.random() * 10e9),
        username: 'Carl Johnson',
        avatar: 'https://avatars.dicebear.com/api/open-peeps/default-user.svg',
    },
    {
        id: Math.floor(Math.random() * 10e9),
        username: 'Sweat',
        avatar: 'https://avatars.dicebear.com/api/open-peeps/default-user.svg',
    },
]

import Section from './components/Section.jsx'
import TextInput from './components/TextInput.jsx'
import Avatar from './components/Avatar.jsx'
import { Link } from 'react-router-dom'

export default function Search() {
	return <>
        <Section>
            <TextInput type="search" placeholder="search user ..." />
        </Section>
        <Section>
            <div className="flex flex-col gap-4">
                {users.map(user => (
                    <UserCard user={user} key={user.id} />
                ))}
            </div>
        </Section>
    </>
}

const UserCard = ({user}) => (
    <Link to="#" className="flex items-center gap-4 p-2 hover:bg-base-200 rounded-lg ">
        <Avatar 
            size="sm"
            src={user.avatar} 
            alt={`${user.username} profile picture`} 
        />
        <h2 className="text-xl">{user.username}</h2>
    </Link>
)