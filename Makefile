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

mycli:
	@set -a; . ./docker/.env; set +a; cd cli && go run . $(s)

# --- TEST ---

test: init # Démarre MinIO, attend sa disponibilité et lance tous les tests Go
	@./docker/docker.sh up -d minio; \
	echo "Waiting for MinIO on port $${MINIO_PORT:-9000}..."; \
	ready=0; \
	for attempt in $$(seq 1 30); do \
		if curl --fail --silent "http://localhost:$${MINIO_PORT:-9000}/minio/health/ready" >/dev/null; then \
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
	printf '%s\n' '{"default":"minio-test","aliases":{"minio-test":{"url":"http://localhost:'$${MINIO_PORT:-9000}'","access_key":"test-access","secret_key":"test-secret","region":"us-east-1"}}}' > "$$test_config"; \
	cd cli && MYCLI_CONFIG="$$test_config" go test ./...

test-all: test # Alias historique pour lancer la suite complète


# --- Aide ---

help: # Affiche la liste et la description de toutes les commandes disponibles
	@echo "Commandes disponibles :"
	@echo ""
	@grep -E '^[a-zA-Z0-9_-]+:.*?# .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[1;32m%-20s\033[0m %s\n", $$1, $$2}'