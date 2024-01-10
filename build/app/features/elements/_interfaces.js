import * as React from 'react'

import { Elements } from './elements'
import { Element } from './element'
import { NewElement } from './newElement'
import { EditElement } from './editElement'
import { UploadElement } from './uploadElement'
import { DeleteElement } from './deleteElement'
import { InitUploadElement } from './initUploadElement'
import { InitUploadElements } from './initUploadElements'

export var ElementInterfaces = {
	"deleteelement": {
		level: -1,
		name: "Delete", 
		component: (<DeleteElement/>),
	},
	"newelement": {
		level: 6+2,
		name: "New Element",
		component: (<NewElement />),
	},  
	"inituploadelement": {
		level: 6+2,
		name: "Upload Element",
		component: (<InitUploadElement />),
	},
	"inituploadelements": {
		level: 6+2,
		name: "Upload Elements",
		component: (<InitUploadElements />),
	},
	"editelement": {
		level: -1,
		name: "Edit Element", 
		component: (<EditElement />),
	},  
	"uploadelement": {
		level: 6+3,
		name: "Upload File", 
		component: (<UploadElement />),
	},  
	"elements": {
		level: 6+1,
		name: "Manage Elements", 
		component: (<Elements />),
		subsublinks: ["newelement", "inituploadelements"],
	},
	"element": {
		level: 6+2,
		name: "Element",
		sublinks: ["editelement", "uploadelement", "deleteelement"],
		subsublinks: ["tags",""],
		component: (<Element />),
	},
}
