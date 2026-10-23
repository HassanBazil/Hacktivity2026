package main

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"

	validator "github.com/go-playground/validator/v10"

	StockJSON "encoding/json"
	StockJSON2 "encoding/json/v2"

	BugerJSON "github.com/buger/jsonparser"
	GoJay "github.com/francoispqt/gojay"
	goccyJSON "github.com/goccy/go-json"
	JSONiter "github.com/json-iterator/go"
	SegmentIoJSON "github.com/segmentio/encoding/json"
	//SIMD doesn't work on apple silicon
	//jettison doesn't decode
	//easyJSON is a parser code generator
)

const NumberOfParsers = 8

//go:generate ffjson $GOFILE

// Yes you can use regex, but should you?
type User struct {
	Name   string  `validate:"required" json:"name"`
	YoE    int     `validate:"required,gte=0,lte=50" json:"yoe"`
	Email  string  `validate:"required,email" json:"email"`
	Avatar string  `validate:"required" json:"path"`
	Vibe   float32 `validate:"required" json:"vibe"`
}

func main() {

	users := make([]User, NumberOfParsers, NumberOfParsers)

	//🌍 Origin - ip address, access key (token)
	input := []byte(exampleInput)

	fmt.Println("JSON processed: \n", string(input))

	// 📏 Size - Is message reasonably long? Are all fields reasonably long as well?
	if len(input) > 1000 {
		fmt.Println("❗️ Input is larger than should be.")
	} else {
		fmt.Println("✅ Input is not too long.")
	}

	// 📚 Lexical validation - Content uses right characters and encoding?
	if !utf8.Valid(input) {
		fmt.Println("❗️ Incorrect character encoding!")
	} else {
		fmt.Println("✅ Character encoding is okay!")
	}

	wrongClass := false
	for _, runeValue := range string(input) {
		if !(unicode.IsNumber(runeValue) || unicode.IsLetter(runeValue) || unicode.IsPunct(runeValue) || unicode.IsSpace(runeValue)) && !unicode.Is(unicode.Latin, runeValue) {
			fmt.Println("❗️ Incorrect character class!", string(runeValue))
			wrongClass = true
		}
	}
	if !wrongClass {
		fmt.Println("✅ Character class is okay!")
	}

	//🪭 Normalise unicode bytes using composing representation for more compact representation. Consider
	//input = norm.NFC.Bytes(input)

	// 📝 Syntax validation - Is the format is right? (JSON double property etc. )
	fmt.Println("")
	if !StockJSON.Valid([]byte(input)) {
		fmt.Println("❗️ Incorrect syntax by stock JSON.")
	} else {
		fmt.Println("✅ JSON is valid.")
	}

	fmt.Println("")

	// ⚙️ Parsing
	var parsers = map[int]string{
		0: "Stock",
		1: "goccy",
		2: "JSONiter",
		3: "JSONiter compatible mode",
		4: "GoJay",
		5: "SegmentIOJSON",
		6: "BugerJSON",
		7: "StockJSON2",
	}

	StockJSON.Unmarshal(input, &users[0])

	goccyJSON.Unmarshal(input, &users[1])

	JSONiter.Unmarshal(input, &users[2])

	var JSONIterCompatibilityMode = JSONiter.ConfigCompatibleWithStandardLibrary
	JSONIterCompatibilityMode.Unmarshal(input, &users[3])

	GoJay.UnmarshalJSONObject(input, &users[4])

	SegmentIoJSON.Unmarshal(input, &users[5])

	// Buger JSON
	users[6].Name, _ = BugerJSON.Lenient.GetString(input, "Name")
	a, _, _, _ := BugerJSON.Lenient.Get(input, "yoe")
	users[6].YoE, _ = strconv.Atoi(string(a))
	users[6].Email, _ = BugerJSON.Lenient.GetString(input, "eMail")
	users[6].Avatar, _ = BugerJSON.Lenient.GetString(input, "path")
	b, _, _, _ := BugerJSON.Lenient.Get(input, "vibe")
	c, _ := strconv.ParseFloat(string(b), 32)
	users[6].Vibe = float32(c)

	StockJSON2.Unmarshal(input, &users[7])

	list := []int{0, 6, 7}
	for _, i := range list {
		fmt.Println(parsers[i], users[i])
	}

	fmt.Println("")

	// 🧠 Semantic validation - Does it make sense?
	fmt.Println("Let's validate semantics!")

	validate := validator.New()
	for _, i := range list {
		err := validate.Struct(users[i])
		if err != nil {
			fmt.Println("❗️Using ", parsers[i], " parser struct did not pass validation.")
		} else {
			fmt.Println("✅ Using", parsers[i], "  parser struct did passed validation.")
		}
		//fmt.Println("According to ", parsers[i], "user looks like this :\n", users[i], "\n")
	}

	fmt.Println("")

	for _, i := range list {
		if users[i].Name != "Pérez Zoé" {
			fmt.Println("Not the Zoé", parsers[i], "expected. This is ", users[i].Name, ".")
		} else {
			fmt.Println(" Hello Zoé", parsers[i], "expected you.")
		}
	}
}
