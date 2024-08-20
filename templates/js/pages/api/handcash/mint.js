import { HandCashMinter } from '@handcash/handcash-connect';

export default async(req, res) => {


    const jsonData = req.body;

	console.log('creating items', jsonData)

    // Parse JSON data into an object
    const mint = [{
		name: 'An example item',
		rarity: 'Common',
		attributes: [
		  {
			name: 'Edition',
			value: 'First',
			displayType: 'string',
		  },
		],
		mediaDetails: {
		  image: {
			url: 'https://res.cloudinary.com/handcash-iae/image/upload/v1702398906/items/da2qv0oqma0hs3gqevg7.webp',
			imageHighResUrl: 'https://res.cloudinary.com/handcash-iae/image/upload/v1697465892/items/gh7tsn11svhx7z943znv.png',
			contentType: 'image/png',
		  },
		},
		color: '#B19334',
		quantity: 5,
	}]

	const { authToken } = req.query;

	var creds = { 
		authToken: authToken,
		appId: process.env.HANDCASH_APP_ID,
		appSecret: process.env.HANDCASH_APP_SECRET,
	}
	
	const handCashMinter = HandCashMinter.fromAppCredentials(creds);

	(async () => {
		const creationOrder = await handCashMinter.createCollectionOrder({
			name: 'Leap Test',
			description: 'A test of Leap',
			mediaDetails: {
			  image: {
				url: 'https://res.cloudinary.com/handcash-iae/image/upload/v1685141160/round-handcash-logo_cj47fp_xnteyo_oy3nbd.png',
				contentType: 'image/png'
			  }
			}
		  });
	   
		console.log(`Items order created, items are being created asynchronously`);
		const items = await handCashMinter.getOrderItems(creationOrder.id);
		console.log(`Collection Created, collectionId: ${items[0].id}`);
	})();
	
	console.log('Items created', mint);

    res.status(200).json("OK")
}