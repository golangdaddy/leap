import * as React from 'react'

import { Layers } from './layers'
import { Layer } from './layer'
import { NewLayer } from './newLayer'
import { EditLayer } from './editLayer'
import { UploadLayer } from './uploadLayer'
import { DeleteLayer } from './deleteLayer'
import { InitUploadLayer } from './initUploadLayer'
import { InitUploadLayers } from './initUploadLayers'

export var LayerInterfaces = {
	"deletelayer": {
		level: -1,
		name: "Delete", 
		component: (<DeleteLayer/>),
	},
	"newlayer": {
		level: 4+2,
		name: "New Layer",
		component: (<NewLayer />),
	},  
	"inituploadlayer": {
		level: 4+2,
		name: "Upload Layer",
		component: (<InitUploadLayer />),
	},
	"inituploadlayers": {
		level: 4+2,
		name: "Upload Layers",
		component: (<InitUploadLayers />),
	},
	"editlayer": {
		level: -1,
		name: "Edit Layer", 
		component: (<EditLayer />),
	},  
	"uploadlayer": {
		level: 4+3,
		name: "Upload File", 
		component: (<UploadLayer />),
	},  
	"layers": {
		level: 4+1,
		name: "Manage Layers", 
		component: (<Layers />),
		subsublinks: ["newlayer"],
	},
	"layer": {
		level: 4+2,
		name: "Layer",
		sublinks: ["editlayer", "deletelayer"],
		subsublinks: ["overlays","elements","tags",""],
		component: (<Layer />),
	},
}
