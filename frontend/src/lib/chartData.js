/**
 * Prepare the complete history once, independent of ECharts' display sampling.
 * Missing metrics are not zeroes, and API responses need not be chronological.
 * @param {Array<Record<string, any>> | null | undefined} stats
 * @param {string} metric
 * @returns {Array<[number, number]>}
 */
export function preparePoints(stats, metric) {
  if (!Array.isArray(stats)) return [];
  return stats.flatMap(stat => {
    const timestamp = stat?.timestamp;
    const raw = stat?.[metric];
    // osu!track serializes score totals as decimal strings.
    const value = typeof raw === 'string' && /^\d+(\.\d+)?$/.test(raw) ? Number(raw) : raw;
    if ((typeof timestamp !== 'string' && typeof timestamp !== 'number') ||
        timestamp === '' || typeof value !== 'number' || !Number.isFinite(value)) return [];
    const time = new Date(timestamp).getTime();
    return Number.isFinite(time) ? [/** @type {[number, number]} */ ([time, value])] : [];
  }).sort((a, b) => a[0] - b[0]);
}

/**
 * Find the nearest observation in a sorted history in O(log n).
 * Equal distances choose the earlier observation; endpoints remain available.
 * @param {Array<[number, number]>} points
 * @param {number} timestamp
 * @returns {[number, number] | null}
 */
export function findNearestPoint(points, timestamp) {
  if (!points.length || !Number.isFinite(timestamp)) return null;
  let low = 0;
  let high = points.length;
  while (low < high) {
    const middle = (low + high) >>> 1;
    if (points[middle][0] < timestamp) low = middle + 1;
    else high = middle;
  }
  if (low === 0) return points[0];
  if (low === points.length) return points[low - 1];
  const before = points[low - 1];
  const after = points[low];
  return timestamp - before[0] <= after[0] - timestamp ? before : after;
}
