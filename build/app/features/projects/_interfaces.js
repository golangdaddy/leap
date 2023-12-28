import * as React from 'react'

import { Projects } from './projects'
import { Project } from './project'
import { NewProject } from './newProject'
import { EditProject } from './editProject'

export var ProjectInterfaces = {
    "newproject": {
        level: -1,
        name: "New Project", 
        component: (<NewProject />),
    },  
    "editproject": {
        level: -1,
        name: "Edit Project", 
        component: (<EditProject />),
    },  
    "projects": {
        level: 7,
        name: "Manage Projects", 
        component: (<Projects />),
    },
    "project": {
        level: 8,
        name: "Project", 
        component: (<Project />),
    },
}
