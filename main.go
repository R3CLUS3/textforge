// on peut utiliser  --number-lines : pour numéroter les lignes,
// --stats : pour afficher les statistiques du texte,
// --preview : pour afficher le résultat dans le terminal
// --replace "ancien=nouveau" : pour remplacer du texte
// --capitalize : pour mettre en majuscule le début des phrases
// --trim : pour nettoyer les espaces et la ponctuation
// --lowercase : pour mettre le texte en minuscules
// --uppercase : pour mettre le texte en majuscules
// --help : pour afficher l'aide
// voici le code:
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Déclaration des options qu'on peut passer en ligne de commande
	trimFlag := flag.Bool("trim", false, "Nettoyer les espaces")
	lowerFlag := flag.Bool("lowercase", false, "Convertir en minuscules")
	upperFlag := flag.Bool("uppercase", false, "Convertir en majuscules")
	capFlag := flag.Bool("capitalize", false, "Majuscule en début de phrase")
	replaceFlag := flag.String("replace", "", "Remplacer du texte (ancien=nouveau)")
	numberFlag := flag.Bool("number-lines", false, "Numéroter les lignes")
	statsFlag := flag.Bool("stats", false, "Afficher les statistiques")
	previewFlag := flag.Bool("preview", false, "Aperçu sans modifier le fichier")

	// Récupération effective des arguments passés par l'utilisateur
	flag.Parse()

	// Incompatibilité : impossible de mettre le texte à la fois tout en majuscules et tout en minuscules
	if *lowerFlag && *upperFlag {
		fmt.Println("Erreur: les options --lowercase et --uppercase ne peuvent pas être utilisées ensemble.")
		os.Exit(1)
	}

	// On vérifie qu'on a bien au moins le fichier d'entrée et le fichier de sortie
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Utilisation: textforge [options] <input.txt> <output.txt>")
		os.Exit(1)
	}

	inputFile := args[0]
	outputFile := args[1]

	// Lecture du fichier source
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Erreur: le fichier '%s' est introuvable ou inaccessible.\n", inputFile)
		os.Exit(1)
	}

	// Conversion des octets dus au fichier en chaîne de caractères
	content := string(data)

	// Affichage des infos sur le texte si l'option est présente
	if *statsFlag {
		PrintStats(content)
	}

	// Application du remplacement de texte
	if *replaceFlag != "" {
		content, err = ReplaceText(content, *replaceFlag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Nettoyage des espaces et de la ponctuation
	if *trimFlag {
		content = CleanText(content)
	}

	// Changement de la casse (minuscules ou majuscules)
	if *lowerFlag {
		content = strings.ToLower(content)
	} else if *upperFlag {
		content = strings.ToUpper(content)
	}

	// Passage des débuts de phrases en majuscule
	if *capFlag {
		content = Capitalize(content)
	}

	// Ajout de la numérotation en début de ligne
	if *numberFlag {
		content = NumberLines(content)
	}

	// Envoi du résultat dans le terminal ou enregistrement sur le disque
	if *previewFlag {
		fmt.Println(content)
	} else {
		err = os.WriteFile(outputFile, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Erreur lors de l'écriture du fichier '%s'\n", outputFile)
			os.Exit(1)
		}
		fmt.Println("Traitement terminé")
	}
}
