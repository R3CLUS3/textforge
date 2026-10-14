package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// CleanText nettoie les espaces et corrige la ponctuation
func CleanText(content string) string {
	lines := strings.Split(content, "\n")
	var result []string

	for _, line := range lines {
		// strings.Fields découpe par tous les espaces/tabulations et supprime les espaces aux extrémités
		words := strings.Fields(line)
		if len(words) == 0 {
			result = append(result, "")
			continue
		}

		cleanedLine := strings.Join(words, " ")

		// Supprime l'espace inutile avant la ponctuation (. , ! ?)
		punctuations := []string{" .", " ,", " !", " ?"}
		for _, p := range punctuations {
			target := string(p[1]) // la ponctuation sans l'espace
			cleanedLine = strings.ReplaceAll(cleanedLine, p, target)
		}

		result = append(result, cleanedLine)
	}

	return strings.Join(result, "\n")
}

// Capitalize met en majuscule le début des phrases
func Capitalize(content string) string {
	if content == "" {
		return ""
	}

	runes := []rune(content)
	capitalizeNext := true

	for i, r := range runes {
		if capitalizeNext && r >= 'a' && r <= 'z' {
			runes[i] = r - ('a' - 'A')
			capitalizeNext = false
		} else if r == '.' || r == '!' || r == '?' {
			capitalizeNext = true
		}
	}

	return string(runes)
}

// Replace text au format "ancien=nouveau"
func ReplaceText(content, replaceFlag string) (string, error) {
	parts := strings.Split(replaceFlag, "=")
	if len(parts) != 2 || parts[0] == "" {
		return "", fmt.Errorf("erreur: format invalide pour --replace. Utilisation: --replace \"ancien=nouveau\"")
	}

	return strings.ReplaceAll(content, parts[0], parts[1]), nil
}

// NumberLines ajoute "1 | ", "2 | ", etc.
func NumberLines(content string) string {
	lines := strings.Split(content, "\n")
	var result []string

	for i, line := range lines {
		result = append(result, fmt.Sprintf("%d | %s", i+1, line))
	}

	return strings.Join(result, "\n")
}

// PrintStats affiche les statistiques simples
func PrintStats(content string) {
	lines := strings.Split(content, "\n")
	words := strings.Fields(content)
	chars := utf8.RuneCountInString(content)

	fmt.Println("STATISTIQUES")
	fmt.Printf("Lignes : %d\n", len(lines))
	fmt.Printf("Mots : %d\n", len(words))
	fmt.Printf("Caractères : %d\n", chars)
}
