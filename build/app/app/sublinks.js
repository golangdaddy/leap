import * as React from 'react'

import { useLocalContext } from '@/context/local'
import { GetInterfaces } from '@/features/interfaces'
import VisitTab from '@/features/interfaces'

export default function Sublinks(props) {

	const [localdata, setLocaldata] = useLocalContext()

	console.log("SUBLINKS", localdata)

    function updateBackEvent(e) {
        updateBack(parseInt(e.target.id.split("crumb_")[1]))
    }

    function updateBack(crumbIndex) {
		console.log("C-INDEX", crumbIndex)
		const tab = localdata.breadcrumbs[crumbIndex]
        setLocaldata(VisitTab(localdata, tab.id, tab.context))
    }

    // update tabs handles the updated context and sends the user to a new interface
    function updateTabEvent(e) {
        updateTab(e.target.id)
    }

    // update tabs handles the updated context and sends the user to a new interface
    function updateTab(tabname) {
        setLocaldata(VisitTab(localdata, tabname, localdata.tab.context))
    }

	const interfaces = GetInterfaces()


	return (
		<>
		{
			(localdata.tab.sublinks?.length > 0) && <div className='flex flex-row justify-center w-full bg-gray-600 text-white px-2 py-2 text-sm'>
				{
					localdata.tab.sublinks.map(function (tabname, i) {
						const tab = interfaces[tabname]
						return (
							<div key={i} className='flex flex-col justify-center'>
								<div className='flex flex-row rounded-md border py-1 px-2 mx-2 bg-white'>
									<div id={tab.id} className='cursor-pointer text-gray-800' onClick={updateTabEvent}>{tab.name}</div>
								</div>
							</div>
						)
					})
				}
			</div>
		}
		</>
	)
}
