package pkg

import "flag"

var (
	Reject       = flag.String("R", "", "Rejeter les fichiers spécifiés")
	BackGround   = flag.Bool("B", false, "Lancer le téléchargement en arrière-plan")
	Exclude      = flag.String("X", "", "Exclure les dossiers spécifiés")
	XLong        = flag.String("exclude", "", "Exclure les dossiers spécifiés version longue")
	RLong        = flag.String("reject", "", "Rejeter les fichiers spécifiés version longue")
	Output       = flag.String("O", "", "Nom du fichier de sortie")
	FileInput    = flag.String("i", "", "Lire les URL à partir d'un fichier text")
	RateLimit    = flag.String("rate-limit", "", "Limiter la vitesse de téléchargement")
	NewPath      = flag.String("P", "./", "Répertoire de destination")
	Mirroring    = flag.Bool("mirror", false, "Activer le mirroir d'un site web")
	ConvertLinks = flag.Bool("convert-links", false, "Convertir les liens pour le mirroir")
)
