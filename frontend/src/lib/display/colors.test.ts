import { expect, describe, it } from 'vitest';
import { colorCodes, formatColors } from './colors.js';

describe('formatColors', () => {
  it('should surround the string with a span', () => {
    expect(formatColors('foo')).toBe('<span>foo</span>');
  });

  it('should represent ^#123456/^# pairs as color spans', () => {
    expect(formatColors('one ^#123456two^# three')).toBe("<span>one <span style='color: #123456'>two</span> three</span>");
  });

  it('should understand exported color code constant', () => {
    expect(formatColors(`^${colorCodes.MARAUDER}marauder^#`)).toBe("<span><span style='color: #E05030'>marauder</span></span>");
  });

  it('should close unmatched format specifiers at the end of the string', () => {
    expect(formatColors('one ^#123456two three')).toBe("<span>one <span style='color: #123456'>two three</span></span>");
  });

  it('should silently ignore unmatched close tags', () => {
    expect(formatColors('one ^#^#123456two^# three')).toBe(formatColors('one ^#123456two^# three'));
  });
});
