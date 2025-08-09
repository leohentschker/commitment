# Committing CLI

A command-line tool built with Bubble Tea to help create conventional branches and commits.

## Installation

```bash
go build -o committing
```

### Add to PATH

To use `committing` from anywhere, add it to your shell profile:

**For bash (`~/.bashrc` or `~/.bash_profile`):**
```bash
echo 'export PATH="$PATH:$(pwd)"' >> ~/.bashrc
source ~/.bashrc
```

**For zsh (`~/.zshrc`):**
```bash
echo 'export PATH="$PATH:$(pwd)"' >> ~/.zshrc
source ~/.zshrc
```

**For fish (`~/.config/fish/config.fish`):**
```bash
echo 'set -gx PATH $PATH (pwd)' >> ~/.config/fish/config.fish
source ~/.config/fish/config.fish
```

After adding to PATH, you can use `committing` instead of `./committing`.

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
