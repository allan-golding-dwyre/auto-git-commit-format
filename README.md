# auto-git-commit-format (AGCF)

Permet de formatter automatiquement un commit vers le type voulu. Il peut aussi permettre d'invalider le commit si il y a trop de caractère, ou si le commit n'est pas en anglais.

## Installation:

``` sh
curl -sSL https://raw.githubusercontent.com/allan-golding-dwyre/auto-git-commit-format/main/install.sh | bash
```

## Usage:
  agcf [command]

### Available Commands:
  * `build`       Build / déploiement
  * `deps`        Mise à jour dépendances
  * `docs`        Documentation
  * `feat`        Nouvelle fonctionnalité
  * `fix`         Correction de bug
  * `refactor`    Refactoring
  * `remove`      Suppression de code
  ---
  * `help`        Help about any command
  * `completion`  Generate the autocompletion script for the specified shell


### Example:

```bash

agcf feat "Adding a feature with a message"

# Also works:
agcf feat Adding a feature with a message
```


