import { defineConfig } from 'vitest/config';

export default defineConfig({
  test: {
    globals: true,
    environment: 'node',

    include: ['tests/**/*.test.ts'],
    exclude: ['dist/**', 'node_modules/**'],

    reporters: [
      'default',
      ['junit', { outputFile: 'reports/test-junit.xml' }]
    ],

    coverage: {
      provider: 'istanbul',
      reporter: ['text', 'lcov'],
      reportsDirectory: './coverage',
      exclude: ['tests/**']
    }
  }
});
