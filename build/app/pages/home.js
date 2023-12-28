import { useRouter } from 'next/navigation'

import Layout from '@/app/layout';

import Dashboard from '@/features/dashboard';

export default function HomePage({ payload }) {

	const router = useRouter()

	if (payload.redirect) {
		useEffect(() => {
		router.replace("/dashboard/otp")
	}, [])
	}

	return (
		<Layout className="flex flex-col items-center h-full">
			<Dashboard otp={payload.otp} region={payload.region}/>
		</Layout>
	)
}

// This gets called on every request
export async function getServerSideProps(context) {

	var payload = {}
	if (context.query.otp) {
		payload["otp"] = context.query.otp
	} else {
		payload["redirect"] = true
	}
 
	return { props: { payload } }
}