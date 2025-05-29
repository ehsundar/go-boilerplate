# Commit Message Convention

To maintain a clean and readable project history, please follow these commit message rules:

## Commit Message Structure

Each commit message must consist of a **header**, an optional **body**, and an optional **footer**:

```
type(scope?): subject

body (optional)

footer (optional)
```

### Header
- **type**: The type of change (see below)
- **scope**: The section of the codebase affected (optional)
- **subject**: A short description of the change (max 72 characters, no period at end)

### Body (optional)
- Use to explain the motivation and context for the change.
- Wrap lines at 72 characters.

### Footer (optional)
- Use for breaking changes or issues closed (e.g., `BREAKING CHANGE: ...` or `Closes #123`).

## Allowed Types
- **feat**: A new feature
- **fix**: A bug fix
- **docs**: Documentation only changes
- **style**: Changes that do not affect meaning (formatting, etc.)
- **refactor**: Code change that neither fixes a bug nor adds a feature
- **perf**: Performance improvement
- **test**: Adding or correcting tests
- **chore**: Maintenance tasks

## Using Emojis

You are encouraged to add a relevant emoji at the start of the subject line to make commit messages more visually engaging. Place the emoji after the type (and optional scope). For example:

- **feat**: ✨, 🚀, 🆕
- **fix**: 🐛, 🛠️, 🚑️
- **docs**: 📝, 📚
- **style**: 💄, 🎨
- **refactor**: ♻️, 🔨
- **perf**: ⚡, 🚤
- **test**: ✅, 🧪
- **chore**: 🔧, ⬆️

Choose emojis that best represent the change. This is optional but encouraged.

## Examples

```
feat(api): add user authentication endpoint

fix(items): correct item deletion logic

docs: update README with setup instructions

refactor(storage): simplify query builder

chore: update dependencies
```

Thank you for following these guidelines to keep the project history clean and useful.
