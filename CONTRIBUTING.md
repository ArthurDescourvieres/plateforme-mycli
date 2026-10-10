# Contribuer

Merci de contribuer à `plateforme-mycli`.

## Prérequis

- Go 1.26.6 ou supérieur
- Docker avec Docker Compose
- Make

Toutes les commandes `make` doivent être exécutées depuis la racine du projet.

## Préparer l'environnement

Initialiser l'environnement local :

```bash
make init
```

Cette commande crée `docker/.env` à partir de
`docker/.env.example` s'il n'existe pas déjà.

Lancer les services Docker :

```bash
make launch
```

Le projet Compose utilise le nom défini par `PROJECT` dans `docker/.env`.
Ce nom peut être remplacé ponctuellement :

```bash
make PROJECT=mon-projet launch
```

Les identifiants `admin` / `password` sont réservés au développement local.
N'utilisez jamais de véritables identifiants dans `docker/.env.example`.

## Vérifications locales

Avant de créer une Pull Request, exécuter les vérifications suivantes.

### Formatage et lint

Appliquer le formatage et les corrections automatiques disponibles :

```bash
make go-lint
```

Vérifier le formatage et le lint sans modifier les fichiers :

```bash
make go-lint-check
```

### Dépendances

Vérifier les modules Go et rechercher les vulnérabilités connues :

```bash
make go-deps-check
```

### Tests

Les tests utilisent un serveur MinIO local. La commande démarre MinIO si
nécessaire, attend que son endpoint de santé soit disponible, puis exécute les
tests Go :

```bash
make test
```

## CI

Le workflow GitHub Actions est défini dans
[`.github/workflows/ci.yml`](.github/workflows/ci.yml). Il exécute quatre jobs :

- `format-lint` vérifie `gofmt` et lance `golangci-lint` ;
- `dependencies` vérifie les modules Go et lance `govulncheck` ;
- `leaks` recherche les secrets exposés avec Gitleaks ;
- `test` démarre MinIO et exécute la suite de tests.

Une Pull Request doit passer l'ensemble de ces vérifications.

## Workflow de contribution

1. Créer une branche dédiée à partir de `main`.
2. Implémenter les changements en limitant le périmètre de la Pull Request.
3. Ajouter ou mettre à jour les tests concernés.
4. Exécuter les vérifications locales.
5. Vérifier qu'aucun secret, identifiant réel ou fichier local sensible n'est ajouté.
6. Créer une Pull Request vers `main` en décrivant les changements et les tests effectués.

## Pull Requests

Une Pull Request doit :

- expliquer clairement le problème corrigé ou la fonctionnalité ajoutée ;
- décrire les changements importants ;
- indiquer les commandes de test exécutées et leur résultat ;
- signaler les changements de configuration ou de comportement ;
- rester ciblée et ne pas mélanger de changements sans rapport.

## Secrets et fichiers locaux

Ne committez pas :

- `docker/.env` ;
- des clés d'accès, mots de passe ou tokens réels ;
- des fichiers de configuration propres à votre machine ;
- des artefacts de build ou des fichiers temporaires.

Utilisez `docker/.env.example` pour documenter uniquement les valeurs de
développement sans caractère sensible.
