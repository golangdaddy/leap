import dotenv from 'dotenv';
import { HandCashMinter, Environments, Types, HandCashConnect } from '@handcash/handcash-connect';

dotenv.config();

export default async(req, res) => {

	console.log('creating collection')

	const { authToken } = req.query;

	var creds = { 
		authToken: authToken,
		appId: process.env.HANDCASH_APP_ID,
		appSecret: process.env.HANDCASH_APP_SECRET,
	}
	
	console.log(creds)

	const handCashMinter = HandCashMinter.fromAppCredentials(creds);

	(async () => {
		const creationOrder = await handCashMinter.createCollection({
			name: 'HandCash Team Caricatures',
			description: 'A unique collection of caricatures of the HandCash team',
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

    res.status(200).json([])
}