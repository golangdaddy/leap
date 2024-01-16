export default function Input(props) {

	console.log("SHOW INPUT", props)

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
		console.log("ONLOAD", props.type, data)
		props.inputChange(
			data
		)
	}

	function changeEvent(e) {
		const id = e.target.id
		var value = e.target.value
		if (props.type == "number") {
			value = parseFloat(value)
		}
		props.inputChange(
			{
				"id": id,
				"type": props.type,
				"value": value,
				"required": props.required,
			}
		)
	}

	return (
		<>
			<div className="flex flex-col">
				<div className="text-l font-bold">{props.title}{props.required && "*"}</div>
				<div className="m-2"></div>
				{
					props.type == "number" && <div>
						<input disabled={(props.disabled == true)} className="p-4 border" id={props.id} type={props.type} defaultValue={props.value} onChange={changeEvent} onLoad={changeEventOnload}/>
					</div>
				}
				{
					props.type != "number" && <input disabled={(props.disabled == true)} className="p-4 border" id={props.id} type={props.type} defaultValue={props.value} onChange={changeEvent} onLoad={changeEventOnload} placeholder={props.placeholder} />
				}
			</div>
		</>
	);
}
