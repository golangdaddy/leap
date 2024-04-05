import Layout from '@/app/layout';
import { useEffect } from 'react';
import { useRouter } from 'next/router';
import { HandCashConnect } from '@handcash/handcash-connect';

export default function PaymentPage({ data }) {

	const router = useRouter();
	useEffect(() => {
//		router.push(data.redirect)
	}, [])

	return (
		<Layout/>
	);
}

// This gets called on every request
export async function getServerSideProps(context) {

	var creds = { 
		appId: process.env.HANDCASH_APP_ID,
		appSecret: process.env.HANDCASH_APP_SECRET,
	}

	console.log(creds)
	
	const handCashConnect = new HandCashConnect(creds);

	const account = handCashConnect.getAccountFromAuthToken(context.query.authToken);
	const paymentParameters = {
		description: "Hold my beer!🍺",
		payments: [
			{ currencyCode: 'USD', sendAmount: 0.05, destination: '$gopher' },
		]
	};
	const paymentResult = await account.wallet.pay(paymentParameters);
	console.log(paymentResult)

	var data = {}

	// Pass data to the page via props
	return { props: { data } }
}