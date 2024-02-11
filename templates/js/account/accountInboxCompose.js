import { useState } from "react"
import { useUserContext } from "@/context/user"

import { UserAutocompleteGET } from "@/app/fetch"

import { InboxSendMessage } from "@/app/fetch"

export default function AccountInboxCompose(props) {

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
		//props.inputChange(data)
		console.log("input change", data)
	}

	function changeEvent(e) {
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
		const n = parseInt(e.target.id)
		const item = included[n]
		var a = []
		for (var s in included) {
			if (included[s] == item.ID) {
				continue
			}
			a.push(included[s])
		}
		inputChange(a)
		console.log("REMOVE", id)
	}

	function chooseSuggestion(e) {
		// removeSuggestion(e)
		const n = parseInt(e.target.id)
		const item = suggestions[n]
		console.log("CHOOSE", item)
		var data = []
		data.push(item)
		inputChange(data)
		// clear inputs ready for next suggestions
		setSuggestions([])
		document.getElementById("searchInput").value = ""
	}

	function sendMessage() {
		const payload = {
			"recipients": included,
			"body": document.getElementById("body").value,
		}
		InboxSendMessage(userdata, payload)
		.then((res) => console.log(res))
		.catch(function (e) {
			console.error("FAILED TO SEND", e)
		})
	}

	const buttonStyle = {
		borderRadius: "12px",
		backgroundColor: "rgb(96, 165, 250)",
		border: "solid 0px",
		color: "white",
		padding: "6px 10px"
	}

	const bodystyle = {
		borderRadius: "12px",
		border: "solid 1px black",
	}

	return (
		<div className="flex flex-col w-full">
			{
				included && <div className="flex flex-wrap p-4 round-lg">
					{
						included.map(function(item, i) {
							return (
								<div key={i} className="m-4 shadow-lg p-2">
									<div className="flex flex-row">
										<div className="uppercase">{item.Username}</div>
										<div id={i} className="mx-2 cursor-pointer" onClick={removeSuggestion}>
											<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6" style={ {pointerEvents:"none"} }>
											<path strokeLinecap="round" strokeLinejoin="round" d="M9.75 9.75l4.5 4.5m0-4.5l-4.5 4.5M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
											</svg>
										</div>
									</div>
								</div>
							)
						})
					}
				</div>
			}
			<input className="py-2 px-4 border" id="searchInput" placeholder="Find users" onChange={changeEvent} />
			<div className="flex flex-col">
				{
					suggestions && suggestions.map(function(item, i) {
						return (
							<div id={item.Username} key={i} className="uppercase m-4">
								<div id={i} className="flex flex-row" onClick={chooseSuggestion}>
									<div style={ {pointerEvents:"none"} }>{item.Username}</div>
									<div className="mx-2 cursor-pointer" style={ {pointerEvents:"none"} }>
										<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
										<path strokeLinecap="round" strokeLinejoin="round" d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
										</svg>
									</div>
								</div>
							</div>
						)
					})
				}
			</div>
			<textarea id='body' className="w-full my-2 p-2" placeholder="your message" style={bodystyle}></textarea>
			<button style={buttonStyle} onClick={sendMessage}>Send</button>
		</div>
	);
}
