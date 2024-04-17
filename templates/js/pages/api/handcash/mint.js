import { HandCashMinter } from '@handcash/handcash-connect';

export default async(req, res) => {

	console.log('creating items')

    const jsonData = req.body;

	const { authToken } = req.query;

    // Parse JSON data into an object
    const mint = JSON.parse(jsonData);

	var creds = { 
		authToken: authToken,
		appId: process.env.HANDCASH_APP_ID,
		appSecret: process.env.HANDCASH_APP_SECRET,
	}
	
	const handCashMinter = HandCashMinter.fromAppCredentials(creds);

	const creationOrder = await handCashMinter.createItemsOrder(mint);
	const items = await handCashMinter.getOrderItems(creationOrder.id);
	
	console.log('Items created', items);

    res.status(200).json(paymentResult)
}