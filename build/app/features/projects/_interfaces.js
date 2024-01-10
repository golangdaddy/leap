import * as React from 'react'

import { Projects } from './projects'
import { Project } from './project'
import { NewProject } from './newProject'
import { EditProject } from './editProject'
import { UploadProject } from './uploadProject'
import { DeleteProject } from './deleteProject'
import { InitUploadProject } from './initUploadProject'
import { InitUploadProjects } from './initUploadProjects'

export var ProjectInterfaces = {
	"deleteproject": {
		level: -1,
		name: "Delete", 
		component: (<DeleteProject/>),
	},
	"newproject": {
		level: 0+2,
		name: "New Project",
		component: (<NewProject />),
	},  
	"inituploadproject": {
		level: 0+2,
		name: "Upload Project",
		component: (<InitUploadProject />),
	},
	"inituploadprojects": {
		level: 0+2,
		name: "Upload Projects",
		component: (<InitUploadProjects />),
	},
	"editproject": {
		level: -1,
		name: "Edit Project", 
		component: (<EditProject />),
	},  
	"uploadproject": {
		level: 0+3,
		name: "Upload File", 
		component: (<UploadProject />),
	},  
	"projects": {
		level: 0+1,
		name: "Manage Projects", 
		component: (<Projects />),
		subsublinks: ["newproject"],
	},
	"project": {
		level: 0+2,
		name: "Project",
		sublinks: ["editproject", "deleteproject"],
		subsublinks: ["collections","fonts",""],
		component: (<Project />),
	},
}
