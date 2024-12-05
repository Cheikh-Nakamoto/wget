# README - Clone de Wget

## Logo
![Gitea](./logo.png)

---

## Introduction

Ce projet a pour objectif de recréer certaines fonctionnalités de l’outil `wget` à l’aide d’un langage compilé de votre choix (C, Rust, Go, etc.). L’objectif est de maîtriser les concepts liés au téléchargement de fichiers, à la gestion de ressources web, ainsi qu’aux algorithmes de parcours et de récursivité.

---

## Fonctionnalités

Le programme implémentera les fonctionnalités suivantes :

1. **Téléchargement classique d’un fichier :**  
   `$ ./program https://example.com/file.zip`

2. **Téléchargement avec nom personnalisé :**  
   `$ ./program -O=new_name.zip https://example.com/file.zip`

3. **Téléchargement dans un répertoire spécifique :**  
   `$ ./program -P=/chemin/vers/dossier https://example.com/file.zip`

4. **Limitation de la vitesse de téléchargement :**  
   `$ ./program --rate-limit=200k https://example.com/file.zip`

5. **Téléchargement en arrière-plan :**  
   `$ ./program -B https://example.com/file.zip`  
   La sortie sera redirigée vers `wget-log`.

6. **Téléchargement multiple via un fichier :**  
   `$ ./program -i=links.txt`

7. **Mirroring d’un site web :**  
   `$ ./program --mirror https://example.com`  
   Le site sera sauvegardé localement avec sa structure.

---

## Utilisation

Voici les étapes principales :

1. **Lancer un téléchargement simple :**
   ```bash
   $ ./program https://example.com/file.zip
   Début : 2024-12-04 15:00:00
   Réponse : 200 OK
   Taille : 10 MB
   Enregistrement sous : ./file.zip
   Progression : [====================] 100% 1 MB/s
   Fin : 2024-12-04 15:00:10
   ```

2. **Téléchargement avec options :**
   - **Nom personnalisé** : `-O=fichier.zip`
   - **Répertoire spécifique** : `-P=/chemin/vers/dossier`
   - **Limiter la vitesse** : `--rate-limit=200k`
   - **Arrière-plan** : `-B` (sortie dans `wget-log`)

3. **Téléchargement multiple avec `-i=links.txt` :**
   Exemple de contenu de `links.txt` :
   ```
   https://example.com/file1.zip
   https://example.com/file2.zip
   ```
   Commande :  
   `$ ./program -i=links.txt`

4. **Mirror de site web avec `--mirror` :**
   ```bash
   $ ./program --mirror https://example.com
   ```

---

## Flags supplémentaires

### Pour `--mirror`
- **Exclusion de fichiers par type :** `-R=jpg,png`
- **Exclusion de répertoires :** `-X=/assets,/css`
- **Conversion des liens pour usage hors ligne :** `--convert-links`

---

## Pré-requis

- Connaissances sur les protocoles HTTP, FTP et sur la gestion des fichiers.
- Familiarité avec un langage compilé (C, Go, Rust…).

---

## Contribution

Si vous trouvez un problème ou souhaitez proposer une amélioration, n’hésitez pas à soumettre une issue ou une pull request !

### Architecture du projet "Clone de Wget"

---

Une architecture bien structurée est essentielle pour assurer la lisibilité, la maintenabilité et l'évolutivité du projet. Voici une proposition d’architecture modulaire pour ce projet :

---

### 1. **Structure des dossiers**
```
.
├── main.rs                            # Point d'entrée principal
├── cmd/                               # Gestion des commandes et des arguments
│   ├── flags.rs                       # Parsing et gestion des flags
│   ├── commands.rs                    # Logique liée aux différentes commandes
├── internal/                          # Fonctionnalités principales
│   ├── downloader/                    # Module de téléchargement
│   │   ├── download.rs                # Téléchargement de fichiers
│   │   ├── rate_limiter.rs            # Gestion de la vitesse de téléchargement
│   │   ├── progress.rs                # Gestion de la barre de progression
│   └── mirror/                        # Module pour le mirroring
│       ├── mirror.rs                  # Téléchargement d’un site complet
│       ├── filters.rs                 # Gestion des rejets (types, répertoires)
├── utils/                             # Fonctions utilitaires
│   ├── logger.rs                      # Gestion des logs
│   ├── file_utils.rs                  # Manipulation des fichiers et répertoires
│   ├── http_utils.rs                  # Fonctions liées aux requêtes HTTP
└── README.md                          # Documentation
```

