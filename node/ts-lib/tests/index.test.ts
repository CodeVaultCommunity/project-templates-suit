import { describe, it, expect } from 'vitest';
import { soma } from '../src';

describe('soma', () => {
  it('soma dois números', () => {
    expect(soma(2, 3)).toBe(5);
  });
});
    