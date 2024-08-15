# SAM - AI CLI Tool

SAM is a CLI tool that generates commit messages based on the changes in a Git repository using OpenAI's GPT-4o-mini language model.

## Installation

### Prerequisites

- Go (version 1.16 or later)
- Git

### Steps

1. Clone the SAM repository:

   ```bash
   git clone https://github.com/blackestwhite/sam.git
   ```

2. Change to the SAM directory:

   ```bash
   cd sam
   ```

3. Build the SAM binary:

   ```bash
   go build
   ```

4. Move the SAM binary to a directory in your PATH (e.g., `/usr/local/bin`):

   ```bash
   sudo mv sam /usr/local/bin
   ```

## Usage

Once SAM is installed, you can use it as follows:

```bash
sam <command>
```

### Available Commands

- `commit`: Generates a commit message based on the changes in the Git repository.
- `improve`: Suggest improvements, fix bugs, and propose new features for the project
- `help`: Displays usage information.

## Configuration

SAM requires an API key from OpenAI to generate commit messages. Before using SAM, you need to set up your API key by creating a `.samrc` file in your home directory (`$HOME`) and adding your API key to it.

Example `.samrc` file:

```
OPENAI_API_KEY="sk-xxxxxxxxxxxxxxxxxx"
```

## Donations

Ethereum: blackestwhite.eth

TON(Telegram Open Network): blackestwhite.ton