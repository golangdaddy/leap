import { useState, useEffect } from "react"
import { useUserContext } from "@/context/user"
import { useLocalContext } from "@/context/local"

export default function CollectionSelect(props) {

	console.log("SHOW INPUT", props)

	const [userdata] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()

	const [options, setOptions] = useState()

	function getList() {
		props.listFunction(
			userdata,
			props.parent,
		)
		.then((res) => res.json())
		.then((data) => {
			// return to previous interface
			var options = []
			data.forEach(item => {
				options.push(item[props.field])
			})
			console.log("CSELECT OPTIONS", options)
			setOptions(options)
		})
		.catch(function (e) {
			console.error(e)
			setLocaldata(GoBack(localdata))
		})
	}

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

	useEffect(() => {
		getList()
	}, [])

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
							<option key={i+1}>
							{item}
							</option>
						)
					})
				}
			</select>
			</div>
		</>
	);
}
