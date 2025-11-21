// eslint.config.js
import js from "@eslint/js";
import tseslint from "typescript-eslint";
import prettier from "eslint-config-prettier";
import importPlugin from "eslint-plugin-import";

export default [
  // ignorar dist / node_modules
  {
    ignores: ["dist/**", "node_modules/**"]
  },

  // configurações base do ESLint para JS
  js.configs.recommended,

  // recommended do typescript-eslint (spread das configs internas)
  ...tseslint.configs.recommended,

  // Regras específicas para TS/TSX
  {
    files: [
      "src/**/*.{ts,tsx}",
      // "tests/**/*.{ts,tsx}"
    ],
    plugins: { import: importPlugin },
    languageOptions: {
      parser: tseslint.parser,
      parserOptions: {
        project: ["./tsconfig.json"],
        sourceType: "module"
      }
    },
    rules: {
      // ===== Naming conventions =====
      "@typescript-eslint/naming-convention": [
        "error",
        // classes / interfaces / type aliases / enums -> PascalCase
        {
          selector: "typeLike",
          format: ["PascalCase"]
        },

        // functions -> camelCase (inclui function declarations)
        {
          selector: "function",
          format: ["camelCase"]
        },

        // methods -> camelCase
        {
          selector: "method",
          format: ["camelCase"]
        },

        // variables -> snake_case OR UPPER_CASE (consts)
        {
          selector: "variable",
          format: ["snake_case", "UPPER_CASE"],
          leadingUnderscore: "allow",
          trailingUnderscore: "allow"
        },

        // parameters -> camelCase (boas práticas)
        {
          selector: "parameter",
          format: ["camelCase"],
          leadingUnderscore: "allow"
        },

        // propriedades de objetos -> permitir qualquer formato (API/JSON)
        {
          selector: "property",
          format: null
        }
      ],

      // ===== Brace style (Allman) =====
      // coloca '{' em linha separada:
      "brace-style": ["error", "allman", { "allowSingleLine": false }],

      // ===== Import order =====
      // exige plugin eslint-plugin-import estar instalado
      "import/order": [
        "error",
        {
          groups: [
            "builtin",
            "external",
            "internal",
            ["parent", "sibling", "index"]
          ],
          "newlines-between": "always",
          alphabetize: { order: "asc", caseInsensitive: true }
        }
      ],

      // ===== Proibições / boas práticas TS =====
      "@typescript-eslint/no-explicit-any": "error", // proíbe any
      "@typescript-eslint/no-unused-vars": [
        "error",
        { vars: "all", args: "after-used", ignoreRestSiblings: true }
      ],
      "@typescript-eslint/explicit-function-return-type": [
        "warn",
        { allowExpressions: true, allowTypedFunctionExpressions: true }
      ],

      // ===== Espaçamento / organização =====
      "lines-between-class-members": ["error", "always", { exceptAfterSingleLine: false }],
      "padding-line-between-statements": [
        "error",
        // blank line before return
        { blankLine: "always", prev: "*", next: "return" },
        // blank lines around exports
        { blankLine: "always", prev: "*", next: "export" }
      ],

      // ===== Estilo geral =====
      "no-tabs": "error",
      "no-console": ["warn", { allow: ["warn", "error"] }],
      "prefer-const": ["error", { destructuring: "all" }]
    }
  },

  // Prettier por último para desabilitar regras conflituosas
  prettier
];
