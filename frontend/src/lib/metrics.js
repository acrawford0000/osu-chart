export const STAT_GROUPS = [
  { label: 'Performance', stats: [
    { key: 'pp_raw', label: 'Performance (pp)', description: 'Total performance points', decimals: 2 },
    { key: 'pp_rank', label: 'Global rank', description: 'Worldwide player ranking · lower is better', decimals: 0 },
    { key: 'accuracy', label: 'Accuracy (%)', description: 'Overall hit accuracy', decimals: 2, suffix: '%' },
    { key: 'level', label: 'Level', description: 'Player level', decimals: 2 }
  ]},
  { label: 'Activity', stats: [
    { key: 'playcount', label: 'Play count', description: 'Total plays', decimals: 0 },
    { key: 'ranked_score', label: 'Ranked score', description: 'Combined ranked score', decimals: 0 },
    { key: 'total_score', label: 'Total score', description: 'Combined score across all plays', decimals: 0 }
  ]},
  { label: 'Grades', stats: [
    { key: 'count_rank_ss', label: 'SS ranks', description: 'Number of SS grades', decimals: 0 },
    { key: 'count_rank_s', label: 'S ranks', description: 'Number of S grades', decimals: 0 },
    { key: 'count_rank_a', label: 'A ranks', description: 'Number of A grades', decimals: 0 }
  ]},
  { label: 'Hit results', stats: [
    { key: 'count300', label: '300 hits', description: 'Total 300 hit results', decimals: 0 },
    { key: 'count100', label: '100 hits', description: 'Total 100 hit results', decimals: 0 },
    { key: 'count50', label: '50 hits', description: 'Total 50 hit results', decimals: 0 }
  ]}
];
export const getStat = key => STAT_GROUPS.flatMap(group => group.stats).find(stat => stat.key === key) || STAT_GROUPS[2].stats[0];
export const GAME_MODES = [
  { key: 'standard', label: 'osu!' }, { key: 'taiko', label: 'osu!taiko' },
  { key: 'mania', label: 'osu!mania' }, { key: 'catch', label: 'osu!catch' }
];
