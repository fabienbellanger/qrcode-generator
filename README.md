# qrcode-generator

CLI en Go pour générer des QR Codes au format PNG.

## Installation

### Depuis les sources

```bash
go install .
```

### Via Make

```bash
make build
```

Les binaires sont générés dans le dossier `build/`.

## Utilisation

```bash
qrcode-generator [options] <texte>
```

### Options

| Flag | Description                             | Obligatoire |
| ---- | --------------------------------------- | ----------- |
| `-o` | Chemin du fichier image de sortie (PNG) | Oui         |
| `-s` | Taille de l'image en pixels             | Oui         |
| `-h` | Affiche l'aide                          | Non         |

### Exemples

```bash
# Générer un QR Code de 512x512 px
qrcode-generator -o qrcode.png -s 512 "https://example.com"

# Afficher l'aide
qrcode-generator -h
```

## Cross-compilation

Le Makefile permet de compiler pour plusieurs plateformes :

```bash
# Compiler pour toutes les plateformes (Mac + Linux, 32 et 64 bit)
make all

# Compiler uniquement pour Mac
make darwin

# Compiler uniquement pour Linux
make linux

# Nettoyer les binaires
make clean
```

### Plateformes supportées

| OS    | Architecture | Binaire                               |
| ----- | ------------ | ------------------------------------- |
| macOS | amd64        | `build/qrcode-generator-darwin-amd64` |
| macOS | arm64        | `build/qrcode-generator-darwin-arm64` |
| Linux | 386          | `build/qrcode-generator-linux-386`    |
| Linux | amd64        | `build/qrcode-generator-linux-amd64`  |

## Dépendances

- [go-qrcode](https://github.com/skip2/go-qrcode) - Génération de QR Codes en Go
