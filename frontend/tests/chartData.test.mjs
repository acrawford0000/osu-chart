import test from 'node:test';
import assert from 'node:assert/strict';
import { preparePoints, findNearestPoint } from '../src/lib/chartData.js';
const day = value => Date.parse(`2020-10-${value}T00:00:00Z`);

test('finds each player independently when their dates do not overlap', () => {
  const a = [[day('01'), 10], [day('05'), 50]];
  const b = [[day('03'), 30], [day('09'), 90]];
  assert.deepEqual(findNearestPoint(a, day('04')), [day('05'), 50]);
  assert.deepEqual(findNearestPoint(b, day('04')), [day('03'), 30]);
});
test('ties select the earlier observation; exact matches and endpoints work', () => {
  const points = [[day('01'), 0], [day('03'), 3], [day('09'), 9]];
  assert.deepEqual(findNearestPoint(points, day('02')), points[0]);
  assert.deepEqual(findNearestPoint(points, day('03')), points[1]);
  assert.deepEqual(findNearestPoint(points, day('01') - 1e9), points[0]);
  assert.deepEqual(findNearestPoint(points, day('09') + 1e9), points[2]);
  assert.equal(findNearestPoint([], day('01')), null);
  assert.equal(findNearestPoint(points, NaN), null);
});
test('prepares chronological points without mutating history or inventing zeroes', () => {
  const stats = [
    { timestamp: '2020-10-03T00:00:00Z', rank: 3, pp: 30 },
    { timestamp: '2020-10-01T00:00:00Z', rank: 0, pp: 10 },
    { timestamp: '2020-10-02T00:00:00Z', rank: null },
    { timestamp: 'invalid', rank: 100 },
    { timestamp: null, rank: 100 },
    { timestamp: '2020-10-04T00:00:00Z', rank: Infinity }
  ];
  assert.deepEqual(preparePoints(stats, 'rank'), [[day('01'), 0], [day('03'), 3]]);
  assert.deepEqual(preparePoints(stats, 'pp'), [[day('01'), 10], [day('03'), 30]]);
  assert.equal(stats[0].rank, 3);
  assert.deepEqual(preparePoints(null, 'rank'), []);
});
test('keeps full-resolution observations for the tooltip even in dense histories', () => {
  const points = Array.from({length:100000}, (_, i) => [i * 1000, i]);
  assert.deepEqual(findNearestPoint(points, 45678123), [45678000, 45678]);
});

test('score totals serialized as strings plot correctly without treating missing values as zero', () => {
  const stats = [
    {timestamp:'2020-10-01T00:00:00Z',ranked_score:'1234567890123'},
    {timestamp:'2020-10-02T00:00:00Z',ranked_score:''},
    {timestamp:'2020-10-03T00:00:00Z',ranked_score:null},
    {timestamp:'2020-10-04T00:00:00Z',ranked_score:'0'}
  ];
  assert.deepEqual(preparePoints(stats,'ranked_score'), [[day('01'),1234567890123],[day('04'),0]]);
});
