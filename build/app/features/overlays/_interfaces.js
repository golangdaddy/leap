import * as React from 'react'

import { Overlays } from './overlays'
import { Overlay } from './overlay'
import { NewOverlay } from './newOverlay'
import { EditOverlay } from './editOverlay'
import { UploadOverlay } from './uploadOverlay'
import { DeleteOverlay } from './deleteOverlay'
import { InitUploadOverlay } from './initUploadOverlay'
import { InitUploadOverlays } from './initUploadOverlays'

export var OverlayInterfaces = {
	"deleteoverlay": {
		level: -1,
		name: "Delete", 
		component: (<DeleteOverlay/>),
	},
	"newoverlay": {
		level: 6+2,
		name: "New Overlay",
		component: (<NewOverlay />),
	},  
	"inituploadoverlay": {
		level: 6+2,
		name: "Upload Overlay",
		component: (<InitUploadOverlay />),
	},
	"inituploadoverlays": {
		level: 6+2,
		name: "Upload Overlays",
		component: (<InitUploadOverlays />),
	},
	"editoverlay": {
		level: -1,
		name: "Edit Overlay", 
		component: (<EditOverlay />),
	},  
	"uploadoverlay": {
		level: 6+3,
		name: "Upload File", 
		component: (<UploadOverlay />),
	},  
	"overlays": {
		level: 6+1,
		name: "Manage Overlays", 
		component: (<Overlays />),
		subsublinks: ["newoverlay"],
	},
	"overlay": {
		level: 6+2,
		name: "Overlay",
		sublinks: ["editoverlay", "deleteoverlay"],
		subsublinks: [""],
		component: (<Overlay />),
	},
}
