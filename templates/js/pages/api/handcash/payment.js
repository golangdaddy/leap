import { HandCashConnect } from '@handcash/handcash-connect';

export default async(req, res) => {

    const jsonData = req.body;

	const { authToken } = req.query;

    // Parse JSON data into an object
    const payment = JSON.parse(jsonData);

	var creds = {
		appId: process.env.HANDCASH_APP_ID,
		appSecret: process.env.HANDCASH_APP_SECRET,
	}
	
	const handCashConnect = new HandCashConnect(creds);

	const account = handCashConnect.getAccountFromAuthToken(authToken);
/*
    const paymentParameters = {
		description: "Hold my beer!🍺",
		payments: [
			{ currencyCode: 'USD', sendAmount: 0.01, destination: 'gopher' },
		]
	};
*/
    const paymentResult = await account.wallet.pay(payment);
	console.log(paymentResult)

    res.status(200).json(paymentResult)
}