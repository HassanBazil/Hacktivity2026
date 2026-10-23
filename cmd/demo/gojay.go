package main

import (
	GoJay "github.com/francoispqt/gojay"
	//SIMD doesn't work on m1
	// jettison doesn't decode
	// easyJSON is a parser code generator
)

// implement gojay. Unmarshaler helper function
func (u *User) UnmarshalJSONObject(dec *GoJay.Decoder, key string) error {
	switch key {
	case "yoe":
		return dec.Int(&u.YoE)
	case "name":
		return dec.String(&u.Name)
	case "email":
		return dec.String(&u.Email)
	case "path":
		return dec.String(&u.Avatar)
	}
	return nil
}

func (u *User) NKeys() int {
	return 3
}
