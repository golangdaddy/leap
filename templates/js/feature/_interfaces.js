import * as React from 'react'

import { {{titlecase .Object.Name}}s } from './{{lowercase .Object.Name}}s'
import { {{titlecase .Object.Name}} } from './{{lowercase .Object.Name}}'
import { New{{titlecase .Object.Name}} } from './new{{titlecase .Object.Name}}'
import { Edit{{titlecase .Object.Name}} } from './edit{{titlecase .Object.Name}}'

export var {{titlecase .Object.Name}}Interfaces = {
    "new{{lowercase .Object.Name}}": {
        level: -1,
        name: "New {{titlecase .Object.Name}}", 
        component: (<New{{titlecase .Object.Name}} />),
    },  
    "edit{{lowercase .Object.Name}}": {
        level: -1,
        name: "Edit {{titlecase .Object.Name}}", 
        component: (<Edit{{titlecase .Object.Name}} />),
    },  
    "{{lowercase .Object.Name}}s": {
        level: 7,
        name: "Manage {{titlecase .Object.Name}}s", 
        component: (<{{titlecase .Object.Name}}s />),
    },
    "{{lowercase .Object.Name}}": {
        level: 8,
        name: "{{titlecase .Object.Name}}", 
        component: (<{{titlecase .Object.Name}} />),
    },
}
