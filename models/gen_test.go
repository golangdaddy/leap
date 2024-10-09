package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGen(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(TEST_1, TEST_1)
}

var TEST_1 string = `{"siteName":"Rescue Centre","Config":{"webAPI":"https://newtown.vercel.app/","hostAPI":"https://server-go-gen-test-da7z6jf32a-nw.a.run.app/","websocketHost":"server-go-gen-test-da7z6jf32a-nw.a.run.app","repoURI":"github.com/golangdaddy/newtown","projectID":"npg-generic","projectName":"go-gen-test","projectRegion":"europe-west2-b"},"objects":[{"name":"animal","names":["NAME"],"plural":"animals","json":"","context":"Define the main object for storing information about each rescued animal","children":[{"name":"healthCheckup","names":null,"plural":"checkups","json":"","context":"A record of each health checkup per animal, detailing health-related observations","parents":["animal"],"fields":null,"inputs":[{"id":"NOTES","context":"notes about the animal's health checkup","name":"notes","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":10000},"regexp":"","regexpHex":""}],"listMode":"","options":{"readonly":false,"admin":false,"member":null,"job":false,"comment":false,"order":false,"file":false,"image":false,"photo":false,"exif":false,"font":false,"topicCreate":null,"topics":null,"assetlayer":null,"handcash":{"Type":"","Payments":null,"Mint":null},"pusher":false,"permissions":{"AdminsOnly":false,"AdminsEdit":false},"filterFields":null},"tags":null,"childTags":null}],"fields":[{"id":"NAME","context":"The name of the animal","name":"name","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":30},"regexp":"","regexpHex":""},{"id":"SPECIES","context":"The species of the animal","name":"species","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":30},"regexp":"","regexpHex":""},{"id":"AGE","context":"The age of the animal","name":"age","type":"uint","element":{"Name":"INT","Go":"int","Input":"input","Type":"number"},"inputReference":"","required":true,"filter":false,"range":{"min":0,"max":-1},"regexp":"","regexpHex":""},{"id":"BIRTHDAY","context":"The D.O.B. of the animal","name":"birthday","type":"date","element":{"Name":"DATE","Go":"string","Input":"input","Type":"date"},"inputReference":"","required":true,"filter":false,"regexp":"","regexpHex":""},{"id":"ADDRESS","context":"The D.O.B. of the animal","name":"address","type":"address","element":null,"inputs":[{"id":"BUILDING_NUMBER","context":"the number of the building on the street","name":"building number","type":"int","element":{"Name":"INT","Go":"int","Input":"input","Type":"number"},"inputReference":"","required":true,"filter":false,"regexp":"","regexpHex":""},{"id":"APARTMENT_NUMBER","context":"if applicable, the number of the unit or apartment in the building","name":"apartment number","type":"int","element":{"Name":"INT","Go":"int","Input":"input","Type":"number"},"inputReference":"","required":false,"filter":false,"regexp":"","regexpHex":""},{"id":"STREET","context":"the street where the building is","name":"street","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":75},"regexp":"","regexpHex":""},{"id":"TOWN_OR_CITY","context":"the town or city where the street is","name":"town or city","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":50},"regexp":"","regexpHex":""},{"id":"COUNTRY","context":"the country where the town or city is","name":"country","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":50},"regexp":"","regexpHex":""}],"inputReference":"","required":true,"filter":false,"regexp":"","regexpHex":""}],"inputs":[{"id":"NAME","context":"The name of the animal","name":"name","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":30},"regexp":"","regexpHex":""},{"id":"SPECIES","context":"The species of the animal","name":"species","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":30},"regexp":"","regexpHex":""},{"id":"AGE","context":"The age of the animal","name":"age","type":"uint","element":{"Name":"INT","Go":"int","Input":"input","Type":"number"},"inputReference":"","required":true,"filter":false,"range":{"min":0,"max":-1},"regexp":"","regexpHex":""},{"id":"BIRTHDAY","context":"The D.O.B. of the animal","name":"birthday","type":"date","element":{"Name":"DATE","Go":"string","Input":"input","Type":"date"},"inputReference":"","required":true,"filter":false,"regexp":"","regexpHex":""},{"id":"BUILDING_NUMBER","context":"the number of the building on the street","name":"building number","type":"int","element":{"Name":"INT","Go":"int","Input":"input","Type":"number"},"inputReference":"","required":true,"filter":false,"regexp":"","regexpHex":""},{"id":"APARTMENT_NUMBER","context":"if applicable, the number of the unit or apartment in the building","name":"apartment number","type":"int","element":{"Name":"INT","Go":"int","Input":"input","Type":"number"},"inputReference":"","required":false,"filter":false,"regexp":"","regexpHex":""},{"id":"STREET","context":"the street where the building is","name":"street","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":75},"regexp":"","regexpHex":""},{"id":"TOWN_OR_CITY","context":"the town or city where the street is","name":"town or city","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":50},"regexp":"","regexpHex":""},{"id":"COUNTRY","context":"the country where the town or city is","name":"country","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":50},"regexp":"","regexpHex":""}],"listMode":"","options":{"readonly":false,"admin":true,"member":null,"job":false,"comment":false,"order":false,"file":false,"image":false,"photo":false,"exif":false,"font":false,"topicCreate":null,"topics":null,"assetlayer":null,"handcash":{"Type":"","Payments":null,"Mint":null},"pusher":false,"permissions":{"AdminsOnly":false,"AdminsEdit":false},"filterFields":null},"tags":null,"childTags":null},{"name":"healthCheckup","names":null,"plural":"checkups","json":"","context":"A record of each health checkup per animal, detailing health-related observations","parents":["animal"],"parentCount":1,"fields":[{"id":"NOTES","context":"notes about the animal's health checkup","name":"notes","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":10000},"regexp":"","regexpHex":""}],"inputs":[{"id":"NOTES","context":"notes about the animal's health checkup","name":"notes","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":10000},"regexp":"","regexpHex":""}],"listMode":"","options":{"readonly":false,"admin":false,"member":null,"job":false,"comment":false,"order":false,"file":false,"image":false,"photo":false,"exif":false,"font":false,"topicCreate":null,"topics":null,"assetlayer":null,"handcash":{"Type":"","Payments":null,"Mint":null},"pusher":false,"permissions":{"AdminsOnly":false,"AdminsEdit":false},"filterFields":null},"tags":null,"childTags":null},{"name":"adopter","names":null,"plural":"adopters","json":"","context":"Stores information about individuals who adopt animals","fields":[{"id":"PERSON.NAME","context":"The name of the adopter","name":"person.name","type":"person.name","element":null,"inputs":[{"id":"FIRST_NAME","context":"A name or names of something or someone","name":"first-name","type":"name","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":2,"max":50},"regexp":"","regexpHex":""},{"id":"MIDDLE_NAMES","context":"A name or names of something or someone","name":"middle-names","type":"name","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":false,"filter":false,"range":{"min":2,"max":50},"regexp":"","regexpHex":""},{"id":"LAST_NAME","context":"A name or names of something or someone","name":"last-name","type":"name","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":2,"max":50},"regexp":"","regexpHex":""}],"inputReference":"","required":true,"filter":false,"regexp":"","regexpHex":""},{"id":"PHONE_NUMBER","context":"The phone number of the adopter","name":"phone number","type":"phone","element":{"Name":"PHONE","Go":"string","Input":"input","Type":"tel"},"inputReference":"","required":true,"filter":false,"regexp":"^\\+?[1-9]\\d{1,14}$","regexpHex":"5e5c2b3f5b312d395d5c647b312c31347d24"}],"inputs":[{"id":"FIRST_NAME","context":"A name or names of something or someone","name":"first-name","type":"name","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":2,"max":50},"regexp":"","regexpHex":""},{"id":"MIDDLE_NAMES","context":"A name or names of something or someone","name":"middle-names","type":"name","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":false,"filter":false,"range":{"min":2,"max":50},"regexp":"","regexpHex":""},{"id":"LAST_NAME","context":"A name or names of something or someone","name":"last-name","type":"name","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":2,"max":50},"regexp":"","regexpHex":""},{"id":"PHONE_NUMBER","context":"The phone number of the adopter","name":"phone number","type":"phone","element":{"Name":"PHONE","Go":"string","Input":"input","Type":"tel"},"inputReference":"","required":true,"filter":false,"regexp":"^\\+?[1-9]\\d{1,14}$","regexpHex":"5e5c2b3f5b312d395d5c647b312c31347d24"}],"listMode":"","options":{"readonly":false,"admin":true,"member":null,"job":false,"comment":false,"order":false,"file":false,"image":false,"photo":false,"exif":false,"font":false,"topicCreate":null,"topics":null,"assetlayer":null,"handcash":{"Type":"","Payments":null,"Mint":null},"pusher":false,"permissions":{"AdminsOnly":false,"AdminsEdit":false},"filterFields":null},"tags":null,"childTags":null}],"entrypoints":[{"name":"animal","names":["NAME"],"plural":"animals","json":"","context":"Define the main object for storing information about each rescued animal","children":[{"name":"healthCheckup","names":null,"plural":"checkups","json":"","context":"A record of each health checkup per animal, detailing health-related observations","parents":["animal"],"fields":null,"inputs":[{"id":"NOTES","context":"notes about the animal's health checkup","name":"notes","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":10000},"regexp":"","regexpHex":""}],"listMode":"","options":{"readonly":false,"admin":false,"member":null,"job":false,"comment":false,"order":false,"file":false,"image":false,"photo":false,"exif":false,"font":false,"topicCreate":null,"topics":null,"assetlayer":null,"handcash":{"Type":"","Payments":null,"Mint":null},"pusher":false,"permissions":{"AdminsOnly":false,"AdminsEdit":false},"filterFields":null},"tags":null,"childTags":null}],"fields":[{"id":"NAME","context":"The name of the animal","name":"name","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":30},"regexp":"","regexpHex":""},{"id":"SPECIES","context":"The species of the animal","name":"species","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":30},"regexp":"","regexpHex":""},{"id":"AGE","context":"The age of the animal","name":"age","type":"uint","element":{"Name":"INT","Go":"int","Input":"input","Type":"number"},"inputReference":"","required":true,"filter":false,"range":{"min":0,"max":-1},"regexp":"","regexpHex":""},{"id":"BIRTHDAY","context":"The D.O.B. of the animal","name":"birthday","type":"date","element":{"Name":"DATE","Go":"string","Input":"input","Type":"date"},"inputReference":"","required":true,"filter":false,"regexp":"","regexpHex":""},{"id":"ADDRESS","context":"The D.O.B. of the animal","name":"address","type":"address","element":null,"inputs":[{"id":"BUILDING_NUMBER","context":"the number of the building on the street","name":"building number","type":"int","element":{"Name":"INT","Go":"int","Input":"input","Type":"number"},"inputReference":"","required":true,"filter":false,"regexp":"","regexpHex":""},{"id":"APARTMENT_NUMBER","context":"if applicable, the number of the unit or apartment in the building","name":"apartment number","type":"int","element":{"Name":"INT","Go":"int","Input":"input","Type":"number"},"inputReference":"","required":false,"filter":false,"regexp":"","regexpHex":""},{"id":"STREET","context":"the street where the building is","name":"street","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":75},"regexp":"","regexpHex":""},{"id":"TOWN_OR_CITY","context":"the town or city where the street is","name":"town or city","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":50},"regexp":"","regexpHex":""},{"id":"COUNTRY","context":"the country where the town or city is","name":"country","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":50},"regexp":"","regexpHex":""}],"inputReference":"","required":true,"filter":false,"regexp":"","regexpHex":""}],"inputs":[{"id":"NAME","context":"The name of the animal","name":"name","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":30},"regexp":"","regexpHex":""},{"id":"SPECIES","context":"The species of the animal","name":"species","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":30},"regexp":"","regexpHex":""},{"id":"AGE","context":"The age of the animal","name":"age","type":"uint","element":{"Name":"INT","Go":"int","Input":"input","Type":"number"},"inputReference":"","required":true,"filter":false,"range":{"min":0,"max":-1},"regexp":"","regexpHex":""},{"id":"BIRTHDAY","context":"The D.O.B. of the animal","name":"birthday","type":"date","element":{"Name":"DATE","Go":"string","Input":"input","Type":"date"},"inputReference":"","required":true,"filter":false,"regexp":"","regexpHex":""},{"id":"BUILDING_NUMBER","context":"the number of the building on the street","name":"building number","type":"int","element":{"Name":"INT","Go":"int","Input":"input","Type":"number"},"inputReference":"","required":true,"filter":false,"regexp":"","regexpHex":""},{"id":"APARTMENT_NUMBER","context":"if applicable, the number of the unit or apartment in the building","name":"apartment number","type":"int","element":{"Name":"INT","Go":"int","Input":"input","Type":"number"},"inputReference":"","required":false,"filter":false,"regexp":"","regexpHex":""},{"id":"STREET","context":"the street where the building is","name":"street","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":75},"regexp":"","regexpHex":""},{"id":"TOWN_OR_CITY","context":"the town or city where the street is","name":"town or city","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":50},"regexp":"","regexpHex":""},{"id":"COUNTRY","context":"the country where the town or city is","name":"country","type":"string","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":1,"max":50},"regexp":"","regexpHex":""}],"listMode":"","options":{"readonly":false,"admin":true,"member":null,"job":false,"comment":false,"order":false,"file":false,"image":false,"photo":false,"exif":false,"font":false,"topicCreate":null,"topics":null,"assetlayer":null,"handcash":{"Type":"","Payments":null,"Mint":null},"pusher":false,"permissions":{"AdminsOnly":false,"AdminsEdit":false},"filterFields":null},"tags":null,"childTags":null},{"name":"adopter","names":null,"plural":"adopters","json":"","context":"Stores information about individuals who adopt animals","fields":[{"id":"PERSON.NAME","context":"The name of the adopter","name":"person.name","type":"person.name","element":null,"inputs":[{"id":"FIRST_NAME","context":"A name or names of something or someone","name":"first-name","type":"name","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":2,"max":50},"regexp":"","regexpHex":""},{"id":"MIDDLE_NAMES","context":"A name or names of something or someone","name":"middle-names","type":"name","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":false,"filter":false,"range":{"min":2,"max":50},"regexp":"","regexpHex":""},{"id":"LAST_NAME","context":"A name or names of something or someone","name":"last-name","type":"name","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":2,"max":50},"regexp":"","regexpHex":""}],"inputReference":"","required":true,"filter":false,"regexp":"","regexpHex":""},{"id":"PHONE_NUMBER","context":"The phone number of the adopter","name":"phone number","type":"phone","element":{"Name":"PHONE","Go":"string","Input":"input","Type":"tel"},"inputReference":"","required":true,"filter":false,"regexp":"^\\+?[1-9]\\d{1,14}$","regexpHex":"5e5c2b3f5b312d395d5c647b312c31347d24"}],"inputs":[{"id":"FIRST_NAME","context":"A name or names of something or someone","name":"first-name","type":"name","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":2,"max":50},"regexp":"","regexpHex":""},{"id":"MIDDLE_NAMES","context":"A name or names of something or someone","name":"middle-names","type":"name","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":false,"filter":false,"range":{"min":2,"max":50},"regexp":"","regexpHex":""},{"id":"LAST_NAME","context":"A name or names of something or someone","name":"last-name","type":"name","element":{"Name":"STRING","Go":"string","Input":"input","Type":"text"},"inputReference":"","required":true,"filter":false,"range":{"min":2,"max":50},"regexp":"","regexpHex":""},{"id":"PHONE_NUMBER","context":"The phone number of the adopter","name":"phone number","type":"phone","element":{"Name":"PHONE","Go":"string","Input":"input","Type":"tel"},"inputReference":"","required":true,"filter":false,"regexp":"^\\+?[1-9]\\d{1,14}$","regexpHex":"5e5c2b3f5b312d395d5c647b312c31347d24"}],"listMode":"","options":{"readonly":false,"admin":true,"member":null,"job":false,"comment":false,"order":false,"file":false,"image":false,"photo":false,"exif":false,"font":false,"topicCreate":null,"topics":null,"assetlayer":null,"handcash":{"Type":"","Payments":null,"Mint":null},"pusher":false,"permissions":{"AdminsOnly":false,"AdminsEdit":false},"filterFields":null},"tags":null,"childTags":null}],"options":{"sidebar":false,"chatgpt":true,"assetlayer":false,"wallets":null,"whitelistDomains":false,"registrationDomains":null,"whitelistEmails":false,"registrationEmails":null,"pusher":null,"handcash":{"appId":"660c209b9295c1bcf6312def","appSecret":"7b7489072ece66e7f93867ba6ff638a1f80943ebb51629e6bfc6b17d85dbb1b1"}}}`