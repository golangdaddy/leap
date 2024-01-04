import { PublicFetch, AxiosPOST } from '@/app/fetch';
import SessionFetch from '@/app/fetch';
import InputFormat from '@/inputs/inputFormat';

export function FontsInitPOST(user, parentID, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/fonts?function=init&parent="+parentID, p)
}

export function FontUpdatePOST(user, id, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/font?function=update&id="+id, p)
}

export function FontObjectGET(user, id) {
    return SessionFetch(user, "GET", "api/font?function=object&id="+id)
}

export function FontsListGET(user, parentID, limit) {
    return SessionFetch(user, "GET", "api/fonts?function=list&parent="+parentID+"&limit="+limit)
}

export function FontsCountGET(user, collectionID) {
    return SessionFetch(user, "GET", "api/fonts?function=count&collection="+collectionID)
}

export function FontDELETE(user, id) {
    return SessionFetch(user, "DELETE", "api/font?id="+id)
}

export function FontFunctionPOST(user, id, func) {
    return SessionFetch(user, "POST", "api/font?function="+func+"&id="+id)
}

export function FontFileUpload(user, id, formData) {
    return AxiosPOST(user, "api/font?function=upload&id="+id, formData)
}