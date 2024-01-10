import { PublicFetch, AxiosPOST } from '@/app/fetch';
import SessionFetch from '@/app/fetch';
import InputFormat from '@/inputs/inputFormat';

export function LayersInitPOST(user, parentID, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/layers?function=init&parent="+parentID, p)
}

export function LayerUpdatePOST(user, id, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/layer?function=update&id="+id, p)
}

export function LayerObjectGET(user, id) {
    return SessionFetch(user, "GET", "api/layer?function=object&id="+id)
}

export function LayersListGET(user, parentID, limit) {
    return SessionFetch(user, "GET", "api/layers?function=list&parent="+parentID+"&limit="+limit)
}

export function LayersCountGET(user, parentID) {
    return SessionFetch(user, "GET", "api/layers?function=count&parent="+parentID)
}

export function LayerMoveUpPOST(user, id) {
    return SessionFetch(user, "POST", "api/layer?function=up&id="+id)
}

export function LayerMoveDownPOST(user, id) {
    return SessionFetch(user, "POST", "api/layer?function=down&id="+id)
}

export function LayerDELETE(user, id) {
    return SessionFetch(user, "DELETE", "api/layer?id="+id)
}

export function LayerFunctionPOST(user, id, func) {
    return SessionFetch(user, "POST", "api/layer?function="+func+"&id="+id)
}

export function LayerFileUpload(user, id, formData) {
    return AxiosPOST(user, "api/layer?function=upload&id="+id, formData)
}
