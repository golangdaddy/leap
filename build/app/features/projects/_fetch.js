import { PublicFetch, AxiosPOST } from '@/app/fetch';
import SessionFetch from '@/app/fetch';
import InputFormat from '@/inputs/inputFormat';

export function ProjectsInitPOST(user, parentID, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/projects?function=init&parent="+parentID, p)
}

export function ProjectUpdatePOST(user, id, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/project?function=update&id="+id, p)
}

export function ProjectObjectGET(user, id) {
    return SessionFetch(user, "GET", "api/project?function=object&id="+id)
}

export function ProjectsListGET(user, parentID, limit) {
    return SessionFetch(user, "GET", "api/projects?function=list&parent="+parentID+"&limit="+limit)
}

export function ProjectsCountGET(user, parentID) {
    return SessionFetch(user, "GET", "api/projects?function=count&parent="+parentID)
}

export function ProjectMoveUpPOST(user, id) {
    return SessionFetch(user, "POST", "api/project?function=up&id="+id)
}

export function ProjectMoveDownPOST(user, id) {
    return SessionFetch(user, "POST", "api/project?function=down&id="+id)
}

export function ProjectDELETE(user, id) {
    return SessionFetch(user, "DELETE", "api/project?id="+id)
}

export function ProjectFunctionPOST(user, id, func) {
    return SessionFetch(user, "POST", "api/project?function="+func+"&id="+id)
}

export function ProjectFileUpload(user, id, formData) {
    return AxiosPOST(user, "api/project?function=upload&id="+id, formData)
}
