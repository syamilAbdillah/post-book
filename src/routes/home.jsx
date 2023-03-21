const posts = [
	{
		id: Math.floor((Math.random() * 10e9)),
		content: 'Lorem, ipsum dolor sit amet consectetur adipisicing elit. Reprehenderit quisquam error suscipit, quam. Temporibus perspiciatis, doloribus quaerat placeat. Nesciunt, reprehenderit!',
		created_at: Date.now() - 4000,
		updated_at: 0,
		user: {
			id: Math.floor((Math.random() * 10e9)),
			username: 'Matt flock',
			avatar: 'https://avatars.dicebear.com/api/open-peeps/default-user.svg',
		}
	},
	{
		id: Math.floor((Math.random() * 10e9)),
		content: 'Lorem, ipsum dolor sit amet consectetur adipisicing elit. Reprehenderit quisquam error suscipit, quam. Temporibus perspiciatis, doloribus quaerat placeat. Nesciunt, reprehenderit! amet consectetur adipisicing elit. Reprehenderit quisquam error suscipit, quam. Temporibus perspiciatis, doloribus quaerat placeat. Nesciunt, reprehenderit!',
		created_at: Date.now() - 3000,
		updated_at: 0,
		user: {
			id: Math.floor((Math.random() * 10e9)),
			username: 'Matt Demon',
			avatar: 'https://avatars.dicebear.com/api/open-peeps/default-user.svg',
		}
	},
	{
		id: Math.floor((Math.random() * 10e9)),
		content: 'Lorem, ipsum dolor sit amet consectetur adipisicing elit. Reprehenderit',
		created_at: Date.now() - 2000,
		updated_at: 0,
		user: {
			id: Math.floor((Math.random() * 10e9)),
			username: 'John Doe',
			avatar: 'https://avatars.dicebear.com/api/open-peeps/default-user.svg',
		}
	},
]


import {Form} from 'react-router-dom'
import Section from './components/Section.jsx'
import Avatar from './components/Avatar.jsx'
import Textarea from './components/Textarea.jsx'

export async function action({ request }) {
	const formData = await request.formData()
	console.log(Object.fromEntries(formData))
	return null
}

export default function Home() {
	return <div className="space-y-8">
		<PostInput/>
		{posts.map(post => (
			<DisplayPost post={post} key={post.id}/>
		))}
	</div>
}

function PostInput() {
	return <Section>
		<Form method="post" className="grid gap-6">
			<Textarea label="create new post"/>
			<div className="flex justify-end items-center">
				<button className="btn" type="submit">post</button>
			</div>
		</Form>
	</Section>
}

function DisplayPost({ post }) {
	return <Section>
		<div className="flex gap-4">
			<Avatar src={post.user.avatar} alt={post.user.username} />
			<div className="space-y-6">
				<hgroup>
					<h5 className="font-semibold text-xl">{post.user.username}</h5>
					<p className="text-gray-500">{formatDate(post.created_at)}</p>
				</hgroup>
				<p>{post.content}</p>
				<div className="grid grid-cols-2">
					<button className="btn btn-ghost">like</button>
					<button className="btn btn-ghost">comments</button>
				</div>
			</div>
		</div>
	</Section>
}

const formatDate = (unixEpochMili) => new Intl.DateTimeFormat("en", {timeStyle: "medium", dateStyle: "short"}).format(unixEpochMili)