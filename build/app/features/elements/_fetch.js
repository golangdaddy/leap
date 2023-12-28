import { PublicFetch } from '@/app/fetch';
import SessionFetch from '@/app/fetch';
import InputFormat from '@/inputs/inputFormat';

export function ElementsInitPOST(user, layer, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/elements?function=init&layer="+layer, p)
}

export function ElementUpdatePOST(user, id, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/element?function=update&id="+id, p)
}

export function ElementObjectGET(user, id) {
    return SessionFetch(user, "GET", "api/element?function=object&id="+id)
}

export function ElementsListGET(user, layer, limit) {
    return SessionFetch(user, "GET", "api/elements?function=list&layer="+layer+"&limit="+limit)
}

export function ElementsCountGET(user, collectionID) {
    return SessionFetch(user, "GET", "api/elements?function=count&collection="+collectionID)
}

export function ElementDELETE(user, id) {
    return SessionFetch(user, "DELETE", "api/element?id="+id)
}

export function ElementFunctionPOST(user, id, func) {
    return SessionFetch(user, "POST", "api/element?function="+func+"&id="+id)
}