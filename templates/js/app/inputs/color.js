
import React from 'react'; 
import { ColorPicker, useColor } from "react-color-palette"; 
import "react-color-palette/css"; 

export default function Color(props) {

	const [color, updateColor] = useColor("hex", props.value); 

	console.log("SHOW INPUT", props)

	function setColor(x) {
		updateColor(x)
		props.inputChange(
			{
				"id": props.id,
				"ftype": props.ftype,
				"value": x.hex,
				"required": props.required,
			}
		)
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
		<div className="flex flex-col" style={{maxWidth:"50vw"}}>
			<div className="text-l font-bold">{props.title}{props.required && "*"}</div>
			<div className="m-2"></div>
			<ColorPicker width={456} height={200} color={color} onChange={setColor} hideHSV /> 
		</div>
  	);
}
