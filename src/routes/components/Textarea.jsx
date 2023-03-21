export default function Textarea({ className: _, error, label, value, ...rest }) {
	return <div className="form-control">
		{label && (
			<label className="label">
				<span className="label-text">{label}</span>
			</label>
		)}
		<textarea className="textarea textarea-bordered h-24" {...rest}>{value || ''}</textarea>
		{error && (
			<label className="label">
				<span className="label-text-alt text-error">{error}</span>
			</label>
		)}
	</div>
}