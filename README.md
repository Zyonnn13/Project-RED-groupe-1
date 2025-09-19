## Cyberpunk 2077 :

Un jeu vidéo narratif a choix mutiples dans un univers cybernétique et futuriste développé en language Go, inspiré de l’univers du jeux vidéo Cyberpunk 2077.

## Fonctionnalités :

- Système de combat modulaire (piratage, surchauffe, armes)
- Scénario à embranchements (Nomade, Corpo, Gosse)
- Inventaire et boutique interactifs
- Gestion de la monnaie (Eddies)
- Bande-son immersive
- Nombreuses variantes d'armes
- Système de craft 
- Système de vie 

# 📁 Structure du projet

Voici l’organisation des fichiers et dossiers de votre projet de jeu :

armes/ # Gestion des armes du jeu
     -armes.go
 combat/ # Mécaniques de combat et ennemis
    - hack.go # Piratage en combat
    - surchauffe.go # Effets de surchauffe
    - combat.go # Logique générale du combat
    - Ennemis.go # Définition des ennemis
│
├── histoire/ # Narration et scénarios
│ ├── Acte1.go # Premier acte de l’histoire
│ ├── Acte2.go # Deuxième acte
│ ├── corpo.go # Scénario lié aux corpos
│ ├── Gosse.go # Scénario lié aux gosses des rues
│ ├── Nomade.go # Scénario lié aux nomades
│ └── scenario.go # Structure narrative globale
│
├── Inventaire/ # Gestion de l’inventaire du joueur
│ └── Inventaire.go
│
├── monnaie/ # Système monétaire (Eddies)
│ └── eddies.go
│
├── player/ # Définition et gestion du joueur
│ └── player.go
│
├── PvO/ # Interactions pour joueur
│ └── Shop/ # Système de boutique
│ ├── craft.go # Fabrication d’objets
│ ├── sell.go # Vente d’objets
│ └── shop.go # Interface de la boutique
│
├── sound/ # Ressources sonores
│ └── musique.mp3
│
├── go.mod # Fichier de configuration des dépendances Go
├── go.sum # Sommes de contrôle des modules
├── main.go # Point d’entrée principal du jeu
└── README.md # Documentation du projet


## Lien vers Github : 
https://github.com/Zyonnn13/Project-RED-groupe-1

## Language de Programation de notre projet :
- Golang 
