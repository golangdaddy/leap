import * as React from 'react'

import { {{titlecase .Object.Name}}sMatrix } from './{{lowercase .Object.Name}}sMatrix'
import { {{titlecase .Object.Name}}s } from './{{lowercase .Object.Name}}s'
import { {{titlecase .Object.Name}} } from './{{lowercase .Object.Name}}'
import { {{titlecase .Object.Name}}Admin } from './{{lowercase .Object.Name}}Admin'
import { {{titlecase .Object.Name}}Admins } from './{{lowercase .Object.Name}}Admins'
import { Assets } from './assets'
import { New{{titlecase .Object.Name}} } from './new{{titlecase .Object.Name}}'
import { Edit{{titlecase .Object.Name}} } from './edit{{titlecase .Object.Name}}'
import { Delete{{titlecase .Object.Name}} } from './delete{{titlecase .Object.Name}}'
import { InitUpload{{titlecase .Object.Name}} } from './initUpload{{titlecase .Object.Name}}'
import { InitUpload{{titlecase .Object.Name}}s } from './initUpload{{titlecase .Object.Name}}s'
import { Upload{{titlecase .Object.Name}} } from './upload{{titlecase .Object.Name}}'

export var {{titlecase .Object.Name}}Interfaces = {
	"delete{{lowercase .Object.Name}}": {
		level: -1,
		name: "Delete", 
		component: (<Delete{{titlecase .Object.Name}}/>),
	},
	"new{{lowercase .Object.Name}}": {
		level: {{parentcount .Object}}+2,
		name: "New {{titlecase .Object.Name}}",
		component: (<New{{titlecase .Object.Name}} />),
	},
	{{if .Object.Options.File}}
	"initupload{{lowercase .Object.Name}}": {
		level: {{parentcount .Object}}+2,
		name: "Upload {{titlecase .Object.Name}}",
		component: (<InitUpload{{titlecase .Object.Name}} />),
	},
	"initupload{{lowercase .Object.Name}}s": {
		level: {{parentcount .Object}}+2,
		name: "Upload {{titlecase .Object.Name}}s",
		component: (<InitUpload{{titlecase .Object.Name}}s />),
	},
	"upload{{lowercase .Object.Name}}": {
		level: {{parentcount .Object}}+3,
		name: "Upload File", 
		component: (<Upload{{titlecase .Object.Name}} />),
	},
	{{end}}
	"edit{{lowercase .Object.Name}}": {
		level: -1,
		name: "Edit {{titlecase .Object.Name}}", 
		component: (<Edit{{titlecase .Object.Name}} />),
	},  
	"{{lowercase .Object.Name}}s": {
		level: {{parentcount .Object}}+1,
		name: "{{titlecase .Object.Name}}s", 
		component: (<{{titlecase .Object.Name}}s />),
		subsublinks: ["{{lowercase .Object.Name}}smatrix", "new{{lowercase .Object.Name}}"{{if .Object.Options.File}}, "initupload{{lowercase .Object.Name}}", "initupload{{lowercase .Object.Name}}s"{{end}}],
		hasNewButton: true,
		hasSpreadsheetButton: true,
	},
	"{{lowercase .Object.Name}}smatrix": {
		level: {{parentcount .Object}}+2,
		name: "{{titlecase .Object.Name}}s Matrix", 
		component: (<{{titlecase .Object.Name}}sMatrix />),
		subsublinks: ["new{{lowercase .Object.Name}}"{{if .Object.Options.File}}, "initupload{{lowercase .Object.Name}}", "initupload{{lowercase .Object.Name}}s"{{end}}],
		hasNewButton: true,
		hasListButton: true,
	},
	"{{lowercase .Object.Name}}": {
		level: {{parentcount .Object}}+2,
		name: "{{titlecase .Object.Name}}",
		subsublinks: [{{ range .Object.Children }}"{{lowercase .Name}}s",{{end}}{{if .Object.Options.Admin}}"{{lowercase .Object.Name}}admins"{{end}}],
		component: (<{{titlecase .Object.Name}} />),
		hasDeleteButton: true,
		hasEditButton: true,
	},
	"{{lowercase .Object.Name}}admin": {
		level: {{parentcount .Object}}+2,
		name: "Admin",
		component: (<{{titlecase .Object.Name}}Admin />),
	},
	"{{lowercase .Object.Name}}admins": {
		level: {{parentcount .Object}}+2,
		name: "Admins",
		component: (<{{titlecase .Object.Name}}Admins />),
	},
	"{{lowercase .Object.Name}}assets": {
		level: {{parentcount .Object}}+2,
		name: "{{titlecase .Object.Name}} Assets",
		component: (<Assets />),
	},
}
