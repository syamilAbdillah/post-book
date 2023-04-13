
export const formatDate = (unixEpochMili) => new Intl.DateTimeFormat("en", {timeStyle: "medium", dateStyle: "short"}).format(unixEpochMili)
export const delay = ms => new Promise(res => {
	setTimeout(res, ms)
})

export const formAction = fn => ev => {
	ev.preventDefault()
	const data = Object.fromEntries(new FormData(ev.currentTarget))
	fn(data, ev)
}