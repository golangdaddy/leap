export function TagPreview(props) {

	var style = {
		backgroundColor: props.backgroundColor,
		color: props.textColor,
		whiteSpace: "nowrap",
	}

	switch (props.size) {
		case "tiny":

			break
		case "small":
		case "medium":
		case "large":
	} 

	return (
		<div className='m-2 p-2 text-2xl rounded-lg uppercase' style={style}>
			{props.name}
		</div>
	)

}