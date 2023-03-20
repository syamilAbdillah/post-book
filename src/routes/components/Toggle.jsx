export default function Toggle({ label, type: _, className: __, ...rest }) {
	return <div className="form-control">
		<label className="label cursor-pointer inline-flex justify-start gap-2 items-center">
			<input type="checkbox" className="toggle toggle-primary" {...rest} />
			{label && (
				<span className="label-text">{label}</span> 
			)}
		</label>
	</div>
}