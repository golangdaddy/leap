import { PublicFetch, AxiosPOST } from '@/app/fetch';
import SessionFetch from '@/app/fetch';
import InputFormat from '@/inputs/inputFormat';

export function OverlaysInitPOST(user, parentID, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/overlays?function=init&parent="+parentID, p)
}

export function OverlayUpdatePOST(user, id, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/overlay?function=update&id="+id, p)
}

export function OverlayObjectGET(user, id) {
    return SessionFetch(user, "GET", "api/overlay?function=object&id="+id)
}

export function OverlaysListGET(user, parentID, limit) {
    return SessionFetch(user, "GET", "api/overlays?function=list&parent="+parentID+"&limit="+limit)
}

export function OverlaysCountGET(user, parentID) {
    return SessionFetch(user, "GET", "api/overlays?function=count&parent="+parentID)
}

export function OverlayMoveUpPOST(user, id) {
    return SessionFetch(user, "POST", "api/overlay?function=up&id="+id)
}

export function OverlayMoveDownPOST(user, id) {
    return SessionFetch(user, "POST", "api/overlay?function=down&id="+id)
}

export function OverlayDELETE(user, id) {
    return SessionFetch(user, "DELETE", "api/overlay?id="+id)
}

export function OverlayFunctionPOST(user, id, func) {
    return SessionFetch(user, "POST", "api/overlay?function="+func+"&id="+id)
}

export function OverlayFileUpload(user, id, formData) {
    return AxiosPOST(user, "api/overlay?function=upload&id="+id, formData)
}
