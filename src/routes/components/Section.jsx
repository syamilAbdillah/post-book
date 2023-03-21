export default function Section({ children, ...rest }) {
	return <section className="p-6 border-t first:border-0 max-w-screen-md">{children}</section>
}