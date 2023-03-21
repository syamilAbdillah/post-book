import { createContext, useContext, useState, useEffect } from 'react'
import { delay } from '../components/utils'

const AuthContext = createContext({
	user: undefined,
})

export const useAuth = () => useContext(AuthContext)

export function AuthProvider({ children }) {
	const [user, setUser] = useState()
	const [loading, setLoading] = useState(true)
	const authenticating = async () => {
		await delay(2000)

		if(Math.random() < 0.5) {
			return {id: 'somerandomid'}
		}

		return undefined
	}

	useEffect(() => {
		authenticating()
			.then(setUser)
			.catch(error => console.log(error))
			.finally(() => setLoading(false))
	}, [])

	return <AuthContext.Provider value={{ user }}>
		{loading ? <Loading/>: children}
	</AuthContext.Provider>
}

const Loading = () => <div className="min-h-screen grid place-content-center bg-base-200">
	<button className="btn btn-ghost loading" disabled>
		<i>loading...</i>
	</button>
</div>