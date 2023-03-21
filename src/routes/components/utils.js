
export const formatDate = (unixEpochMili) => new Intl.DateTimeFormat("en", {timeStyle: "medium", dateStyle: "short"}).format(unixEpochMili)