# Conventions du projet


## Langue

**code** doit etre  en anglais : noms de variables, de fonctions et de fichiers,
messages d'erreur, textes d'aide du CLI, noms de tests.
La **documentation** est en français : ce fichier, le README, les issues.

## Nommage des fichiers

**Un fichier par commande, un fichier par opération.** 

```
cmd/bucket_create.go          la commande « bucket create »
internal/s3/create_bucket.go  l'opération S3 « CreateBucket »
```

Les deux moitiés sont inversées, et c'est voulu : dans `cmd/` le nom suit ce que
tape l'utilisateur (`mycli bucket create`), dans `internal/s3/` il suit le nom
de l'opération S3 (`CreateBucket`).

Minuscules et tirets bas, jamais de tiret simple ni de majuscule.


## Structure d'un fichier de commande

Toujours les mêmes cinq étapes, dans cet ordre :

1. la déclaration cobra — `Use`, `Short`, et `RunE`
2. la déclaration des options (`--bucket`, `--file`)
3. la vérification des options obligatoires
4. l'appel d'une fonction de `internal/s3/`
5. l'affichage du résultat, ou le renvoi de l'erreur

Puis une fonction `init()` qui enregistre la commande auprès de son groupe.

## `RunE`, jamais `Run`

Cobra propose deux champs pour l'action d'une commande :

- `Run` ne renvoie rien — l'erreur devrait être traitée sur place
- `RunE` renvoie une `error` — elle remonte jusqu'à `root.go`

## Les erreurs

- Le **résultat** va sur la sortie standard et les **erreurs** vont sur la sortie
  d'erreur, avec un code de sortie non nul.
- `os.Exit` apparaît juste une fois dans tout le projet, dans `cmd/root.go`.
  Une commande ne termine jamais le programme elle-même : elle renvoie son
  erreur.
- Les sept opérations renvoient le type d'erreur commun de
  `internal/s3/errors.go`, jamais une erreur brute.
- Chaque message dit **quoi** a échoué et **quoi faire**.

## Les commandes

- Chaque commande a un `Short` d'une ligne, commençant par un verbe, sans
  article ni point final. Il s'affiche dans la liste du `--help` parent.

## Les tests

Ils vivent à côté du code testé, suffixés `_test.go` — chacun teste ce qu'il
écrit. Ils tournent contre **MinIO réel**, pas contre un serveur simulé.

## Taille des fichiers

400 lignes maximum. Avec un fichier par commande et par opération, aucun ne
devrait dépasser 60.
