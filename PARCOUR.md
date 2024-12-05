### Plan de parcours pour réaliser ce projet "Clone de Wget"

---

### Étape 1 : **Comprendre les besoins**
1. Lisez attentivement la documentation du projet (les fonctionnalités à implémenter, les flags, les exemples).
2. Familiarisez-vous avec l’outil `wget` en testant ses fonctionnalités principales :
   - `$ wget https://example.com/file.zip`
   - `$ wget --mirror https://example.com`

---

### Étape 2 : **Préparer l’environnement**
1. **Choisir le langage :**
   - Optez pour un langage compilé (Go, Rust, C, etc.).
   - Installez les outils nécessaires (compilateur, éditeur, gestionnaire de dépendances).

2. **Installer les bibliothèques nécessaires :**
   - Pour Go : `net/http`, `os`, `io`.
   - Pour Rust : `reqwest`, `tokio` (pour l’asynchrone).
   - Pour C : `libcurl` ou une alternative pour gérer HTTP.

---

### Étape 3 : **Planification des modules**
Divisez le projet en sous-modules pour mieux structurer votre code :
1. **Téléchargement basique d’un fichier.**
2. **Gestion des flags pour personnalisation :**
   - Nom personnalisé (`-O`).
   - Répertoire cible (`-P`).
3. **Limitation de vitesse de téléchargement.**
4. **Téléchargement multiple.**
5. **Mode arrière-plan (`-B`).**
6. **Mirror d’un site web avec options :**
   - Rejet de fichiers (`-R`).
   - Exclusion de répertoires (`-X`).
   - Conversion de liens (`--convert-links`).

---

### Étape 4 : **Développement progressif**
#### 4.1. Téléchargement d’un fichier simple
- **Objectif :** Implémentez une fonction pour télécharger un fichier donné une URL.
- **Tests :** Vérifiez que le fichier est téléchargé correctement avec la taille et le nom exact.

#### 4.2. Gestion des options basiques
- Implémentez les flags :
  - `-O` : Enregistrement avec un nom personnalisé.
  - `-P` : Enregistrement dans un répertoire spécifique.
- **Tests :** Vérifiez que les fichiers sont sauvegardés dans les bons chemins avec les bons noms.

#### 4.3. Limitation de la vitesse de téléchargement
- Ajoutez une fonctionnalité pour limiter la vitesse (`--rate-limit`).
- **Tests :** Simulez un téléchargement lent pour valider la limitation.

#### 4.4. Téléchargement multiple
- Implémentez le flag `-i` pour lire une liste de liens depuis un fichier.
- **Tests :** Vérifiez que tous les fichiers de la liste sont téléchargés.

#### 4.5. Téléchargement en arrière-plan
- Utilisez des threads ou des processus pour exécuter en arrière-plan (`-B`).
- Redirigez la sortie vers un fichier log (`wget-log`).
- **Tests :** Validez la bonne exécution via les logs.

#### 4.6. Mirroring d’un site web
- Implémentez le téléchargement d’un site complet (`--mirror`).
- Ajoutez les options avancées :
  - Rejet de fichiers par type (`-R`).
  - Exclusion de répertoires (`-X`).
  - Conversion de liens (`--convert-links`).
- **Tests :** Validez en téléchargeant un site simple.

---

### Étape 5 : **Optimisation et gestion des erreurs**
1. **Gestion des erreurs HTTP :** Implémentez une gestion des codes HTTP (404, 403…).
2. **Progression visuelle :** Ajoutez une barre de progression et affichez la vitesse de téléchargement.
3. **Asynchrone :** Implémentez des téléchargements simultanés pour accélérer les tâches multiples.

---

### Étape 6 : **Test complet**
1. Créez différents scénarios pour tester toutes les fonctionnalités.
2. Vérifiez les comportements inattendus (fichiers inexistants, URL incorrectes, connexion lente).
3. Validez avec des sites de test ou des fichiers volumineux.

---

### Étape 7 : **Documentation et finalisation**
1. Rédigez un guide utilisateur clair dans le `README`.
2. Nettoyez le code et ajoutez des commentaires si nécessaire.
3. Créez une démonstration (GIF ou vidéo) pour illustrer les fonctionnalités.

---

### Étape 8 : **Améliorations possibles**
1. Implémentez une interface graphique pour simplifier l’utilisation.
2. Ajoutez des statistiques avancées (temps total, vitesse moyenne).
3. Supportez plus de protocoles (FTP, SFTP…).

---

### Résultat final attendu
Un programme fonctionnel capable de reproduire une large gamme des fonctionnalités de `wget`, accompagné d’une documentation claire, de tests complets et d’une bonne gestion des erreurs.