# TypeScript Starter — Linting, SonarQube, Commitlint & GitHub Workflows

A complete TypeScript boilerplate including:
- ESLint with custom rules
- SonarQube reporting (XML + JSON dynamic handling)
- Cross‑platform lint report generation (Windows + Unix)
- Commitlint + Husky Git hooks
- GitHub Actions templates for CI/CD
- Preconfigured scripts for local development

---

## 📚 Table of Contents
- [Features](#features)
- [Scripts](#scripts)
- [Setup](#setup)
- [GitHub Workflows](#github-workflows)
- [Português (PT-BR)](#pt-br)

---

## Features
- Strict TypeScript config  
- ESLint + Prettier integration  
- Dynamic lint report generator  
- SonarQube environment‑driven token loading  
- Conventional commit enforcement  
- GitHub workflows: CI, sonar scan, test matrix, PR checks, etc.

---

## Scripts
```bash
npm run build
npm run lint
npm run lint:report
npm run test
npm run sonar
```

---

## Setup

### 1. Install dependencies
```bash
npm install
```

### 2. Configure environment variables
Create a `.env` file:

```
SONAR_TOKEN=your_token_here
SONAR_HOST_URL=http://localhost:9000
```

### 3. Enable Husky
```bash
npx husky
```

---

## GitHub Workflows Included
- **CI Build & Lint**
- **SonarQube Cloud Scan**
- **Pull Request Checks**
- **Test Matrix**
- **Lint‑Only Workflow**
- **Auto‑Formatting on PR**

---

## PT-BR
Leia esta documentação em português aqui:  
[LEIA-ME](./LEIA-ME.md)
