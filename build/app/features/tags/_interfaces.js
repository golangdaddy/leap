import * as React from 'react'

import { Tags } from './tags'
import { Tag } from './tag'
import { NewTag } from './newTag'
import { EditTag } from './editTag'
import { UploadTag } from './uploadTag'
import { DeleteTag } from './deleteTag'
import { InitUploadTag } from './initUploadTag'
import { InitUploadTags } from './initUploadTags'

export var TagInterfaces = {
	"deletetag": {
		level: -1,
		name: "Delete", 
		component: (<DeleteTag/>),
	},
	"newtag": {
		level: 6+2,
		name: "New Tag",
		component: (<NewTag />),
	},  
	"inituploadtag": {
		level: 6+2,
		name: "Upload Tag",
		component: (<InitUploadTag />),
	},
	"inituploadtags": {
		level: 6+2,
		name: "Upload Tags",
		component: (<InitUploadTags />),
	},
	"edittag": {
		level: -1,
		name: "Edit Tag", 
		component: (<EditTag />),
	},  
	"uploadtag": {
		level: 6+3,
		name: "Upload File", 
		component: (<UploadTag />),
	},  
	"tags": {
		level: 6+1,
		name: "Manage Tags", 
		component: (<Tags />),
		subsublinks: ["newtag"],
	},
	"tag": {
		level: 6+2,
		name: "Tag",
		sublinks: ["edittag", "deletetag"],
		subsublinks: [""],
		component: (<Tag />),
	},
}
