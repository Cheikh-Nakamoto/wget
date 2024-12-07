# Rebuild complet et exécution
run:
	cargo build --release
	rm -rf wget_rust
	mv target/release/wget_rust .

# Indiquer à make que les règles ne sont pas des fichiers
.PHONY: run
