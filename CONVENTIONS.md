# Conventions du projet

À lire avant d'écrire la première ligne. L'organisation des fichiers et le
contrat du client S3 sont dans l'issue épinglée du dépôt.

## Langue

Le **code** est en anglais : noms de variables, de fonctions et de fichiers,
messages d'erreur, textes d'aide du CLI, noms de tests.
La **documentation** est en français : ce fichier, le README, les issues.

## Nommage des fichiers

**Un fichier par commande, un fichier par opération.** Personne ne partage un
fichier avec personne : c'est ce qui permet de travailler à quatre sans conflit.

```
cmd/bucket_create.go          la commande « bucket create »
internal/s3/create_bucket.go  l'opération S3 « CreateBucket »
```

Les deux moitiés sont inversées, et c'est voulu : dans `cmd/` le nom suit ce que
tape l'utilisateur (`mycli bucket create`), dans `internal/s3/` il suit le nom
de l'opération S3 (`CreateBucket`). Chaque fichier porte le nom de ce qu'il est
dans son monde.

Minuscules et tirets bas, jamais de tiret simple ni de majuscule.

## Nommage en Go

| Quoi | Forme | Exemple |
|---|---|---|
| Commande cobra | `<nom>Cmd` | `bucketCmd`, `bucketCreateCmd` |
| Fonction exportée | Majuscule | `Load`, `ListBuckets` |
| Variable interne au paquet | minuscule | `rootCmd`, `bucketFlag` |

Une **majuscule** rend le nom visible depuis les autres paquets, une
**minuscule** l'enferme dans son dossier. Ce n'est pas un choix de style : c'est
le mécanisme de visibilité de Go.

Pas d'abréviations inventées : `bucket`, pas `bkt`.

## Structure d'un fichier de commande

Toujours les mêmes cinq étapes, dans cet ordre :

1. la déclaration cobra — `Use`, `Short`, et `RunE`
2. la déclaration des options (`--bucket`, `--file`)
3. la vérification des options obligatoires
4. l'appel d'une fonction de `internal/s3/`
5. l'affichage du résultat, ou le renvoi de l'erreur

Puis une fonction `init()` qui enregistre la commande auprès de son groupe.
Chaque fichier s'enregistre seul : ajouter une commande, c'est ajouter un
fichier, sans toucher à aucune liste centrale.

## `RunE`, jamais `Run`

Cobra propose deux champs pour l'action d'une commande :

- `Run` ne renvoie rien — l'erreur devrait être traitée sur place
- `RunE` renvoie une `error` — elle remonte jusqu'à `root.go`

**On utilise `RunE` partout.** C'est ce qui permet d'avoir un seul endroit dans
tout le projet qui affiche l'erreur et termine le programme.

## Les erreurs

- Le **résultat** va sur la sortie standard, les **erreurs** sur la sortie
  d'erreur, avec un code de sortie non nul.
- `os.Exit` n'apparaît **qu'une fois** dans tout le projet, dans `cmd/root.go`.
  Une commande ne termine jamais le programme elle-même : elle renvoie son
  erreur.
- Les sept opérations renvoient le type d'erreur commun de
  `internal/s3/errors.go`, jamais une erreur brute.
- Chaque message dit **quoi** a échoué et **quoi faire**. « bucket "photos"
  introuvable » vaut mieux qu'une trace technique.

## Les commandes

- **Un tiret = une option, un mot = une action.** `mycli object upload`, jamais
  `mycli object -u`.
- Pas de raccourci à une lettre pour une sous-commande. Seules exceptions,
  fournies par cobra : `-h/--help` et `-v/--version`.
- Les paramètres passent par des **options nommées** (`--bucket`, `--file`),
  pas par des arguments devinés à leur position.
- Chaque commande a un `Short` d'une ligne, commençant par un verbe, sans
  article ni point final. Il s'affiche dans la liste du `--help` parent.

## Les frontières entre paquets

- Un fichier de `cmd/` ne fait **jamais** de HTTP. Il ne connaît ni URL, ni
  en-tête, ni signature.
- `internal/s3/` ne connaît **pas** cobra. Aucun import de cobra, aucun
  affichage à l'écran : il renvoie des valeurs et des erreurs.
- Les sept opérations **ne signent pas**. Elles construisent leur requête et la
  passent à la fonction d'envoi du client, qui signe.

## Les commentaires

Un commentaire explique un **pourquoi** qu'on ne peut pas deviner en lisant le
code : un contournement, une contrainte externe, un choix surprenant. Une ligne
suffit.

Pas de commentaire qui répète ce que le code dit déjà. Si un bloc a besoin de
cinq lignes d'explication pour être compris, c'est le bloc qu'il faut
simplifier.

## Les tests

Ils vivent à côté du code testé, suffixés `_test.go` — chacun teste ce qu'il
écrit. Ils tournent contre **MinIO réel**, pas contre un serveur simulé.

Deux règles qui évitent des heures perdues :

- **nom de bucket unique** à chaque exécution : Go lance les tests de paquets
  différents en parallèle, deux tests qui partagent un nom se marchent dessus ;
- **nettoyage garanti même en cas d'échec**, sinon l'exécution suivante échoue
  sur « ce bucket existe déjà ».

## Git

- Branche par issue : `feat/14-commandes-bucket`, `docs/conventions`.
- Messages de commit préfixés : `feat:`, `fix:`, `test:`, `docs:`, `chore:`.
- **Pull request relue par quelqu'un d'autre** avant fusion.
- Pour fermer une issue automatiquement, le mot-clé doit être **en anglais** :
  `Closes #14`. En français, ça ne crée qu'un lien.
- **Pas de pull requests empilées.** Chacun part de `main` et y revient. Comme
  personne ne partage de fichier, il n'y a aucune raison d'empiler.

## Sécurité

Le dépôt est **public**.

- Aucune clé d'accès, aucun secret dans le dépôt — même temporairement, même
  dans un exemple. Un secret poussé une fois reste dans l'historique.
- Le fichier de configuration local est ignoré par Git.
- Ne jamais afficher ni journaliser une clé secrète ou un en-tête
  `Authorization`, y compris dans un message d'erreur.
- Les jetons de publication vivent dans les secrets du dépôt GitHub.

## Taille des fichiers

400 lignes maximum. Avec un fichier par commande et par opération, aucun ne
devrait dépasser 60.
