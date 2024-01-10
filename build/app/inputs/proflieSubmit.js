import { useUserContext } from "@/context/user"
import { useLocalContext } from "@/context/local"
import { ProfileChangePATCH } from "@/features/account/_fetch"
import { UserSessionGET } from "@/app/fetch"

import { GoBack } from "@/app/interfaces"

export default function ProfileSubmit(props) {

	const [userdata, setUserdata] = useUserContext()
    const [localdata, setLocaldata] = useLocalContext()

    var text = props.text
    if (!text) {
        text = "Submit"
    }

    function isValid() {
        var checked = 0
        if (!props.assert) {
            return true
        }
        for (var x = 0; x < props.assert.length; x++) {
            const key = props.assert[x]
            if (props.inputs[key]?.length) {
                checked++
            }
        }
        return (checked == props.assert?.length)
    }

	function submitUpdate() {
		
		ProfileChangePATCH(
			userdata,
			props.account,
			props.inputs
		)
		.then((res) => console.log(res))
		.then(function () {
			updateUser()
		})
		.then(function () {
			// return to previous interface
			setLocaldata(GoBack(localdata))
		})
		.catch(function (e) {
			console.error("FAILED TO SEND", e)
		})
	}

	function updateUser() {
		UserSessionGET(userdata)
		.then((res) => res.json())
		.then((data) => {
			var newData = Object.assign({}, userdata)
			for (var k in data) {
				newData[k] = data[k]
			}
			setUserdata(newData)
			console.log("UPDATE USER", userdata)
		})
	}

    return (
        <div className="flex flex-col">
        <hr className='my-4'/>
        { !isValid() &&
            <div>
                Fields Marked with * are required...
            </div>
        }
        { isValid() &&
            <div>
                <button onClick={submitUpdate} className="my-4 text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">
                    { text }
                </button>
            </div>
        }
        </div>
    );
}
  