import { useState } from 'react';
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';

import VisitTab, { GoBack } from '@/features/interfaces';
import Loading from '@/app/loading';
import { titlecase } from './_interfaces';

import { AttributeDELETE } from './_fetch'

export function DeleteAttribute(props) {

	const [userdata] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()
	const [object] = useState(localdata.tab.context.object)
	const [loading, setLoading] = useState(false)

	function confirm() {
		setLoading(true)
		console.log("DELETING", object)
		AttributeDELETE(userdata, object.Meta.ID)
		.then((res) => console.log(res))
		.then(function () {
			setLocaldata(GoBack(localdata))
		})
		.catch(function (e) {
			console.error("FAILED TO DELETE", e)
		})
	}

	return (
		<div className='p-2 flex flex-col items-center'>
			{
				loading && <Loading/>
			}
			{
				!loading && <div>
					<button onClick={confirm} className="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">
						Confirm Delete attribute
					</button>
				</div>
			}
		</div>
	);
}
