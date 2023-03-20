import { useId } from 'react'

const TextInput = ({ label, error, ...rest }) => {
	const id = useId()

	return (<div className="form-control w-full">
		{label && (
			<label htmlFor={id} className="label">
				<span className="label-text">{label}</span>
			</label>
		)}
		<input {...rest} id={id} className="input input-bordered w-full" />
		{error && (
		  	<label htmlFor={id} className="label">
				<span className="label-text-alt text-error">{error}</span>
			</label>
		)}
	</div>)
}

export default TextInput