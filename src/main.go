package main

import (
	"fmt"
	"sort"
	"strings"
)

var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVXWYZ"

func main() {
	key := 5
	original := "A ligeira raposa marrom saltou sobre o cachorro cansado"

	fmt.Printf("Original: %s\n", original)

	ciphered := caesarEncrypt(original, key)
	fmt.Printf("Encriptada: %s\n", ciphered)

	plain := caesarDecrypt(ciphered, key)
	fmt.Printf("Decriptada: %s\n", plain)

	mapaPortugues := sortKeys(estatisticaPortugues())
	keys := make([]string, 0, len(mapaPortugues))
	for _, k := range keys {
		fmt.Printf(k, mapaPortugues[k], " ")
	}

	mapaCifra := sortKeys(freq(ciphered))
	keys = make([]string, 0, len(mapaCifra))
	for _, k := range keys {
		fmt.Printf(k, mapaCifra[k], " ")
	}

}

func sortKeys(mapa map[string]float32) map[string]float32 {
	keys := make([]string, 0, len(mapa))
	for k := range mapa {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return mapa[keys[i]] > mapa[keys[j]]
	})
	for _, k := range keys {
		fmt.Println(k, mapa[k])
	}
	return mapa
}

func freq(text string) map[string]float32 {
	freq := make(map[string]float32)
	for _, num := range text {
		num := strings.ToLower(string(num))
		if num != " " {
			freq[string(num)] = freq[string(num)] + 1.00
		}
	}
	return freq
}

func estatisticaPortugues() map[string]float32 {
	m := map[string]float32{
		"a": 14.63,
		"b": 1.04,
		"c": 3.88,
		"d": 4.99,
		"e": 12.5,
		"f": 1.02,
		"g": 1.30,
		"h": 1.28,
		"i": 6.18,
		"j": 0.40,
		"k": 0.02,
		"l": 2.78,
		"m": 4.74,
		"n": 5.05,
		"o": 10.73,
		"p": 2.52,
		"q": 1.20,
		"r": 6.53,
		"s": 7.81,
		"t": 4.34,
		"u": 4.63,
		"v": 1.67,
		"w": 0.01,
		"x": 0.21,
		"y": 0.01,
		"z": 0.47}

	return m
}

func caesarEncrypt(data string, key int) string {
	newData := ""
	for _, p := range data {
		index := strings.IndexRune(alphabet, p)
		if index == -1 {
			newData = newData + string(p)
		} else {
			newIndex := index + key
			newIndex = modulus(newIndex, len(alphabet))
			newData += alphabet[newIndex : newIndex+1]
		}
	}
	return newData
}

func caesarDecrypt(data string, key int) string {
	newData := ""
	for _, c := range data {
		index := strings.IndexRune(alphabet, c)
		if index == -1 {
			newData = newData + string(c)
		} else {
			newIndex := index - key
			newIndex = modulus(newIndex, len(alphabet))
			newData += alphabet[newIndex : newIndex+1]
		}
	}
	return newData
}

func modulus(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
