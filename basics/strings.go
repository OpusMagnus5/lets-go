package basics

import (
	"fmt"
	"unicode/utf8"
)

func TestStrings() {
	fmt.Println("--Strings--------------------------------------------------------------------------------------------")
	/*
	Go używa wartości typu rune do reprezentowania znaków Unicode. Język Go definiuje typ rune jako alias dla typu int32.
	Co więcej, można założyć, że ciągi znaków są nie tylko sekwencjami bajtów, ale także sekwencjami run.

	W zależności od przypadku użycia, ciągi znaków są powszechnie traktowane jako sekwencje bajtów
	podczas przesyłania danych i jako sekwencje run, gdy wymagane jest sprawdzenie każdego pojedynczego znaku ciągu.

	Jeśli jesteś zainteresowany długością łańcucha w znakach, użyj funkcji RuneCountInString z pakietu unicode/utf8
	*/
	// Emoji example 🗿
	emoji := "🙋🌍❗"
	// len(emoji) // 11
	utf8.RuneCountInString(emoji) // 3
}