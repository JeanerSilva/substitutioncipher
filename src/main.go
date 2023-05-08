package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func main() {
	key := 5
	original := "A ligeira raposa marrom saltou sobre o cachorro cansado"
	//original := "A compreensão e interpretação de texto são duas ações que estão relacionadas, uma vez que quando se compreende corretamente um texto e seu propósito comunicativo chegamos a determinadas conclusões (interpretação)."

	//original := "Quem se interessa por aprender a falar Português já pode contar com um ensino eficiente. Com os nossos métodos conseguimos ensinar, sobretudo alunos iniciantes, por meio de textos práticos, que favorecem a boa leitura e consequente compreensão do que é ensinado."
	//original := "Qual é a velocidade dos seus downloads? Em poucos segundos, o teste do FAST.com faz uma estimativa da velocidade do seu provedor."

	fmt.Printf("Original: %s\n\n", original)

	originalAdaptada := replaceAscii(strings.ToLower(original))

	fmt.Printf("Original adaptada: %s\n\n", originalAdaptada)

	ciphered := caesarEncrypt(originalAdaptada, key)
	fmt.Printf("Encriptada: %s\n\n", ciphered)

	plain := caesarDecrypt(ciphered, key)
	fmt.Printf("Decriptada: %s\n\n", plain)

	mapaPortugues := sortKeys(estatisticaPortugues())
	fmt.Println("Mapa portugues", mapaPortugues)

	mapaOriginal := sortKeys(freq(originalAdaptada))
	fmt.Println("Mapa original ", mapaOriginal)

	mapaCifra := sortKeys(freq(ciphered))
	fmt.Println("Mapa cifrado  ", mapaCifra)

	chave := sortKeys(freqInt(calculaChave(mapaPortugues, mapaCifra)))
	fmt.Println("\nChaves prováveis em ordem descrescente:", chave)

	keyGuessed := chave[0]
	keyGuessedInt, _ := strconv.Atoi(keyGuessed)
	fmt.Println("\nMensagem decifrada com a chave mais provável:\n", caesarDecrypt(ciphered, keyGuessedInt))

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

func calculaChave(keysPortugues, keysCifra []string) [][]int {

	var res [][]int
	var reslinha []int
	fmt.Println("\nMatriz de probabilidades:")
	for x := 0; x < 5; x++ {
		for i := 0; i < 5; i++ {
			c := int([]byte(keysCifra[x])[0])
			p := int([]byte(keysPortugues[i])[0])
			v := modulus(c-p, len(alphabet))
			if v < 0 {
				reslinha = append(reslinha, v*-1)
			} else {
				reslinha = append(reslinha, v)
			}
		}
		fmt.Println(reslinha)
		res = append(res, reslinha)
		reslinha = nil
	}

	return res

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

func freqInt(arr [][]int) map[string]float32 {
	freq := make(map[string]float32)
	for x := 0; x < 5; x++ {
		for i := 0; i < 5; i++ {
			freq[fmt.Sprint(arr[x][i])] = freq[fmt.Sprint(arr[x][i])] + 1
		}
	}
	fmt.Println("\nMapa de frequência:", freq)
	return freq
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
