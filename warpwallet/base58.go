// code from https://github.com/ThePiachu/Split-Vanity-Miner-Golang/blob/master/src/pkg/mymath/base58.go, on 2013-11-28

// Copyright (c) 2013 Charles M. Ellison III.  All rights reserved.
// Copyright (c) 2012 ThePiachu. All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:

//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * The name of ThePiachu may not be used to endorse or promote products
// derived from this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package warpwallet

import "math/big"

//Useful materials:
//https://en.bitcoin.it/wiki/Base_58_Encoding
//http://www.strongasanox.co.uk/2011/03/11/base58-encoding-in-python/

//alphabet used by Bitcoins
var alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

//type to hold the Base58 string
type Base58 string

//reverse alphabet used for quckly converting base58 strings into numbers
var revalp = map[string]int{
	"1": 0, "2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6, "8": 7, "9": 8, "A": 9,
	"B": 10, "C": 11, "D": 12, "E": 13, "F": 14, "G": 15, "H": 16, "J": 17, "K": 18, "L": 19,
	"M": 20, "N": 21, "P": 22, "Q": 23, "R": 24, "S": 25, "T": 26, "U": 27, "V": 28, "W": 29,
	"X": 30, "Y": 31, "Z": 32, "a": 33, "b": 34, "c": 35, "d": 36, "e": 37, "f": 38, "g": 39,
	"h": 40, "i": 41, "j": 42, "k": 43, "m": 44, "n": 45, "o": 46, "p": 47, "q": 48, "r": 49,
	"s": 50, "t": 51, "u": 52, "v": 53, "w": 54, "x": 55, "y": 56, "z": 57,
}

//Convert base58 to big.Int
func (b Base58) ToBig() *big.Int {
	answer := new(big.Int)
	for i := 0; i < len(b); i++ {
		answer.Mul(answer, big.NewInt(58))                              //multiply current value by 58
		answer.Add(answer, big.NewInt(int64(revalp[string(b[i:i+1])]))) //add value of the current letter
	}
	return answer
}

//convert base58 to hex bytes
func (b Base58) ToHex() []byte {
	value := b.ToBig() //convert to big.Int
	oneCount := 0
	for string(b)[oneCount] == '1' {
		oneCount++
	}
	return append(make([]byte, oneCount), value.Bytes()...) //convert big.Int to bytes
}

//convert base58 to hex bytes
func Base582Hex(b string) []byte {
	return Base58(b).ToHex()
}

//convert base58 to hexes used by Bitcoins (keeping the zeroes on the front, 25 bytes long)
func (b Base58) BitHex() []byte {
	value := b.ToBig() //convert to big.Int

	tmp := value.Bytes() //convert to hex bytes
	if len(tmp) == 25 {  //if it is exactly 25 bytes, return
		return tmp
	} else if len(tmp) > 25 { //if it is longer than 25, return nothing
		return nil
	}
	answer := make([]byte, 25)      //make 25 byte container
	for i := 0; i < len(tmp); i++ { //copy converted bytes
		answer[24-i] = tmp[len(tmp)-1-i]
	}
	return answer
}

//encodes big.Int to base58 string
func Big2Base58(val *big.Int) Base58 {
	answer := ""
	valCopy := new(big.Int).Abs(val) //copies big.Int

	if val.Cmp(big.NewInt(0)) <= 0 { //if it is less than 0, returns empty string
		return Base58("")
	}

	tmpStr := ""
	tmp := new(big.Int)
	for valCopy.Cmp(big.NewInt(0)) > 0 { //converts the number into base58
		tmp.Mod(valCopy, big.NewInt(58))                //takes modulo 58 value
		valCopy.Div(valCopy, big.NewInt(58))            //divides the rest by 58
		tmpStr += alphabet[tmp.Int64() : tmp.Int64()+1] //encodes
	}
	for i := (len(tmpStr) - 1); i > -1; i-- {
		answer += tmpStr[i : i+1] //reverses the order
	}
	return Base58(answer) //returns
}

//encodes hex bytes into base58
func Hex2Base58(val []byte) Base58 {
	tmp := Big2Base58(Hex2Big(val)) //encoding of the number without zeroes in front

	//looking for zeros at the beggining
	i := 0
	for i = 0; val[i] == 0 && i < len(val); i++ {
	}
	answer := ""
	for j := 0; j < i; j++ { //adds zeroes from the front
		answer += alphabet[0:1]
	}
	answer += string(tmp) //concatenates

	return Base58(answer) //returns
}

func Hex2Big(b []byte) *big.Int{
	answer:=big.NewInt(0)

	for i:=0; i<len(b); i++{
		answer.Lsh(answer, 8)
		answer.Add(answer, big.NewInt(int64(b[i])))
	}

	return answer
}