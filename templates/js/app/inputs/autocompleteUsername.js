import { useState } from "react"
import { useUserContext } from "@/context/user"

import { UserAutocompleteGET } from "@/app/fetch"

export default function AutocompleteUsername(props) {

	console.log("SHOW INPUT", props)

	const [userdata, setUserdata] = useUserContext()

	const [suggestions, setSuggestions] = useState([])
	const [included, setIncluded] = useState([])

	function inputChange(array) {
		setIncluded(array)
		var data = {
			"id": props.id,
			"type": "array",
			"value": array,
			"required": true,
		}
		props.inputChange(data)
		console.log("input change", data)
	}

	function changeEvent(e) {
		const id = e.target.id
		var value = e.target.value
		if (value.length < 3) {
			return
		}
		UserAutocompleteGET(userdata, value)
		.then((res) => res.json())
		.then((data) => {
			console.log("AC", data)
			setSuggestions(data)
		  })
		.catch(function (e) {
			console.error("FAILED TO SEND", e)
		})
	}

	function removeSuggestion(e) {
		const id = e.target.id
		if (!id.length) {
			return
		}
		var a = []
		for (var s in included) {
			if (included[s] == id) {
				continue
			}
			a.push(included[s])
		}
		inputChange(a)
		console.log("REMOVE", id)
	}

	function chooseSuggestion(e) {
		// removeSuggestion(e)
		const id = e.target.id
		if (!id.length) {
			return
		}
		console.log("CHOOSE", id)
		var data = []
		data.push(id)
		inputChange(data)
		// clear inputs ready for next suggestions
		setSuggestions([])
		document.getElementById("searchInput").value = ""
	}

	return (
		<>
			<div className="flex flex-col">
				<div className="text-l font-bold">{props.title}{props.required && "*"}</div>
				<div className="my-2"></div>
				<div className="flex flex-wrap p-4 round-lg">
					{
						included && included.map(function(username, i) {
							return (
								<div id={username} key={i} className="m-4 shadow-lg p-2">
									<div className="flex flex-row">
										<div className="uppercase">{username}</div>
										<div className="mx-2 cursor-pointer">
											<svg id={username} onClick={removeSuggestion} xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
											<path strokeLinecap="round" strokeLinejoin="round" d="M9.75 9.75l4.5 4.5m0-4.5l-4.5 4.5M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
											</svg>
										</div>
									</div>
								</div>
							)
						})
					}
				</div>
				<input className="py-2 px-4 border" id="searchInput" type={props.type} defaultValue={props.value} onChange={changeEvent} placeholder={props.placeholder} />
				<div className="flex flex-col">
					{
						suggestions && suggestions.map(function(item, i) {
							return (
								<div id={item.Username} key={i} className="uppercase m-4">
									<div className="flex flex-row">
										<div>{item.Username}</div>
										<div className="mx-2 cursor-pointer">
											<svg id={item.Username} onClick={chooseSuggestion} xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
											<path strokeLinecap="round" strokeLinejoin="round" d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
											</svg>
										</div>
									</div>
								</div>
							)
						})
					}
				</div>
			</div>
		</>
	);
}
