import { PublicFetch, AxiosPOST } from '@/app/fetch';
import SessionFetch from '@/app/fetch';
import InputFormat from '@/inputs/inputFormat';

export function ElementsInitPOST(user, parentID, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/elements?function=init&parent="+parentID, p)
}

export function ElementUpdatePOST(user, id, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/element?function=update&id="+id, p)
}

export function ElementObjectGET(user, id) {
    return SessionFetch(user, "GET", "api/element?function=object&id="+id)
}

export function ElementsListGET(user, parentID, limit) {
    return SessionFetch(user, "GET", "api/elements?function=list&parent="+parentID+"&limit="+limit)
}

export function ElementsCountGET(user, parentID) {
    return SessionFetch(user, "GET", "api/elements?function=count&parent="+parentID)
}

export function ElementMoveUpPOST(user, id) {
    return SessionFetch(user, "POST", "api/element?function=up&id="+id)
}

export function ElementMoveDownPOST(user, id) {
    return SessionFetch(user, "POST", "api/element?function=down&id="+id)
}

export function ElementDELETE(user, id) {
    return SessionFetch(user, "DELETE", "api/element?id="+id)
}

export function ElementFunctionPOST(user, id, func) {
    return SessionFetch(user, "POST", "api/element?function="+func+"&id="+id)
}

export function ElementFileUpload(user, id, formData) {
    return AxiosPOST(user, "api/element?function=upload&id="+id, formData)
}

export function ElementArchiveUpload(user, parentID, formData) {
    return AxiosPOST(user, "api/elements?function=archiveupload&parent="+parentID, formData)
}