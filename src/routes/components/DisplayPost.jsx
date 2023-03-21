import Section from './Section.jsx'
import Avatar from './Avatar'
import { formatDate } from './utils'

export default function DisplayPost({ post }) {
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