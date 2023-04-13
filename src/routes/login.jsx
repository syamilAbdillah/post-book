import { useState, useEffect } from 'react'
import { Form, useNavigate } from 'react-router-dom'
import TextInput from './components/TextInput.jsx'
import Toggle from './components/Toggle.jsx'
import { useAuth } from './context/authContext.jsx'
import { formAction } from './components/utils'

function Login() {
	const [show, setShow] = useState(false)
	const handleSubmit = formAction(data => console.log(data))

	return <form className="grid gap-4" onSubmit={handleSubmit}>
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
		<Toggle 
			value={show} 
			onChange={
				ev => setShow(ev.currentTarget.checked)
			} 
			label="show password" 
		/>
		<div className="flex justify-end gap-2">
			<button type="submit" className="btn">login</button>
		</div>
	</form>
}

function Register() {
	const handleSubmit = formAction(data => console.log(data))
	const [show, setShow] = useState(false)

	return <form className="grid gap-4" onSubmit={handleSubmit}>
		<TextInput
			name="username"
			placeholder="username"
			label="username"
			type="text"
		/>
		<TextInput
			name="name"
			placeholder="fullname"
			label="fullname"
			type="text"
		/>
		<div className="grid gap-4 md:grid-cols-2">
			<TextInput
				name="password"
				placeholder="password"
				label="password"
				type={show ? "text":"password"}
			/>
			<TextInput
				name="confirm"
				placeholder="confirm password"
				label="confirm password"
				type={show ? "text":"password"}
			/>
		</div>
		<Toggle 
			value={show} 
			onChange={
				ev => setShow(ev.currentTarget.checked)
			} 
			label="show password" 
		/>
		<div className="flex justify-end gap-2">
			<button type="submit" className="btn">register</button>
		</div>
	</form>
}

const TAB = Object.freeze({
	LOGIN: 'LOGIN',
	REGISTER: 'REGISTER',
})

export default function AuthPage({ children }) {
	const [active, setActive] = useState(TAB.LOGIN)

	return <PageLayout>
		<article className="w-full md:max-w-xl bg-base-100 border border-base-300 shadow rounded-lg">
			<Tabs>
				<Tab 
					active={active == TAB.LOGIN}
					onClick={() => setActive(TAB.LOGIN)}
				>
					Login
				</Tab>
				<Tab 
					onClick={() => setActive(TAB.REGISTER)}
					active={active == TAB.REGISTER}
				>
					Register
				</Tab>
			</Tabs>
			<div className="p-6">
				{active == TAB.LOGIN ? (<Login/>): (<Register/>)}
			</div>
		</article>
	</PageLayout>
}

const PageLayout = ({ children }) => <div className="min-h-screen bg-base-200 flex justify-center items-center p-6">{children}</div>

const Tabs = ({ children }) => <div className="tabs grid grid-cols-2">
	{children}
</div>

const Tab = ({ children, active, ...rest }) => {
	return <a 
		{...rest} 
		className={`tab tab-lg tab-lifted text-2xl ${active && 'tab-active'}`}
	>
	{children}
	</a>
}