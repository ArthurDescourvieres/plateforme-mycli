# plateforme-mycli

## Prérequis

- Go 1.26.6 ou supérieur
- Docker avec Docker Compose
- Make

Toutes les commandes `make` doivent être exécutées à la racine du projet.

## Démarrer le projet

Pour voir toutes les commandes disponibles, exécuter :

```bash
make help
```

Initialiser l'environnement. Cette commande crée `docker/.env` à partir de
`docker/.env.example` et prépare le script Docker :



```bash
make init
```

Lancer MinIO :

```bash
make launch
```

MinIO est ensuite disponible sur :

- API S3 : http://localhost:9000
- Console web : http://localhost:9001

Les identifiants locaux sont `admin` / `password`.


## Tests

Les tests utilisent un vrai serveur MinIO local. Une seule commande démarre
MinIO si nécessaire, attend que son endpoint de santé réponde, puis lance toute
la suite de tests Go :

```bash
make test
```

Les identifiants définis dans `docker/.env.example` sont locaux et réservés au
développement : `admin` / `password`. Ils ne doivent pas être remplacés par des
identifiants réels dans le fichier de démarrage.

## Qualité du code

Formater le code et appliquer les corrections automatiques disponibles :

```bash
make go-lint
```

Vérifier le formatage et le lint sans modifier les fichiers :

```bash
make go-lint-check
```

Vérifier les dépendances :

```bash
make go-deps-check
```