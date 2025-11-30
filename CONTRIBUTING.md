# Contributing to High-Tech Job Interview Study Course

Thank you for your interest in contributing! This is an open project dedicated to helping developers prepare for high-tech job interviews. We welcome contributions of all forms, including new exercises, bug fixes, better explanations, and system design scenarios.

## Getting Started

1.  **Fork the repository**: Click the "Fork" button at the top right of this page.
2.  **Clone your fork**:
    ```bash
    git clone https://github.com/YOUR_USERNAME/study.git
    cd study
    ```
3.  **Create a branch**: Use a descriptive name for your branch.
    ```bash
    git checkout -b feature/add-new-exercise
    ```

## How to Contribute

### Adding a New Exercise

1.  Navigate to the appropriate category (e.g., `01-data-structures/04-stacks`).
2.  Create a new directory for your exercise (e.g., `05-my-new-exercise`).
3.  Add a `main.go` file with the problem description, a `Solution` function, and a `main` function with test cases.
4.  Add an `answer.md` file explaining the approach and complexity ($O$ notation).
5.  Ensure your code runs and passes tests: `go run main.go`.

### Improving Explanations

- Edit the `explanation.md` files in the category directories.
- Use clear, concise language.
- Add Go code snippets or diagrams (Mermaid or images) where helpful.

### Reporting Issues

- If you find a bug or have a suggestion, please open an issue on GitHub.
- Provide as much detail as possible.

## Coding Standards

- **Language**: Go (Golang).
- **Formatting**: Run `go fmt` on your code before committing.
- **Style**: Keep code clean and readable. Variable names should be descriptive.
- **Comments**: Add comments to explain complex logic.
- **Self-Contained**: Each exercise's `main.go` should be self-contained and runnable.

## Pull Request Process

1.  Push your changes to your fork:
    ```bash
    git push origin feature/add-new-exercise
    ```
2.  Open a Pull Request (PR) from your branch to the `main` branch of this repository.
3.  Describe your changes in the PR description.
4.  Wait for review. We may ask for changes before merging.

## Code of Conduct

Please be respectful and considerate of others. We want to create a welcoming environment for everyone learning and sharing knowledge.

Happy Coding!
