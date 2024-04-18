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

	const creationOrder = await handCashMinter.createCollectionOrder({
		name: 'My first collection',
		description: 'This is my first collection. Do not judge me.',
		mediaDetails: {
		  image: {
			url: 'https://res.cloudinary.com/handcash-iae/image/upload/v1685141160/round-handcash-logo_cj47fp_xnteyo_oy3nbd.png',
			contentType: 'image/png'
		  }
		}
	});

    res.status(200).json(creationOrder)
}