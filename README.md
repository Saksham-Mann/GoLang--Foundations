# GoLang Foundations

A dedicated repository documenting my journey through the Exercism Go Track. This project contains practical solutions, core Go concepts, and foundational exercises designed to build proficiency in Go.

---

## About The Project

This repository serves as a hands-on learning log for mastering Go. Every exercise focuses on specific language mechanics, including:

* **Go Basics and Syntax:** Package architecture, variables, constants, and basic types.
* **Control Flow and Functions:** Conditionals, loops, multiple return values, and error handling.
* **Data Structures:** Slices, arrays, maps, and structs.
* **Idiomatic Practices:** Pointers, methods, interfaces, and standard library conventions.

---

## Repository Structure

Solutions are automatically synced and organized by exercise topic under the `solutions/go/` path:

<!-- TREE_START -->
```text
.
├── .github/workflows/        # Automated merge workflows for Exercism sync
└── solutions/
    └── go/
        ├── annalyns-infiltration/
        ├── cars-assemble/
        ├── hello-world/
        ├── lasagna/
        └── weather-forecast/
```
<!-- TREE_END -->

Each exercise directory typically includes:
* The Go source file (`<exercise>.go`) containing the implemented solution.
* Test files provided by the Exercism test runner.

---

## Running Tests Locally

To run the unit tests for any specific exercise on your local machine:

1. Navigate to the exercise directory:
   ```bash
   cd solutions/go/<exercise-name>/1
   ```

2. Run tests using standard Go tooling:
   ```bash
   go test -v
   ```

---

## Automated Sync

Solutions are automatically backed up through the Exercism GitHub integration and merged into the `main` branch using GitHub Actions to keep the repository organized and up to date.
