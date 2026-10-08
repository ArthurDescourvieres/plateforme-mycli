# Variables par défaut
s ?= 

.PHONY: env perms init launch remove start stop restart mycli test test-all help


# --- INSTALL ---

env: # Crée le fichier .env à partir de l'exemple s'il n'existe pas déjà
	@if [ ! -f "./docker/.env" ]; then \
		cp ./docker/.env.example ./docker/.env; \
		echo "Created .env from example"; \
	else \
		echo ".env already exists"; \
	fi

perms: # Donne les droits d'exécution au script docker helper
	@chmod +x ./docker/docker.sh

init: env perms # Exécute toutes les étapes d'initialisation pour préparer l'environnement de développement


# --- DOCKER ---

launch: # Lance les services Docker en mode développement avec docker-compose
	./docker/docker.sh up -d --build

remove: # Arrête les conteneurs et supprime les volumes associés (remise à zéro)
	./docker/docker.sh down -v

start: # Démarre les conteneurs existants (ex: make start ou make start s=api)
	./docker/docker.sh start $(s)

stop: # Arrête temporairement les conteneurs (ex: make stop ou make stop s=api)
	./docker/docker.sh stop $(s)

restart: # Redémarre un ou tous les conteneurs proprement (ex: make restart s=api)
	./docker/docker.sh restart $(s)

# --- DEV ---

mycli: # Lance l'interface en ligne de commande (CLI) pour interagir avec les services (ex: make mycli ou make mycli s="list buckets")
	@set -a; . ./docker/.env; set +a; cd cli && go run . $(s)

go-lint: # Lance le linter Go pour vérifier le code source et corriger automatiquement les problèmes de formatage
	@cd cli && golangci-lint fmt && golangci-lint run --fix

go-lint-check: # Lance le linter Go pour vérifier le code source sans corriger automatiquement les problèmes de formatage
	@test -z "$$(cd cli && gofmt -l .)" && cd cli && golangci-lint run

go-deps-check: # Vérifie que toutes les dépendances Go sont correctement installées et à jour
	@cd cli && go mod verify && go mod tidy -diff && govulncheck ./...
# --- TEST ---

test: init # Démarre MinIO, attend sa disponibilité et lance tous les tests Go
	@set -a; . ./docker/.env; set +a; \
	health_url="http://localhost:$${MINIO_PORT:-9000}/minio/health/ready"; \
	if ! curl --fail --silent "$$health_url" >/dev/null; then \
		./docker/docker.sh up -d minio; \
	fi; \
	echo "Waiting for MinIO on port $${MINIO_PORT:-9000}..."; \
	ready=0; \
	for attempt in $$(seq 1 30); do \
		if curl --fail --silent "$$health_url" >/dev/null; then \
			ready=1; \
			break; \
		fi; \
		sleep 1; \
	done; \
	if [ "$$ready" -ne 1 ]; then \
		echo "MinIO did not become ready after 30 seconds" >&2; \
		exit 1; \
	fi; \
	test_config=$$(mktemp); \
	trap 'rm -f "$$test_config"' EXIT; \
	printf '%s\n' '{"default":"minio-test","aliases":{"minio-test":{"url":"http://localhost:'$${MINIO_PORT:-9000}'","access_key":"'$${MINIO_USER:-admin}'","secret_key":"'$${MINIO_PASS:-password}'","region":"'$${MINIO_REGION:-us-east-1}'"}}}' > "$$test_config"; \
	cd cli && env -u MINIO_PORT -u MINIO_CONSOLE_PORT -u MINIO_USER -u MINIO_PASS -u MINIO_ENDPOINT -u MINIO_REGION MYCLI_CONFIG="$$test_config" go test ./...

test-all: test # Alias historique pour lancer la suite complète


# --- Aide ---

help: # Affiche la liste et la description de toutes les commandes disponibles
	@echo "Commandes disponibles :"
	@echo ""
	@grep -E '^[a-zA-Z0-9_-]+:.*?# .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[1;32m%-20s\033[0m %s\n", $$1, $$2}'