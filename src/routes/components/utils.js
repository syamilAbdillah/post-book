
export const formatDate = (unixEpochMili) => new Intl.DateTimeFormat("en", {timeStyle: "medium", dateStyle: "short"}).format(unixEpochMili)
export const delay = ms => new Promise(res => {
	setTimeout(res, ms)
})