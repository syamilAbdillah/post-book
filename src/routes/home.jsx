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

export default function Home() {
	return <>
		<PostInput/>
		{posts.map(post => (
			<DisplayPost post={post} key={post.id}/>
		))}
	</>
}

function PostInput() {
	return <section>
		<form>
			<label>
				<textarea name="content" rows="4"></textarea>
			</label>
			<input type="submit"/>
		</form>
	</section>
}

function DisplayPost({ post }) {
	return <section>
		<div className="post">
			<img src={post.user.avatar} alt={post.user.username} className="avatar"/>
			<div>
				<hgroup>
					<h5>{post.user.username}</h5>
					<p>{formatDate(post.created_at)}</p>
				</hgroup>
				<p>{post.content}</p>
				<div className="post-action">
					<button className="outline">like</button>
					<button className="outline">comments</button>
				</div>
			</div>
		</div>
	</section>
}

const formatDate = (unixEpochMili) => new Intl.DateTimeFormat("en", {timeStyle: "medium", dateStyle: "short"}).format(unixEpochMili)