import * as React from 'react'

import { DeleteObject } from '@/inputs/deleteObject'
import { ImageUpload } from '@/components/imageUpload'
import { ArchiveUpload } from '@/components/archiveUpload'
import { FontUpload } from '@/components/fontUpload'


import { ProjectInterfaces } from '@/features/projects/_interfaces'
import { CollectionInterfaces } from '@/features/collections/_interfaces'
import { FontInterfaces } from '@/features/fonts/_interfaces'
import { AttributeInterfaces } from '@/features/attributes/_interfaces'
import { LayerInterfaces } from '@/features/layers/_interfaces'
import { ElementInterfaces } from '@/features/elements/_interfaces'

import Home from '@/features/home'

export var Interfaces = {
	"home": {
		level: 0,
		name: "Home",
		component: (<Home/>),
	},
	"deleteobject": {
		level: -1,
		name: "Delete", 
		component: (<DeleteObject/>),
	},
	"imageupload": {
		level: -1,
		name: "Image Upload", 
		component: (<ImageUpload/>),
	},
	"archiveupload": {
		level: -1,
		name: "Archive Upload", 
		component: (<ArchiveUpload/>),
	},
	"fontupload": {
		level: -1,
		name: "Font Upload", 
		component: (<FontUpload/>),
	},
}

export function GetInterfaces() {
	var interfaces = {}
	for (const k in Interfaces) {
		interfaces[k.toLowerCase()] = Interfaces[k]
	}
	// custom features

	// ProjectInterfaces
	for (const k in ProjectInterfaces) {
		interfaces[k.toLowerCase()] = ProjectInterfaces[k]
	}// CollectionInterfaces
	for (const k in CollectionInterfaces) {
		interfaces[k.toLowerCase()] = CollectionInterfaces[k]
	}// FontInterfaces
	for (const k in FontInterfaces) {
		interfaces[k.toLowerCase()] = FontInterfaces[k]
	}// AttributeInterfaces
	for (const k in AttributeInterfaces) {
		interfaces[k.toLowerCase()] = AttributeInterfaces[k]
	}// LayerInterfaces
	for (const k in LayerInterfaces) {
		interfaces[k.toLowerCase()] = LayerInterfaces[k]
	}// ElementInterfaces
	for (const k in ElementInterfaces) {
		interfaces[k.toLowerCase()] = ElementInterfaces[k]
	}
	
	// put id key into the object
	for (const k in interfaces) {
		interfaces[k].id = k
		for (const sub in interfaces[k].subLinks) {
			interfaces[k].subLinks[sub] = interfaces[k].subLinks[sub].toLowerCase()
		}
	}
	console.log("INTERFACES", interfaces)
	return interfaces
}

export function GoBack(localdata) {
	const previousTab = localdata.breadcrumbs[localdata.breadcrumbs.length - 2]
	return VisitTab(localdata, previousTab.id, previousTab.context)
}

export function GoBackBack(localdata) {
	const previousTab = localdata.breadcrumbs[localdata.breadcrumbs.length - 3]
	return VisitTab(localdata, previousTab.id, previousTab.context)
}

// update tabs handles the updated context and sends the user to a new interface
export default function VisitTab(localdata, tabname, context) {

	if (!context) {
		context = {}
	}

	const home = Object.assign({}, GetInterfaces()["home"])

	if (!localdata) {
		localdata = {
			"tab": home, 
			"region": "UK",
			"breadcrumbs": [home],
			"context": {}
		}
	}

	console.log("VISIT TAB", tabname)
	var crumbs = [];

	var tab = Object.assign({}, GetInterfaces()[tabname])
	if (!tab) {
		console.error("TAB NOT FOUND", tabname)
		return
	}

	// assign the context into the tab
	tab.context = Object.assign({}, context)

	if (tab.context._ === localdata.breadcrumbs[localdata.breadcrumbs.length-1].context?._) {
		tab.context._ = ""
	}

	console.log("SWITCHING TAB", tab)

	switch (tab.level) {
	case 0:
		crumbs = [tab]
		break
	case 1:
		crumbs = [home, tab]
		break
	case 2:
		crumbs = [localdata.breadcrumbs[0], localdata.breadcrumbs[1], tab]
		break
	case 3:
		crumbs = [localdata.breadcrumbs[0], localdata.breadcrumbs[1], localdata.breadcrumbs[2], tab]
		break
	case 4:
		crumbs = [localdata.breadcrumbs[0], localdata.breadcrumbs[1], localdata.breadcrumbs[2], localdata.breadcrumbs[3], tab]
		break
	case 5:
		crumbs = [localdata.breadcrumbs[0], localdata.breadcrumbs[1], localdata.breadcrumbs[2], localdata.breadcrumbs[3], localdata.breadcrumbs[4], tab]
		break
	case 6:
		crumbs = [localdata.breadcrumbs[0], localdata.breadcrumbs[1], localdata.breadcrumbs[2], localdata.breadcrumbs[3], localdata.breadcrumbs[4], localdata.breadcrumbs[5], tab]
		break
	case 7:
		crumbs = [localdata.breadcrumbs[0], localdata.breadcrumbs[1], localdata.breadcrumbs[2], localdata.breadcrumbs[3], localdata.breadcrumbs[4], localdata.breadcrumbs[5], localdata.breadcrumbs[6], tab]
		break
	case 8:
		crumbs = [localdata.breadcrumbs[0], localdata.breadcrumbs[1], localdata.breadcrumbs[2], localdata.breadcrumbs[3], localdata.breadcrumbs[4], localdata.breadcrumbs[5], localdata.breadcrumbs[6], localdata.breadcrumbs[7], tab]
		break
	case 9:
		crumbs = [localdata.breadcrumbs[0], localdata.breadcrumbs[1], localdata.breadcrumbs[2], localdata.breadcrumbs[3], localdata.breadcrumbs[4], localdata.breadcrumbs[5], localdata.breadcrumbs[6], localdata.breadcrumbs[7], localdata.breadcrumbs[8], tab]
		break
	case -1:
		localdata.breadcrumbs.map(function(crumb, i) {
			crumbs.push(crumb)
		})
		crumbs.push(tab)
		break
	}
	if (!crumbs.length) {
		console.error("NO CRUMBS")
		return
	}

	var newData = {
		"tab": tab,
		"breadcrumbs": crumbs,	
		"region": localdata.region,
	}

	return newData
}