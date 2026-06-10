// Test setup file for Vitest
import { beforeAll, afterAll } from 'vitest';

// Global test setup
beforeAll(() => {
  // Mock window and document objects if needed
  console.log('Test environment initialized');
});

afterAll(() => {
  console.log('Tests completed');
});
