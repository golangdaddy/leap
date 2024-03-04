import { useState, useEffect } from 'react';
import Layout from '@/app/layout';
import { useRouter } from 'next/router';
import { AuthCheckEmail, AuthOtpGET, AuthRegisterPOST } from '@/app/fetch';

export default function PageRegister({ data }) {

	console.log("data", data)
	const router = useRouter();

	/*
	if (data.email) {
	document.getElementById("otp_email").value = data.email
	}
	*/

	const [isLoading, setLoading] = useState(false)
	const [emailState, setEmailState] = useState(0);

	function checkEmail() {
		const email = document.getElementById("otp_email").value
		console.log("requesting otp for", email)
		AuthCheckEmail(email)
		.then((res) => res.json())
		.then((data) => {
			setEmailState(data["result"])
		})
  	}

	function formatUsername(e) {
		var name = e.target.value.replace(" ", "_").toUpperCase()
		console.log(name)
		e.target.value = name
	}

	function sendRegisterRequest() {
		const email = document.getElementById("otp_email").value
		const username = document.getElementById("otp_username").value
		AuthRegisterPOST(email, username)
		.then(function () {
			router.push("/login?email="+email)
		})
	}

	return (
		<Layout>
			<div className="flex flex-col m-10 p-5">
				<div className='flex flex-col items-center'>
					<div className='font-2xl font-bold'>Register</div>
					<div className='flex flex-row m-5'>
						<input onChange={checkEmail} placeholder="Email Address" className="p-5 rounded-lg border" id="otp_email" type="email"/>
						{{if .Options.WhitelistDomains}}
						<select>
							{{range .Options.RegistrationDomains}}
							<option value="{{.}}">{{.}}</option>
							{{end}}
						</select>
						{{end}}
					</div>
					<div className='m-5'>
						<input onKeyUp={formatUsername} placeholder="username" className="p-5 rounded-lg" id="otp_username" type="text"/>
					</div>
					<button onClick={sendRegisterRequest} className='text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700'>
						Register
					</button>
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