import { useState, useEffect } from 'react';
import Layout from '@/app/layout';
import { useRouter } from 'next/router';

import { AuthCheckEmail, AuthOtpGET, AuthRegisterPOST } from '@/app/fetch';

export default function PageLogin({ data }) {

	console.log("data", data)
	const router = useRouter();
	const [email, setEmail] = useState(data.email)

	const [emailState, setEmailState] = useState(0);

	function checkEmail() {
		const email = document.getElementById("otp_email").value
		if (email.length < 5) {
			return
		}
		console.log("requesting otp for", email)
		AuthCheckEmail(email)
		.then((res) => res.json())
		.then((data) => {
			setEmailState(data["result"])
		})
	}

	useEffect(() => {
		checkEmail()
	}, [])

	function sendOTPRequest() {
		const email = document.getElementById("otp_email").value.toLowerCase()
		console.log("requesting otp for", email)
		AuthOtpGET(email)
		.then(function () {
			router.push("/check?email="+email)
		})
  }

  return (
	<Layout>
		<div className="flex flex-col m-10 p-5 bg-gray-200">
			<div className='font-xl'>
				Get a 1 time password delivered to your email.
			</div>
			<div className='font-xl'>
				If you are already registered a button will appear when you type your email...
			</div>
			<div className='flex flex-col items-center'>
				<div className='m-5'>
				<input onChange={checkEmail} placeholder="Email Address" defaultValue={email} className="p-5 rounded-lg" id="otp_email" type="email"/>
				</div>
				<div className='m-5'>
				{
					(emailState == 1) && <button onClick={sendOTPRequest} className='text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700'>
					Get login link via inbox
					</button>
				}
				</div>
			</div>
		</div>
	</Layout>
  );
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