# Exercism Solutions

A collection of my solutions to [Exercism](https://exercism.org/) exercises. The repository is primarily Go-focused, with selected JavaScript and TypeScript solutions.

## Repository Structure

```text
solutions/
├── go/
├── javascript/
└── typescript/
```

Each language directory contains folders for individual exercises. Some exercises also have numbered attempt directories, for example:

```text
solutions/go/bob/1/
solutions/go/bob/2/
```

These directories preserve multiple approaches or iterations of a solution.

## Running the Go Solutions

Install the [Go toolchain](https://go.dev/doc/install), change into an exercise directory, and run its tests:

```bash
cd solutions/go/<exercise>/<attempt>
go test ./...
```

Example:

```bash
cd solutions/go/bob/1
go test ./...
```

The solutions use standard Go packages and do not require a shared module or project-level build configuration.

## Language Breakdown

- **Go** — primary language and the largest collection of solutions
- **JavaScript** — selected Exercism exercises
- **TypeScript** — selected Exercism exercises

## Automation

The `.github/workflows/automerge.yml` workflow automatically squash-merges pull requests opened or reopened by the `exercism-solutions-syncer[bot]` account and deletes the merged branch.

## License

No license is currently specified for this repository.
