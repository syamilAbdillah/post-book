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

import Section from './components/Section.jsx'
import Avatar from './components/Avatar.jsx'
import DisplayPost from './components/DisplayPost.jsx'

export default function Profile() {
	return <div className="space-y-8">
		<DisplayProfile/>
		{posts.map(post => (
			<DisplayPost post={post} key={post.id} />
		))}
	</div>
}

const DisplayProfile = () => (
	<Section>
		<div className="flex flex-col items-center md:items-start md:flex-row md:justify-between gap-4 pt-12">
			<div className="flex flex-col items-center md:items-start gap-4">
				<Avatar 
					src="https://avatars.dicebear.com/api/open-peeps/default-user.svg" 
					alt=""
				/>
				<hgroup>
					<h2 className="text-xl">some username</h2>
					<p className="text-gray-500" >Lorem ipsum dolor sit</p>
				</hgroup>
			</div>
			<div className="grid grid-cols-3 gap-4">
				<ProfileStat value="10k" label="post" />
				<ProfileStat value="3m" label="followers" />
				<ProfileStat value="10" label="following" />
				<div className="col-span-3 flex justify-center md:justify-end">
					<button className="btn">follow</button>
				</div>
			</div>
		</div>
	</Section>
)

const ProfileStat = ({ value, label }) => <hgroup className="p-2 flex flex-col items-center">
	<h2 className="text-xl">{value}</h2>
	<p className="text-gray-500" >{label}</p>
</hgroup>