[<img src="LLM" align="right" width="25%" padding-right="350">]()

# `PLATEFORME-MYS3`

#### <code>❯ Une plateforme simplifiée compatible avec S3 pour la gestion des fichiers</code>

<p align="left">
	<img src="https://img.shields.io/github/license/tchessi-pre/plateforme-mys3?style=flat&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/tchessi-pre/plateforme-mys3?style=flat&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/tchessi-pre/plateforme-mys3?style=flat&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/tchessi-pre/plateforme-mys3?style=flat&color=0080ff" alt="repo-language-count">
</p>
<p align="left">
		<em>Construit avec les outils et technologies suivants :</em>
</p>
<p align="center">
	<img src="https://img.shields.io/badge/YAML-CB171E.svg?style=flat&logo=YAML&logoColor=white" alt="YAML">
	<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=flat&logo=Go&logoColor=white" alt="Go">
</p>

<br>

##### 🔗 Table des Matières

- [📍 Vue d'ensemble](#-vue-densemble)
- [👾 Fonctionnalités](#-fonctionnalités)
- [📂 Structure du dépôt](#-structure-du-dépôt)
- [🧩 Modules](#-modules)
- [🚀 Pour commencer](#-pour-commencer)
  - [🔖 Prérequis](#-prérequis)
  - [📦 Installation](#-installation)
  - [🤖 Utilisation](#-utilisation)
  - [🧪 Tests](#-tests)
  - [📝 Gestion des buckets et des fichiers avec `mc`](#-gestion-des-buckets-et-des-fichiers-avec-mc)
- [📌 Roadmap du projet](#-roadmap-du-projet)
- [🤝 Contribuer](#-contribuer)
- [🎗 Licence](#-licence)
- [🙌 Remerciements](#-remerciements)

---

## 📍 Vue d'ensemble

<code>❯ Cette plateforme fournit un système de stockage de fichiers compatible S3, développé avec Go et MinIO.</code>

---

## 👾 Fonctionnalités

- Créer et gérer des buckets compatibles S3.
- Upload, téléchargement et suppression de fichiers.
- Support de stockage local en parallèle avec MinIO.
- API REST simple pour les opérations sur les fichiers.

---

## 📂 Structure du dépôt

```sh
└── plateforme-mys3/
    ├── app
    │   ├── go.mod
    │   ├── go.sum
    │   ├── handlers
    │   ├── main.go
    │   ├── storage (Bucket local)
    │   └── tests
    ├── docker-compose.yml
    └── readme.md
```

---

## 🧩 Modules

<details closed><summary>.</summary>

| Fichier                  | Résumé                               |
|--------------------------|--------------------------------------|
| `docker-compose.yml`      | Configuration Docker pour MinIO et Go|

</details>

<details closed><summary>app</summary>

| Fichier       | Résumé                                     |
|---------------|--------------------------------------------|
| `go.sum`      | Gestion des dépendances                    |
| `go.mod`      | Suivi des versions des dépendances         |
| `main.go`     | Point d'entrée principal de l'application Go|

</details>

<details closed><summary>app.handlers</summary>

| Fichier                  | Résumé                                |
|--------------------------|---------------------------------------|
| `create_bucket.go`        | Point d'API pour la création de buckets|
| `delete_file.go`          | API pour la suppression de fichiers   |
| `download_file.go`        | API pour le téléchargement de fichiers|
| `upload_file.go`          | API pour l'upload de fichiers         |
| `list_files.go`           | API pour lister les fichiers dans un bucket|

</details>

<details closed><summary>app.storage</summary>

| Fichier        | Résumé                    |
|----------------|---------------------------|
| `storage.go`   | Gestion du stockage local  |

</details>

<details closed><summary>app.storage</summary>

| File                                                                                                                                                                                                          | Summary                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
| [storage.go](https://github.com/tchessi-pre/plateforme-mys3/blob/main/app/storage/storage.go)                                                                                                                 | <code>❯ Documentation détaillée pour la gestion de buckets et de fichiers dans MinIO à l'aide de mc (MinIO Client) et un serveur Go custom |
| Cette documentation vous guidera à travers les différentes étapes pour gérer des buckets et des fichiers sur MinIO avec l'outil mc ainsi qu'un serveur Go qui utilise MinIO comme backend de stockage.</code> |

</details>

<details closed><summary>app.storage.myBucket3</summary>

| File                                                                                                | Summary                   |
| --------------------------------------------------------------------------------------------------- | ------------------------- |
| [myObject](https://github.com/tchessi-pre/plateforme-mys3/blob/main/app/storage/myBucket3/myObject) | <code>❯ REPLACE-ME</code> |

</details>

---

## 🚀 Getting Started

### 🔖 Prerequisites

**Go**: `version latest`

### 📦 Installation

Build the project from source:

1. Clone the plateforme-mys3 repository:

```sh
❯ git clone https://github.com/tchessi-pre/plateforme-mys3
```

2. Navigate to the project directory:

```sh
❯ cd plateforme-mys3
```

3. Install the required dependencies:

```sh
❯ go build -o myapp
```

### 🤖 Usage

To run the project, execute the following command:

```sh
❯ ./myapp
```

### 🧪 Tests

Execute the test suite using the following command:

```sh
❯ go test
```

---

## 📌 Project Roadmap

- [x] **`Task 1`**: <strike>Implement feature one.</strike>
- [ ] **`Task 2`**: Implement feature two.
- [ ] **`Task 3`**: Implement feature three.

---

## 🤝 Contributing

Contributions are welcome! Here are several ways you can contribute:

- **[Report Issues](https://github.com/tchessi-pre/plateforme-mys3/issues)**: Submit bugs found or log feature requests for the `plateforme-mys3` project.
- **[Submit Pull Requests](https://github.com/tchessi-pre/plateforme-mys3/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.
- **[Join the Discussions](https://github.com/tchessi-pre/plateforme-mys3/discussions)**: Share your insights, provide feedback, or ask questions.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your github account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
   ```sh
   git clone https://github.com/tchessi-pre/plateforme-mys3
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to github**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your contribution!
</details>

<details closed>
<summary>Contributor Graph</summary>
<br>
<p align="left">
   <a href="https://github.com{/tchessi-pre/plateforme-mys3/}graphs/contributors">
      <img src="https://contrib.rocks/image?repo=tchessi-pre/plateforme-mys3">
   </a>
</p>
</details>

---

## 🎗 License

This project is protected under the [SELECT-A-LICENSE](https://choosealicense.com/licenses) License. For more details, refer to the [LICENSE](https://choosealicense.com/licenses/) file.

---

## 🙌 Acknowledgments

- List any resources, contributors, inspiration, etc. here.

---
