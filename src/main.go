package main

import (
	"fmt"
	"sort"
	"strings"
)

var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVXWYZ"

func main() {
	key := 5
	//original := "A ligeira raposa marrom saltou sobre o cachorro cansado"
	//original := "A compreensão e interpretação de texto são duas ações que estão relacionadas, uma vez que quando se compreende corretamente um texto e seu propósito comunicativo chegamos a determinadas conclusões (interpretação)."

	original := "Quem se interessa por aprender a falar Português já pode contar com um ensino eficiente. Com os nossos métodos conseguimos ensinar, sobretudo alunos iniciantes, por meio de textos práticos, que favorecem a boa leitura e consequente compreensão do que é ensinado."
	fmt.Printf("Original: %s\n\n", original)

	original = replaceAscii(original)

	fmt.Printf("Original adaptada: %s\n\n", original)

	ciphered := caesarEncrypt(original, key)
	fmt.Printf("Encriptada: %s\n\n", ciphered)

	plain := caesarDecrypt(ciphered, key)
	fmt.Printf("Decriptada: %s\n\n", plain)

	mapaPortugues := sortKeys(estatisticaPortugues())
	fmt.Println("Mapa portugues", mapaPortugues)

	mapaOriginal := sortKeys(freq(original))
	fmt.Println("Mapa original ", mapaOriginal)

	mapaCifra := sortKeys(freq(ciphered))
	fmt.Println("Mapa cifrado  ", mapaCifra)
	c1, c2, c3 := calculaChave(mapaPortugues, mapaCifra)
	fmt.Printf("\nProváveis chaves: %d, %d, %d.\n", c1, c2, c3)
}

func replaceAscii(result string) string {
	m := map[string]string{
		"ã": "a", "â": "a", "á": "a", "à": "a",
		"ê": "e", "é": "e", "è": "e",
		"í": "i", "ì": "i",
		"õ": "o", "ô": "o", "ó": "o",
		"ú": "u",
		"ç": "c"}

	for v, k := range m {
		result = strings.ReplaceAll(result, v, k)
	}

	return result
}

func calculaChave(keysPortugues, keysCifra []string) (int, int, int) {
	var result0, result1, result2 int

	c0 := int([]byte(keysCifra[0])[0])
	p0 := int([]byte(keysPortugues[0])[0])
	p1 := int([]byte(keysPortugues[1])[0])
	p2 := int([]byte(keysPortugues[2])[0])

	res0 := p0 - c0
	if res0 < 0 {
		result0 = res0 * -1
	}

	res1 := p1 - c0
	if res1 < 0 {
		result1 = res1 * -1
	}

	res2 := p2 - c0
	if res2 < 0 {
		result2 = res2 * -1
	}

	return result0, result1, result2

}

func sortKeys(mapa map[string]float32) []string {
	keys := make([]string, 0, len(mapa))
	var result []string
	for k := range mapa {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return mapa[keys[i]] > mapa[keys[j]]
	})

	for _, k := range keys {
		result = append(result, k)
	}

	return result
}

func freq(text string) map[string]float32 {
	freq := make(map[string]float32)
	for _, num := range text {
		num := strings.ToLower(string(num))
		if num != " " {
			freq[string(num)] = freq[string(num)] + 1.00
		}
	}
	//tem que retornar a frequência e não a quantidade
	return freq
}

func estatisticaPortugues() map[string]float32 {
	m := map[string]float32{
		"a": 14.63, "b": 1.04, "c": 3.88, "d": 4.99, "e": 12.5, "f": 1.02, "g": 1.30, "h": 1.28, "i": 6.18, "j": 0.40, "k": 0.02,
		"l": 2.78, "m": 4.74, "n": 5.05, "o": 10.73, "p": 2.52, "q": 1.20, "r": 6.53, "s": 7.81, "t": 4.34, "u": 4.63, "v": 1.67,
		"w": 0.01, "x": 0.21, "y": 0.01, "z": 0.47}

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

/*
	keysPortugues := make([]string, 0, len(mapaPortugues))
	for v, k := range keysPortugues {
		//	fmt.Printf(k, mapaPortugues[k])
		arrayPort = append(arrayPort, k)
		fmt.Println("array", arrayPort[v])
	}
	fmt.Println(arrayPort)

		mapaCifra := sortKeys(freq(ciphered))
		keysCifra := make([]string, 0, len(mapaCifra))
		for _, k := range keysCifra {
			fmt.Printf(k, mapaCifra[k])
		}

		fmt.Println(keysPortugues[1])

		fmt.Printf("A chave provável é %d", calculaChave(keysPortugues, keysCifra))
*/
