import * as React from 'react'

import { useLocalContext } from '@/context/local'
import { GetInterfaces } from '@/features/interfaces'
import VisitTab from '@/features/interfaces'

export default function Back(props) {

	const [localdata, setLocaldata] = useLocalContext()

	console.log("<<BACK", localdata)

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

	function updateRegion(e) {
		const id = e.target.value
		var data = Object.assign({}, localdata)
		data.region = id
		setLocaldata(VisitTab(data, "home"))
		console.log("NEW REGION", data)

	}

	return (
		<>
		{
			localdata?.breadcrumbs && <div className='w-full bg-gray-700 text-white flex flex-row justify-start'>

				<div className='text-white m-3 flex flex-row'>
					{
						localdata.breadcrumbs.map(function (tab, i) {

							if (!tab) {
								console.log("skipping missed tab")
								return
							}
							var displayName = tab.name

							if (tab.context?._) {
								displayName = tab.context._
							}

							return (
								<div key={i} className='flex flex-col justify-center'>
									<div className='flex flex-row'>
										<div className='mx-2'>/</div>
										<div className="cursor-pointer" id={"crumb_"+i} onClick={ (i == (localdata.breadcrumbs.length - 1) ) ? function () {} : updateBackEvent}>{displayName}</div>
									</div>
								</div>
							)
						})
					}
				</div>
			</div>
		}
		</>
	)
}
