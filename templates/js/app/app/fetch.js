import axios, {isCancel, AxiosError} from 'axios';

export const hostApi = "{{.Config.HostAPI}}"
export const webApi = "{{.Config.WebAPI}}"

export function WebFetch(method, url, body) {

	console.log("PublicFetch >>>", method, url, body);
	console.log(url)

	return fetch(
		webApi + url,
		{
			method: method,
			body: JSON.stringify(body),
		}
	)
}

export function PublicFetch(method, url, body) {

	console.log("PublicFetch >>>", method, url, body);
	console.log(url)

	return fetch(
		hostApi + url,
		{
			method: method,
			body: JSON.stringify(body),
		}
	)
}

export function AxiosPOST(user, url, formData) {
	const config = {
		headers: {
			'Authorization': user.headers.Authorization,
			'content-type': 'multipart/form-data',
		},
	};
	console.log(config)
	return axios.post(host+url, formData, config)
}

export default function SessionFetch(user, method, url, body) {

	console.log("SessionFetch", process.env.HANDCASH_APP_ID, process.env.ENVIRONMENT, ">>>", method, url, body);
	console.log(url)

	if (user == null) {
		console.error("userdata context needs to be provided");
		return
	}

	return fetch(
		hostApi + url,
		{
			method: method,
			body: JSON.stringify(body),
			headers: user.headers
		}
	)
}

export function OTPFetch(url) {

	console.log("OTPFetch >>>", url);
	console.log(url)

	return fetch(hostApi + url, {"method":"POST"})
}

export function UserAutocompleteGET(user, query) {
	return PublicFetch("GET", "api/users?function=autocomplete&query="+query)
}

export function AuthCheckEmail(email) {
	return PublicFetch("GET", "api/auth?function=query&email="+email)
}

export function AuthOtpGET(email) {
	return PublicFetch("GET", "api/auth?function=otp&email="+email)
}

export function AuthRegisterPOST(email, username) {
	return PublicFetch("POST", "api/auth?function=register&email="+email+"&username="+username)
}

export function AuthLoginPOST(otp) {
	return PublicFetch("POST", "api/auth?function=login&otp="+otp)
}

export function UserSessionGET(user) {
	return SessionFetch(user, "GET", "api/users?function=session")
}

export function UsernameGET(user, targetUserID) {
	return SessionFetch(user, "GET", "api/user?function=username&id="+targetUserID)
}

export function UserObjectGET(user, targetUserID) {
	return SessionFetch(user, "GET", "api/user?function=object&id="+targetUserID)
}

export function AssetsUser(user, id) {
	return SessionFetch(user, "GET", "api/assetlayer?function=assets&parent="+id)
}

export function AssetsWallet(user, id) {
	return SessionFetch(user, "GET", "api/assetlayer?function=walletassets&parent="+id)
}

export function ObjectPATCH(user, object, field, value) {
	const c = object.Meta.Class.substring(0, (object.Meta.Class.length-1))
	const payload = {
		"field": field,
		"value": value
	}
	console.log("PATCHING OBJECT:", user, "PATCH", "api/"+c+"?id="+object.Meta.ID, payload)
	return SessionFetch(user, "PATCH", "api/"+c+"?id="+object.Meta.ID, payload)
}

export function InboxConvosGET(user) {
	return SessionFetch(user, "GET", "api/mail?function=convos")
}

export function InboxMessagesGET(user, conversation) {
	return SessionFetch(user, "GET", "api/mail?function=messages&conversation="+conversation)
}

export function InboxSendMessage(user, msg) {
	return SessionFetch(user, "POST", "api/mail?", msg)
}

// handcash

export function HandcashPaymentPOST(authToken, payment) {
    return WebFetch("POST", "api/handcash/payment?authToken="+authToken, payment)
}

export function HandcashMintPOST(authToken, data) {
    return WebFetch("POST", "api/handcash/mint?authToken="+authToken, data)
}