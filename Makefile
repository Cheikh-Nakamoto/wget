# Rebuild complet et exécution
run:
	cargo build --release
	mv target/release/wget_rust .

# Indiquer à make que les règles ne sont pas des fichiers
.PHONY: run
