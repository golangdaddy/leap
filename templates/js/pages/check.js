import { useRouter } from 'next/navigation'

import Layout from '@/app/layout';

export default function PageCheck({ data }) {

	//const router = useRouter()


	return (
		<Layout className="flex flex-col items-center h-full">
			<div style={{minWidth:"40vw"}} className='bg-gray-200 flex flex-col items-center p-8'>
			<div className='m-8 font-bold text-xl'>Check your email: {data.email}</div>
			</div>
		</Layout>
	)
}

// This gets called on every request
export async function getServerSideProps(context) {

	console.log(context)
  
	const email = context.query.email
	const data = {}
	if (email != undefined) {
	  data["email"] = email
	}
   
	// Pass data to the page via props
	return { props: { data } }
  }