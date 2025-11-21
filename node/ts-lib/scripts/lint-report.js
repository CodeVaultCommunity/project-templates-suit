// scripts/lint-report.js
import { exec } from "child_process";
import { mkdirSync } from "fs";

mkdirSync("reports", { recursive: true });

const eslintCmd =
  "eslint src --format node_modules/eslint-formatter-sonarqube/index.js --output-file reports/eslint-report.json";

exec(eslintCmd, (error, stdout, stderr) => {
  if (stdout) console.log(stdout);
  if (stderr) console.error(stderr);

  if (error) {
    // ESLint find errors, but the report was generated normaly
    process.exit(0); // força sucesso
  }

  console.log("✔️  Relatório ESLint gerado com sucesso.");
  process.exit(0);
});

