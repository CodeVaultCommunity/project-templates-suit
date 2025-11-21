# Blueprint Templates -- Projeto TypeScript + SonarQube + CI/CD

Este repositório fornece uma configuração completa e padronizada para
projetos TypeScript modernos, incluindo:

-   ESLint com formato SonarQube\
-   Vitest com cobertura e relatórios JUnit\
-   SonarQube Scanner com configurações dinâmicas via `.env`\
-   Husky + Commitlint para padronização de commits\
-   Workflows GitHub Actions\
-   Estrutura extensível para CI/CD corporativo e profissional

------------------------------------------------------------------------

## 📑 Sumário

-   [Recursos](#recursos)
-   [Como usar](#como-usar)
-   [Scripts disponíveis](#scripts-disponíveis)
-   [Arquivos importantes](#arquivos-importantes)
-   [Workflows GitHub](#workflows-github)
-   [Commits Convencionais](#commits-convencionais)
-   [Como rodar o SonarQube
    localmente](#como-rodar-o-sonarqube-localmente)
-   [Variáveis de ambiente](#variáveis-de-ambiente)
-   [Licença](#licença)

------------------------------------------------------------------------

## 📌 Recursos

### 🔧 Linting

-   ESLint configurado com regras modernas
-   Relatórios SonarQube (`eslint-report.json`)
-   Script de lint cross-platform (Windows + Linux)

### 🧪 Testes

-   Testes com **Vitest**
-   Relatórios JUnit para SonarQube
-   Cobertura configurada com V8

### 📊 Qualidade e Segurança

-   Análise estática via **SonarQube**
-   Scanner integrado e configurável via `.env`
-   Suporte total ao GitHub Actions

### 🔐 Commits

-   Husky ativando hooks
-   Commitlint garantindo convenções padrão
-   Hook automático para proteger o repositório

------------------------------------------------------------------------

## 🚀 Como usar

### 1. Instale dependências

``` bash
npm install
```

### 2. Inicialize o Husky

``` bash
npm run prepare
```

### 3. Crie seu arquivo `.env`

``` bash
SONAR_TOKEN=seu_token
SONAR_HOST_URL=http://localhost:9000
```

### 4. Execute toda a pipeline local

``` bash
npm run generate:reports
npm run sonar:scan
```

------------------------------------------------------------------------

## 📜 Scripts disponíveis

  -----------------------------------------------------------------------
  Script                          Descrição
  ------------------------------- ---------------------------------------
  `npm run lint`                  Executa o lint

  `npm run lint:report`           Gera relatório SonarQube (ESLint)

  `npm run test`                  Roda testes

  `npm run test:report`           Gera testes + cobertura

  `npm run generate:reports`      Executa lint + tests + coverage para o
                                  Sonar

  `npm run sonar:scan`            Executa o sonar com variáveis do `.env`

  `npm run husky:prepare`         Instala hooks do Husky
  -----------------------------------------------------------------------

------------------------------------------------------------------------

## 📁 Arquivos importantes

    .
    ├── sonar-project.properties
    ├── commitlint.config.js
    ├── .husky/
    ├── reports/
    └── .github/workflows/

------------------------------------------------------------------------

## 🤖 Workflows GitHub incluídos

-   **CI Full**\
-   **Lint & Test**\
-   **SonarQube Scan**\
-   **Build Only**\
-   **Release Draft (semântica)**

Todos estão modularizados em `.github/workflows/`.

------------------------------------------------------------------------

## 🧱 Commits Convencionais

Padrão utilizado:

    feat: nova funcionalidade
    fix: correção de bug
    docs: documentação
    chore: manutenção
    refactor: melhoria sem mudar comportamento
    test: testes
    ci: pipelines

------------------------------------------------------------------------

## 🛰️ Como rodar o SonarQube localmente

1.  Baixe o SonarQube Community Edition\
2.  Rode:\

``` bash
./bin/sonar start
```

3.  Crie um token\
4.  Coloque no `.env`

------------------------------------------------------------------------

## 🔧 Variáveis de ambiente

    SONAR_TOKEN=
    SONAR_HOST_URL=

------------------------------------------------------------------------

## 📄 Licença

MIT
