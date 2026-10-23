package main

const exampleInput = `{
	"name": "Pérez Zoé",
	"Name": "Pérez Zoé",
	"yoe": 127,
	"YoE": 27,
	"email": "Bazil@🍕❤️.com",
	"eMail": "!!!!john.smith@example.com",
	"Email": "Bazil@🍕❤️.ws",
	"path":"/mypicture.png",
	"Path":"../../../../etc/shadow",
	"vibe":9999999999999999999999999999999999999999999999999999999
	"isAdmin": true
}/* this is a comment */`

// Emoji in email address is a thing https://mailoji.com/

// note that the Name and Age fields fail the validation

const exampleInput3 = `{
	"Name": "Pérez Zoé2",
	"name": "Pérez Zoé",
	"Name": "John Doe",
	"YoE": 27,
	"Email": "john@@@xuple.com",
	"Email": "john.smith@example.com",
	"path":"/admin\u0000/../public"
}`

const exampleInputfunk = `{
	"Name": "Pérez Zoé2",
	"name": "Pérez Zoé",
	"NaMe": "John Doe",
	"YoE": 27,
	"yoe": 28,
	"EmAil": "john@@@xuple.com",
	"Email": "john.smith@example.com",
	"Path":"/admin\u0000/../public"
	"path":"/mypicture.png"
}/* this is a comment */`

const charTruncationExample = `
{"test": 1, "test\[raw \x0d byte]": 2} 
{"test": 1, "test\ud800": 2}
{"test": 1, "test"": 2}
{"test": 1, "te\st": 2}
`

// 	"Age" : 4569,
// 	"Age": 26,
// 	"Age"": 29,
// 	"Ag\e": 30,
// 	"Age\[raw \x0d byte]": 27,
// 	"Age\ud800": 28,

// // Emoji in email address is a thing https://mailoji.com/

// // note that the Name and Age fields fail the validation

// const exampleInput2 = {

// "name": "John Doe2"

// "nAmE": "John Doe",

// "Age": 26,

// "Age": "4566789",

// "Email": "john.smith@example.com",

// "Email"："johnc@axmple.com"

// // Emoji in email address is a thing https://mailoji.com/

// // note that the Name and Age fields fail the validation

// const exampleInput3 = {

// "name":

// "John Doe2",

// "nAmE": "John Doe",

// "Age": 26,

// "Age": "268",

// "Email": "john@@@xmple.com"

// "Email": "john.smith@example.com"

// const slideExample = {

// "name": "John Doe2"

// "name": "John Doe",

// "Age": 26,

// "Age": "4566789"

// "Email": "john@@@xmple.com",

// "Email": "john.smith@example.com"

// const slideExample2 =

// "name": "John Doe2"

// "name": "John Doe",

// "Age": 26,

// "Age": "4566789"

// "Email": "john.smith@example.com"

// "Email": "john@@@xmple.com"

// const

// slideExample3

// "name": "John Doe",

// "Age": 26,

// "Email": "john@@@xmple.com"

const slideExample4 = `{
"name":"Pérez Zoé",
"age": 26,
"email": "john🐶@💰xmple.com"
}`

// é non combined
// é combined
const slideExample5 = `{
"name":"Pérez Zoé2",
"nAmE": "Pérez Zoé",
"Email": "john🐶@💰xmple.com"
"email": "john.smith@example.com"
"agE": "4566789",
"age": 26
}`
