import axios, {isCancel, AxiosError} from 'axios';

export const host = "http://localhost:8080/"

export function PublicFetch(method, url, body) {

    console.log(">>>", method, url, body);
    console.log(url)

    return fetch(
        host + url,
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

    console.log(">>>", method, url, body);
    console.log(url)

    if (user == null) {
        console.error("fuck");
        return
    }

    return fetch(
        host + url,
        {
            method: method,
            body: JSON.stringify(body),
            headers: user.headers
        }
    )
}

export function OTPFetch(url) {

    console.log(">>>", url);
    console.log(url)

    return fetch(host + url, {"method":"POST"})
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
