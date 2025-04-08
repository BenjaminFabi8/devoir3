# Devoir3

## Fonctionnalités
- **Environnement basé sur une grille** : Une grille personnalisable où les agents se déplacent et interagissent.
- **Types d’agents** :
  - Agent Aléatoire : Se déplace de manière aléatoire sur la grille.
  - Agent A* : Utilise l’algorithme A* pour naviguer vers les objectifs.
  - Agent A* avec Attente : Similaire à l’agent A*, mais inclut un mécanisme de délai.
- **Système de journalisation** : Enregistre les déplacements et interactions des agents dans un format structuré.
- **Opérations atomiques personnalisées** : Implémente des opérations atomiques pour la manipulation des cellules de la grille. Chaque case dans la grille est en format atomic. Donc, les opérations de changer leurs valeurs (lors du déplacement) se font en un seul coup. Cela évite donc que 2 agents se déplacent au même endroit au même moment. Cela permet aussi de ne faire aucune attente entre chaque agent.

## Comment exécuter
1. Assurez-vous que Go est installé (version 1.23.4 ou supérieure).
2. Déplacez-vous dans le dossier `src` :
   ```sh
   cd src
3. Exécuter la commande:
    ```sh
#filename est le fichier pour la carte contenant le jeux
   go run . <filename>


### Commande possible une fois le programme en cours

start           -       Exécute le programme

log <filename>  -       Génère le fichier de log avec le nom passé en paramètre

help            -       Affiche les commandes possibles
