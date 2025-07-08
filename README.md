# term

**term** is a lightweight CLI command shortener and runner that helps you save, manage, and quickly run your frequently used terminal commands with optional arguments.

Think of it as a personal URL shortener—but for your terminal commands.

---

## Features

- Save complex or lengthy CLI commands with simple aliases
- Run saved commands instantly, optionally passing arguments
- Supports argument placeholders (`{}`) in saved commands for dynamic substitution
- List all saved commands in a clean table format
- Lightweight, easy to install and use
- Config and command data stored locally in `~/.term-cli/commands.json`

---

## Installation

1. Clone the repository:

    ```
    git clone https://github.com/srivastava-yash/term.git
    cd term
    ```
2. Build the binary:
    ```
    go build -o term
    ```
3. (Optional) Move binary to your PATH:
    ```
    mv term /usr/local/bin/
    ```

## Usage

### Save Command
Save a command with an alias:

    term save deploy "gcloud app deploy --project=my-app"

### List Command
Display saved commands in a neat table:

    term list

### Run a saved command
Run a saved command by alias, optionally passing arguments to replace {} placeholders:

    term run deploy

If your saved command includes {} placeholders, pass arguments to replace them in order:

    term save deploy "kubectl apply -f {} --namespace {} --record"
    term run deploy deployment.yaml production

## Confiuration and storage
- Commands are stored in JSON format at:
  - ~/.term-cli/commands.json

- You can manually edit this file, but be cautious to keep valid JSON.


## Contributing
Contributions, issues, and feature requests are welcome!
- Fork the repo
- Create your feature branch (git checkout -b feature/my-feature)
- Commit your changes (git commit -m 'Add some feature')
- Push to the branch (git push origin feature/my-feature)
- Open a Pull Request
