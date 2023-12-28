export default function Checkbox(props) {

	console.log("SHOW INPUT", props)

	function changeEvent(e) {
		const id = e.target.id
		var value = e.target.checked

		props.inputChange(
			{
				"id": id,
				"type": "checkbox",
				"value": value,
				"required": false,
			}
		)
	}

	return (
		<>
			<div className="flex flex-row">
				<div>
					<input className="py-2 px-4 border" id={props.id} type="checkbox" defaultValue={props.value} onChange={changeEvent} placeholder={props.placeholder} />
				</div>
				<div className="text-l font-bold mx-3">{props.title}</div>
			</div>
		</>
	);
}
