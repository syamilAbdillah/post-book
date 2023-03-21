

function Avatar({ size, className: _, ...rest }) {
	const base = 'border rounded-full bg-base-100'
	const sizes = {
		sm: `w-14 h-14`,
		md: `w-20 h-20`,
		lg: `w-36 h-36`,
	}

	const className = `${base} ${sizes[size] || sizes.md}`

	return <div className="avatar">
		<div className={className}>
			<img {...rest} />
		</div>
	</div>
}

export default Avatar