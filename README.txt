Ligne de commande
-----------------
Échec en débogue
go run -tags debug,Chess . -unittest

Échec
go run -tags Chess .

Texte en débogue
go run -tags debug,Text . -unittest

Texte
go run -tags Text .

Échec et texte en débogue
go run -tags debug,Chess,Text . -unittest


Échec et texte
go run -tags Chess,Text .

