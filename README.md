# Committing CLI

A command-line tool built with Bubble Tea to help create conventional branches and commits.

## Installation

```bash
go build -o committing
```

### Add aliases to shell profile
Add these aliases to your shell profile for convenient access:
```bash
alias br="$(pwd)/committing br"
alias cm="$(pwd)/committing cm" 
alias st="$(pwd)/committing status"
alias pr="$(pwd)/committing pr"
alias fe="$(pwd)/committing fe"
```

Then reload your shell: `source ~/.zshrc` or `source ~/.bashrc`

## Usage

### Create a new branch
```bash
./committing br
```

### Create a commit
```bash
./committing cm
```

## Flow

Both commands follow the same interactive flow:

1. **Select Type**: Choose from conventional commit types:
   - `feat` - A new feature
   - `fix` - A bug fix
   - `docs` - Documentation only changes
   - `style` - Changes that do not affect the meaning of the code
   - `refactor` - A code change that neither fixes a bug nor adds a feature
   - `perf` - A code change that improves performance
   - `test` - Adding missing tests or correcting existing tests
   - `build` - Changes that affect the build system or external dependencies
   - `ci` - Changes to CI configuration files and scripts
   - `chore` - Other changes that don't modify src or test files

2. **Enter Description**: Write a brief description of the change

3. **Confirm**: Review and confirm the final branch name or commit message

## Examples

### Branch Creation
- Input: `feat` + "add user authentication"
- Result: Creates branch `feat/add-user-authentication`

### Commit Creation  
- Input: `fix` + "resolve memory leak in parser"
- Result: Creates commit with message `fix: resolve memory leak in parser`

## Navigation

- Use arrow keys or `j`/`k` to navigate
- Press `Enter` to select/confirm
- Press `q` or `Ctrl+C` to quit/cancel
