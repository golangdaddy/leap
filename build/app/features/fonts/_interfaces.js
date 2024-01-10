import * as React from 'react'

import { Fonts } from './fonts'
import { Font } from './font'
import { NewFont } from './newFont'
import { EditFont } from './editFont'
import { UploadFont } from './uploadFont'
import { DeleteFont } from './deleteFont'
import { InitUploadFont } from './initUploadFont'
import { InitUploadFonts } from './initUploadFonts'

export var FontInterfaces = {
	"deletefont": {
		level: -1,
		name: "Delete", 
		component: (<DeleteFont/>),
	},
	"newfont": {
		level: 2+2,
		name: "New Font",
		component: (<NewFont />),
	},  
	"inituploadfont": {
		level: 2+2,
		name: "Upload Font",
		component: (<InitUploadFont />),
	},
	"inituploadfonts": {
		level: 2+2,
		name: "Upload Fonts",
		component: (<InitUploadFonts />),
	},
	"editfont": {
		level: -1,
		name: "Edit Font", 
		component: (<EditFont />),
	},  
	"uploadfont": {
		level: 2+3,
		name: "Upload File", 
		component: (<UploadFont />),
	},  
	"fonts": {
		level: 2+1,
		name: "Manage Fonts", 
		component: (<Fonts />),
		subsublinks: ["newfont"],
	},
	"font": {
		level: 2+2,
		name: "Font",
		sublinks: ["editfont", "deletefont"],
		subsublinks: [""],
		component: (<Font />),
	},
}
