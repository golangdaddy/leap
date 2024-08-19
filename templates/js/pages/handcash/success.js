import Layout from '@/app/layout';
import { useEffect } from 'react';
import { useRouter } from 'next/router';
import { HandCashConnect } from '@handcash/handcash-connect';

export default function PageHandcash({ data }) {

	const router = useRouter();
	useEffect(() => {
		router.push(data.redirect)
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

	const handCashConnect = new HandCashConnect(creds);

	var data = {
		'redirect': handCashConnect.getRedirectionUrl(),
	}

	// Pass data to the page via props
	return { props: { data } }
}