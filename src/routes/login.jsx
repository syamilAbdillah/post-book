import { useState, useEffect } from 'react'
import { Form, useNavigate } from 'react-router-dom'
import TextInput from './components/TextInput.jsx'
import Toggle from './components/Toggle.jsx'
import { useAuth } from './context/authContext.jsx'

export async function action({ request }) {
	const formData = await request.formData()
	console.log(Object.fromEntries(formData))
	return null
}

export default function Login() {
	const [show, setShow] = useState(false)
	const handleSubmit = ev => {
		ev.preventDefault()
	}

	return <PageLayout>
		<FormCard>
			<h2 className="text-3xl font-black text-center">login</h2>
			<form className="grid gap-6" onSubmit={handleSubmit}>
				<TextInput
					name="username"
					placeholder="username"
					label="username"
					type="text"
				/>
				<TextInput
					name="password"
					placeholder="password"
					label="password"
					type={show ? "text":"password"}
				/>
				<Toggle value={show} onChange={ev => setShow(ev.currentTarget.checked)} label="show password" />
				<div className="flex justify-end gap-2">
					<button type="submit" className="btn btn-primary">login</button>
				</div>
			</form>
		</FormCard>
	</PageLayout>
}

const FormCard = ({ children }) => <article className="w-full md:max-w-xl px-6 py-12 bg-base-100 border border-base-300 shadow rounded-lg">{children}</article>

const PageLayout = ({ children }) => <div className="min-h-screen bg-base-200 flex justify-center items-center p-6">{children}</div>

