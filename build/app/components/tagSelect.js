import { useEffect, useState } from "react"
import { useUserContext } from "@/context/user"

import { TagsGET } from "@/features/tags/_fetch"

export function TagSelect(props) {
    
	const [userdata, setUserdata] = useUserContext()
	const [tags, setTags] = useState()
	const [selectedTag, setSelectedTag] = useState()

    function inputChange(x) {
        console.log(x.target.value)
		setSelectedTag(x.target.value)

    }

	function addTag() {
		props.addTag(selectedTag)
	}

	useEffect(() => {
		TagsGET(userdata, "collection", props.collectionID)
		.then((res) => res.json())
		.then((data) => {
			console.log("SELECT TAG", data)
			setTags(data)
		})
		.catch(function (e) {
			  console.error("FAILED TO SEND", e)
		})
	}, [])

    return (
		<div className="flex flex-col">
			<select multiple className="py-2 px-4 border h-full"  onChange={inputChange}>
				<option key={0}></option>
				{
					tags && tags.map(function (item, i) {
						return (
							<option key={i+1} className="font-xl p-4 m-4" value={item.Meta.ID} style={{color:item.textColor,backgroundColor:item.backgroundColor}}>
							{item.name}
							</option>
						)	
					})
				}
			</select>
			<div className="my-4">
			{
			confirm && <button onClick={addTag} type="button" className="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900">
					<div style={{whiteSpace:"nowrap"}}>Add Tag</div>
				</button>
			}
			</div>
		</div>
    )
}