---

### 2. **Description des composants**

#### **2.1. Fichier principal : `main.go`**
- **Rôle :**
  - Initialiser le programme.
  - Lire les arguments de la ligne de commande.
  - Appeler les modules appropriés selon les flags et options fournis.
- **Exemple (en Go) :**
  ```go
  package main

  import (
      "os"
      "cmd"
  )

  func main() {
      // Parse des flags et gestion des erreurs
      options := cmd.ParseFlags(os.Args)
      cmd.ExecuteCommand(options)
  }
  ```

---

#### **2.2. Module `cmd` (Commandes et Flags)**

- **Fichier `flags.go` :**
  - Parse les arguments de la ligne de commande.
  - Définit les options valides (`-O`, `-P`, `-B`, `--mirror`, etc.).
  - Valide les combinaisons d’options.

- **Fichier `commands.go` :**
  - Contient la logique pour rediriger les commandes vers les modules correspondants (ex. télécharger un fichier, lancer un mirroring, etc.).

---

#### **2.3. Module `downloader` (Téléchargement)**

- **Fichier `download.go` :**
  - Gère le téléchargement de fichiers à partir d’URL.
  - Vérifie le code de réponse HTTP.
  - Implémente la logique pour gérer les différents chemins (`-P`, `-O`).
  - Exemple (en Go) :
    ```go
    func DownloadFile(url string, outputPath string) error {
        // Requête HTTP et écriture du fichier
    }
    ```

- **Fichier `rate_limiter.go` :**
  - Implémente la limitation de vitesse (`--rate-limit`).
  - Exemple : Implémentez un débit limité en gérant les pauses entre les lectures.

- **Fichier `progress.go` :**
  - Affiche une barre de progression avec le pourcentage, la vitesse et le temps restant.

---

#### **2.4. Module `mirror` (Mirroring)**

- **Fichier `mirror.go` :**
  - Implémente le téléchargement récursif des pages d’un site.
  - Parse les fichiers HTML/CSS pour récupérer les liens à télécharger.
  - Respecte les options comme `-R` et `-X`.

- **Fichier `filters.go` :**
  - Implémente les filtres pour rejeter certains types de fichiers ou exclure des répertoires.

---

#### **2.5. Module `utils` (Utilitaires)**

- **Fichier `logger.go` :**
  - Gère l’écriture des logs dans un fichier (`wget-log`).
  - Exemple (en Go) :
    ```go
    func WriteLog(message string) {
        // Écrit un message dans le fichier log
    }
    ```

- **Fichier `file_utils.go` :**
  - Gère la création de répertoires et la vérification des chemins.

- **Fichier `http_utils.go` :**
  - Contient des fonctions pour effectuer des requêtes HTTP.
  - Gère les redirections ou les erreurs HTTP.

---

### 3. **Flux de travail principal**
1. **Lancement du programme :**
   - Les arguments sont analysés par `cmd.ParseFlags`.
   - Le programme décide de la commande à exécuter (téléchargement, mirroring, etc.).

2. **Téléchargement :**
   - Le module `downloader` est appelé pour télécharger un fichier ou une liste de fichiers.
   - Les options comme `-O`, `-P`, ou `--rate-limit` sont prises en compte.

3. **Mirroring :**
   - Le module `mirror` télécharge récursivement les pages et ressources liées, tout en respectant les filtres (`-R`, `-X`).

4. **Affichage et Logs :**
   - Une barre de progression est affichée.
   - Les événements sont consignés dans le fichier log via `utils.logger`.

---

### 4. **Exemple simplifié de flux d’exécution**

#### Commande :
```bash
./program --mirror -R=jpg,gif https://example.com
```

#### Étapes :
1. **Parsing des flags :**
   - `--mirror` → mode mirroring.
   - `-R=jpg,gif` → filtre pour rejeter les fichiers `.jpg` et `.gif`.

2. **Initialisation du mirroring :**
   - Analyse la page principale.
   - Récupère les fichiers liés via les balises `<a>`, `<img>`, et `<link>`.

3. **Téléchargement récursif :**
   - Appelle `downloader` pour chaque fichier trouvé.
   - Respecte les filtres définis.

4. **Finalisation :**
   - Convertit les liens pour une utilisation hors ligne.
   - Log des actions dans `wget-log`.

---

