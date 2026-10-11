# mycli

Client en ligne de commande pour un serveur de stockage compatible S3, comme
MinIO ou Amazon S3.

## Installation

```bash
npm install -g @derrocethebigorg/mycli
mycli --version
```

À l'installation, le paquet télécharge le binaire de la
[release GitHub](https://github.com/ArthurDescourvieres/plateforme-mycli/releases)
qui correspond à sa version et à votre système (Linux, macOS ou Windows, x64 ou
arm64), puis vérifie sa somme de contrôle SHA-256 avec `checksums.txt`. Aucune
installation de Go n'est nécessaire.

## Utilisation

```bash
export MYCLI_URL=http://localhost:9000
export MYCLI_ACCESS_KEY=<votre-clé-d-accès>
export MYCLI_SECRET_KEY=<votre-clé-secrète>

mycli bucket create --bucket demo
mycli object upload --bucket demo --file rapport.txt
mycli object list --bucket demo
```

La documentation complète (configuration, commandes, correspondance avec
l'API S3) est dans le
[README du projet](https://github.com/ArthurDescourvieres/plateforme-mycli#readme).
