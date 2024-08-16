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

3. Make installer executable

   ```bash
   chmod +x ./install.sh
   ```

4. Run installer

   ```bash
   ./install.sh
   ```

## Usage

Once SAM is installed, you can use it as follows:

```bash
sam <command>
```

### Available Commands

- `commit`: Generates a commit message based on the changes in the Git repository.
- `improve`: Suggest improvements, fix bugs, and propose new features for the project
- `request`: Handle specific requests based on project analysis for example: `sam request "how to add feature x"`
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