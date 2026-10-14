# TextForge CLI

TextForge est un outil en ligne de commande (CLI) développé en Go. Il permet de lire, analyser et transformer facilement des fichiers texte (.txt) grâce à différentes options de traitement (nettoyage, conversion de casse, remplacement, numérotation, etc...).

---

## Présentation

TextForge prend un fichier texte en entrée (input.txt), applique les transformations demandées par l'utilisateur(grace aux commandes ci dessous), puis sauvegarde le résultat dans un fichier de sortie (output.txt).

Le projet a été conçu pour être simple à utiliser, et les commandes sont afficher en clair dans le README et le main(en commentaire)

---

## Arbo

textforge\ 
|
|go.mod
|
|input.txt
|
|output.txt
|
|main.go
|
|transform.go
|
|README.md

---

## liste des commandes 

    Option	                                         Description

--trim	                            Supprime les espaces inutiles (début/fin, espaces multiples) et ajuste la ponctuation.
--lowercase	                       Convertit l'ensemble du texte en minuscules.
--uppercase	                        Convertit l'ensemble du texte en majuscules.
--capitalize	                   Met en majuscule la première lettre de chaque phrase (après ., !, ?).
--replace "ancien=nouveau"	        Remplace toutes les occurrences d'un texte par un autre.
--number-lines	                   Ajoute le numéro de ligne au début de chaque ligne (ex: 1 | ...).
--stats	                            Affiche le nombre de lignes, de mots et de caractères du fichier d'entrée.
--preview	                       Affiche le résultat dans le terminal sans créer ni modifier le fichier de sortie.
--help	                            Affiche la liste des options et l'aide.