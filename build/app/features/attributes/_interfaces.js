import * as React from 'react'

import { Attributes } from './attributes'
import { Attribute } from './attribute'
import { NewAttribute } from './newAttribute'
import { EditAttribute } from './editAttribute'
import { UploadAttribute } from './uploadAttribute'
import { DeleteAttribute } from './deleteAttribute'
import { InitUploadAttribute } from './initUploadAttribute'
import { InitUploadAttributes } from './initUploadAttributes'

export var AttributeInterfaces = {
	"deleteattribute": {
		level: -1,
		name: "Delete", 
		component: (<DeleteAttribute/>),
	},
	"newattribute": {
		level: 4+2,
		name: "New Attribute",
		component: (<NewAttribute />),
	},  
	"inituploadattribute": {
		level: 4+2,
		name: "Upload Attribute",
		component: (<InitUploadAttribute />),
	},
	"inituploadattributes": {
		level: 4+2,
		name: "Upload Attributes",
		component: (<InitUploadAttributes />),
	},
	"editattribute": {
		level: -1,
		name: "Edit Attribute", 
		component: (<EditAttribute />),
	},  
	"uploadattribute": {
		level: 4+3,
		name: "Upload File", 
		component: (<UploadAttribute />),
	},  
	"attributes": {
		level: 4+1,
		name: "Manage Attributes", 
		component: (<Attributes />),
		subsublinks: ["newattribute"],
	},
	"attribute": {
		level: 4+2,
		name: "Attribute",
		sublinks: ["editattribute", "deleteattribute"],
		subsublinks: [""],
		component: (<Attribute />),
	},
}
