import * as React from 'react'

import { Projects } from './projects'
import { Project } from './project'
import { NewProject } from './newProject'
import { EditProject } from './editProject'

export var ProjectInterfaces = {
    "newproject": {
        level: 2,
        name: "New Project", 
        component: (<NewProject />),
    },  
    "editproject": {
        level: 3,
        name: "Edit Project", 
        component: (<EditProject />),
    },  
    "projects": {
        level: 1,
        name: "Manage Projects", 
        component: (<Projects />),
        subsublinks: ["newproject"],
    },
    "project": {
        level: 2,
        name: "Project", 
        subsublinks: ["collections",""],
        component: (<Project />),
    },
}
