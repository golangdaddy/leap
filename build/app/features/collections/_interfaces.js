import * as React from 'react'

import { Collections } from './collections'
import { Collection } from './collection'
import { NewCollection } from './newCollection'
import { EditCollection } from './editCollection'
import { UploadCollection } from './uploadCollection'
import { DeleteCollection } from './deleteCollection'
import { InitUploadCollection } from './initUploadCollection'
import { InitUploadCollections } from './initUploadCollections'

export var CollectionInterfaces = {
	"deletecollection": {
		level: -1,
		name: "Delete", 
		component: (<DeleteCollection/>),
	},
	"newcollection": {
		level: 2+2,
		name: "New Collection",
		component: (<NewCollection />),
	},  
	"inituploadcollection": {
		level: 2+2,
		name: "Upload Collection",
		component: (<InitUploadCollection />),
	},
	"inituploadcollections": {
		level: 2+2,
		name: "Upload Collections",
		component: (<InitUploadCollections />),
	},
	"editcollection": {
		level: -1,
		name: "Edit Collection", 
		component: (<EditCollection />),
	},  
	"uploadcollection": {
		level: 2+3,
		name: "Upload File", 
		component: (<UploadCollection />),
	},  
	"collections": {
		level: 2+1,
		name: "Manage Collections", 
		component: (<Collections />),
		subsublinks: ["newcollection"],
	},
	"collection": {
		level: 2+2,
		name: "Collection",
		sublinks: ["editcollection", "deletecollection"],
		subsublinks: ["attributes","layers",""],
		component: (<Collection />),
	},
}
