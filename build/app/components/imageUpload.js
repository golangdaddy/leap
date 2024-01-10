import { useState } from 'react';

import { useUserContext } from '@/context/user';

import { MediaPOST } from '@/app/fetch';
import { useLocalContext } from '@/context/local';

import Loading from '@/app/loading';
import { GoBack } from '@/features/interfaces';

export function ImageUpload(props) {

	const [userdata, _] = useUserContext();
	const [localdata, setLocaldata] = useLocalContext()
	const [file, setFile] = useState();

	const [ loading, setLoading ] = useState()
	const [ element ] = useState(localdata.tab.context.object)

	function handleChangeFile(event) {
		setFile(event.target.files[0])
	}

	function handleSubmitFile(event) {

		setLoading(true)

		event.preventDefault()
		const formData = new FormData();
		formData.append('file', file);
		formData.append('fileName', file.name);

		MediaPOST(userdata, element.Meta.ID, formData).then((response) => {
			console.log(response.data);
			if (props.done) {
				props.done()
				console.log("IMAGE CALLBACK")
			}
			setLoading(false)
			setLocaldata(GoBack(localdata))
		})
		.catch(function () {
			setLoading(false)
		});

	}

	return (
		<div className="flex flex-col">
			    {
					loading && <Loading/>
				}
				{
					!loading && <>
						<div className='my-3 font-medium'>Upload</div>
						<div className='flex flex-col'>
							<input type="file" onChange={handleChangeFile}/>
							<div>
								<button onClick={handleSubmitFile} className="my-5 text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">
									Upload Image
								</button>
							</div>
						</div>
					</>
				}
		</div>
	)
}