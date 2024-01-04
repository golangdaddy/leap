import * as React from 'react'

import { {{titlecase .Object.Name}}s } from './{{lowercase .Object.Name}}s'
import { {{titlecase .Object.Name}} } from './{{lowercase .Object.Name}}'
import { New{{titlecase .Object.Name}} } from './new{{titlecase .Object.Name}}'
import { Edit{{titlecase .Object.Name}} } from './edit{{titlecase .Object.Name}}'
import { Upload{{titlecase .Object.Name}} } from './upload{{titlecase .Object.Name}}'

export var {{titlecase .Object.Name}}Interfaces = {
    "new{{lowercase .Object.Name}}": {
        level: 2,
        name: "New {{titlecase .Object.Name}}", 
        component: (<New{{titlecase .Object.Name}} />),
    },  
    "edit{{lowercase .Object.Name}}": {
        level: 3,
        name: "Edit {{titlecase .Object.Name}}", 
        component: (<Edit{{titlecase .Object.Name}} />),
    },  
    "upload{{lowercase .Object.Name}}": {
        level: 3,
        name: "Upload File", 
        component: (<Upload{{titlecase .Object.Name}} />),
    },  
    "{{lowercase .Object.Name}}s": {
        level: 1,
        name: "Manage {{titlecase .Object.Name}}s", 
        component: (<{{titlecase .Object.Name}}s />),
        subsublinks: ["new{{lowercase .Object.Name}}"],
    },
    "{{lowercase .Object.Name}}": {
        level: 2,
        name: "{{titlecase .Object.Name}}",
        sublinks: ["edit{{lowercase .Object.Name}}", "upload{{lowercase .Object.Name}}"],
        subsublinks: [{{ range .Object.Children }}"{{lowercase .Name}}s",{{end}}""],
        component: (<{{titlecase .Object.Name}} />),
    },
}
