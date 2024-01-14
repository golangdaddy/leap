import { useState } from "react"

export function RowDelete(props) {
    
    const [confirm, setConfirm] = useState(false)

    function deleteLayer() {
        props.delete(props.id)
    }

    function toggleConfirm() {
        setConfirm(!confirm)
    }

    const iconStyle = {width:"30px",height:"30px"}

    return (
		<div className="flex flex-row">
			<div className="flex flex-row justify-center items-center m-4 cursor-pointer" style={iconStyle} onClick={toggleConfirm}>
				<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
				<path strokeLinecap="round" strokeLinejoin="round" d="M20.25 7.5l-.625 10.632a2.25 2.25 0 01-2.247 2.118H6.622a2.25 2.25 0 01-2.247-2.118L3.75 7.5m6 4.125l2.25 2.25m0 0l2.25 2.25M12 13.875l2.25-2.25M12 13.875l-2.25 2.25M3.375 7.5h17.25c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125z" />
				</svg>
			</div>
			{
				confirm && <div className="flex flex-col justify-center">
					<div>
						<button onClick={deleteLayer} type="button" className="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900">
							<div style={{whiteSpace:"nowrap"}}>Confirm Delete</div>
						</button>
					</div>
				</div>
			}
		</div>
    )
}