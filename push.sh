#!/bin/bash

# Demande à l'utilisateur d'entrer un message de commit
read -p "Entrez le message du commit : " commit_message

# Vérifie si un message a été entré, sinon utilise un message par défaut
if [ -z "$commit_message" ]; then
    commit_message="Mise à jour rapide"
fi

# Ajout des modifications
echo "Ajout des modifications au staging area..."
git add .

# Création du commit
echo "Création du commit avec le message : '$commit_message'..."
git commit -m "$commit_message"

# Poussée vers le dépôt distant
echo "Envoi des modifications au dépôt distant..."
git push

echo "Opérations terminées avec succès !"
