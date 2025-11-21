// scripts/run-sonar.js
import { execSync } from "child_process";
import dotenv from "dotenv";

// Carrega variáveis de ambiente do .env
dotenv.config();

const token = process.env.SONAR_TOKEN;
const host = process.env.SONAR_HOST_URL || "http://localhost:9000";

if (!token) {
  console.error("❌ ERROR: variable SONAR_TOKEN not defined");
  process.exit(1);
}

console.log("🔍 Running SonarQube...");
console.log(`➡ host: ${host}`);

try {
  execSync(
    `sonar-scanner -Dsonar.login=${token} -Dsonar.host.url=${host}`,
    { stdio: "inherit" }
  );
} catch (err) {
  console.error("❌ Failuer when execute sonar-scanner");
  process.exit(1);
}
