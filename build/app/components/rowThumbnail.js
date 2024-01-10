import { useState } from "react"

export function RowThumbnail(props) {
    
    const [confirm, setConfirm] = useState(false)

    function deleteLayer() {
        props.delete(props.id)
    }

    function toggleConfirm() {
        setConfirm(!confirm)
    }

    const iconStyle = {width:"30px",height:"30px"}

    return (
		<div className="flex flex-row" style={{maxWidth:"10vw"}}>
			<img src={props.source}/>
		</div>
    )
}