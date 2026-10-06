# plateforme-mycli

## Lancer les tests

Les tests utilisent un vrai serveur MinIO local. Une seule commande démarre
MinIO si nécessaire, attend que son endpoint de santé réponde, puis lance toute
la suite de tests Go :

```bash
make test
```

Les identifiants définis dans `docker/.env.example` sont locaux et réservés au
développement : `admin` / `password`. Ils ne doivent pas être remplacés par des
identifiants réels dans le fichier de démarrage.

## Développement

Initialiser l’environnement et démarrer MinIO séparément :

```bash
make init
make launch
```