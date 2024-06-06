import { useEffect } from 'react';
import { useRouter } from 'next/navigation'

import Layout from '@/app/layout';

import Dashboard from '@/features/dashboard';

export default function HomePage({ payload }) {

	const router = useRouter()

	if (payload.redirect) {
		useEffect(() => {
		router.replace("/login")
	}, [])
	}

	return (
		<Layout className="flex flex-col items-center h-full">
			<Dashboard otp={payload.otp} handcashToken={payload.handcashToken} region={payload.region}/>
		</Layout>
	)
}

// This gets called on every request
export async function getServerSideProps(context) {

	var payload = {}
	if (context.query.otp) {
		payload["otp"] = context.query.otp
		payload["handcashToken"] = context.query.authToken ? context.query.authToken : ''
	} else {
		payload["redirect"] = true
	}
 
	return { props: { payload } }
}