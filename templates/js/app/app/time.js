
function convert(x) {
    const Tmilliseconds = x * 1000
    const TdateObject = new Date(Tmilliseconds)
    return TdateObject.toLocaleString()
}

export default function Timestamp(props) {
    const t = convert(props.timestamp)
    return (
        <>
            {t}
        </>
    )
}