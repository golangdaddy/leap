export default function Select(props) {

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
					props.options.map(function (item, i) {
						return (
							<option key={i+1} value={item}>
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
