/** Wails may reject with a string rather than an Error instance. */
export function normalizeError(error, fallback = 'Something went wrong. Please try again.') {
  const text = typeof error === 'string' ? error : error?.message;
  if (typeof text !== 'string' || !text.trim()) return { code: 'UNKNOWN', message: fallback };
  const match = text.match(/^([A-Z_]+):\s*([\s\S]*)$/);
  return match ? { code: match[1], message: match[2] || fallback } : { code: 'UNKNOWN', message: text };
}
