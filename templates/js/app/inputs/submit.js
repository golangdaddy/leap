import { useState } from "react"

export default function Submit(props) {

	const [ clicked, setClicked ] = useState(false)

	var requiredIndex = {}
	if (props.assert) {
		for (var i in props.assert) {
			requiredIndex[props.assert[i]] = true
		}
	}

	var text = props.text
	if (!text) {
		text = "Submit"
	}

	var setValid = 0
	var isValid = true
	for (var input in props.inputs) {
		var i = props.inputs[input]
		console.log("VALIDATE INPUTS", input, props.inputs)
		if (requiredIndex[i.id]) {
			if (i.required) {
				switch (i.type) {
					case "color":
						if (i.value.length < 7) {
							isValid = false
							continue
						}
						setValid++
						break
					case "text":
						if (i.value == "") {
							isValid = false
							continue
						}
						setValid++
						break
					case "string":
						if (i.value == "") {
							isValid = false
							continue
						}
						setValid++
						break
					case "number":
						if (parseInt(i.value) < 0) {
							isValid = false
							continue
						}
						setValid++
						break
					case "array":
						if (i.value.length < 1) {
							console.log("array value", i.value)
							isValid = false
							continue
						}
						setValid++
						break
					default:
						console.error("ERROR VALIDATING FORM: "+i.type)
					}
			}
		}
	}
	if (isValid) {
		if (Object.keys(props.inputs).length == 0) {
			isValid = false
		}
		if (setValid != props.assert?.length) {
			isValid = false
		}
	} else {
		console.error("failed to validate")
	}

	function submitForm() {
		setClicked(true)
		props.submit(props.inputs)
	}

	return (
		<div className="flex flex-col">
		<hr className='my-4'/>
		{ !isValid &&
			<div>
				Fields Marked with * are required...
			</div>
		}
		{ isValid && !clicked &&
			<div>
				<button onClick={submitForm} className="my-4 text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">
					{ text }
				</button>
			</div>
		}
		</div>
	);
}
  