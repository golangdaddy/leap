import { useLocalContext } from '@/context/local';

import VisitTab from "@/features/interfaces"

export function RowOrder(props) {
    
	const [localdata, setLocaldata] = useLocalContext()

    const iconStyle = {width:"30px",height:"30px"}

	function moveUp() {
		props.moveUp(props.id)
	}

	function moveDown() {
		props.moveDown(props.id)
	}

    return (
		<div className="flex flex-row">
			<div className="flex flex-row justify-center items-center m-4 cursor-pointer" style={iconStyle}>
			{
				(props.id >= 0) && (props.id < (props.listLength-1)) && <div className="flex flex-row justify-between items-center m-4 cursor-pointer" style={iconStyle} onClick={moveDown}>
				<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
				<path strokeLinecap="round" strokeLinejoin="round" d="M19.5 13.5L12 21m0 0l-7.5-7.5M12 21V3" />
				</svg>
				</div>
			}
			{
				(props.id > 0) && <div className="flex flex-row justify-between items-center m-4 cursor-pointer" style={iconStyle} onClick={moveUp}>
				<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
				<path strokeLinecap="round" strokeLinejoin="round" d="M4.5 10.5L12 3m0 0l7.5 7.5M12 3v18" />
				</svg>
			</div>
			}
			</div>
		</div>
    )
}