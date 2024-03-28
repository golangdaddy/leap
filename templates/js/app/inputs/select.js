import { useState, useEffect } from "react"
import { useUserContext } from "@/context/user"
import SessionFetch from "@/app/fetch"

export default function Select(props) {

	const [userdata, setUserdata] = useUserContext()

	const [options, setOptions] = useState(props.options)

	console.log("SHOW INPUT", props)

	useEffect(() => {
		if (!props.reference) {
			console.log("not getting reference data")
			return
		}
		const parentMeta = props.referenceParent.Meta
		const url = "api/" + props.reference + "?parent=" + parentMeta.Context.Parents[parentMeta.Context.Parents.length -1] + "&function=list&mode=created" 
		console.log("getting reference data:", url)
		SessionFetch(userdata, "GET", url)
		.then((res) => res.json())
		.then((data) => {
			console.log("referened data", data)
			setOptions(data)
		}).catch((e) => console.error(e))
	}, [])

	function changeEventOnload(e) {
		const id = e.target.id
		var value = e.target.value
		if (props.type == "number") {
			value = parseFloat(value)
		}
		const data = {
			"id": id,
			"type": props.type,
			"value": value,
			"required": props.required,
		}
		console.log("ONLOAD", data)
		props.inputChange(data)
	}

	function changeEvent(e) {
		const id = e.target.id
		var value = e.target.value
		if (props.type == "number") {
			value = parseFloat(value)
		}
		props.inputChange({
			"id": id,
			"type": props.type,
			"value": value,
			"required": props.required,
		})
	}

	return (
		<>
		<div className="flex flex-col">
			<div className="text-l font-bold">{props.title}{props.required && "*"}</div>
			<div className="m-2"></div>
			<select disabled={(props.disabled == true)} className="py-2 px-4 border" id={props.id} defaultValue={props.value} onChange={changeEvent} onLoad={changeEventOnload}>
				<option key={0}></option>
				{
					options && options.map(function (item, i) {
						return (
							<option key={i+1} value={item.Meta?.Name ? item.Meta.Name : item}>
							{item.Meta?.Name ? item.Meta.Name : item}
							</option>
						)
					})
				}
			</select>
			</div>
		</>
	);
}
