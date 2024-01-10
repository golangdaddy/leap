import { useState, useEffect } from 'react';

export function LimitedDescription(props) {

		const [description, setDescription] = useState(props.value)
		const [chars, setChars] = useState(0)
		function updateDesc() {
			var content = document.getElementById(props.ident);
			content.value = content.value.substring(0, 800);
			setChars(content.value.length);
			setDescription(content.value)
			props.inputChange(props.id, props.value)
		}

		return (
		<div className="flex flex-col">
			<div className="text-l font-bold">{props.title}</div>
			<div className="m-2"></div>
			<div className='flex flex-col'>
				<textarea className='border' style={{"height":"200px"}} id={props.ident} onKeyUp={updateDesc} defaultValue={description}></textarea>
				<div className='my-2 font-light'>({ props.maxChars - chars } chars left)</div>
			</div>
		</div>
		)
}