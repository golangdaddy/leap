import { useRouter } from 'next/navigation'

import Layout from '@/app/layout';

export default function HomePage({ payload }) {

	//const router = useRouter()


	return (
		<Layout className="flex flex-col items-center h-full">
			Hello World
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