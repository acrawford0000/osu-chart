<script>
  export let name = 'chart';
  export let size = 20;
  const paths = {
    players: 'M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2 M16 3a4 4 0 0 1 0 8 M22 21v-2a4 4 0 0 0-3-3.87 M13 7a4 4 0 1 1-8 0 4 4 0 0 1 8 0',
    chart: 'M3 3v18h18 M7 14l4-5 4 3 6-8',
    settings: 'M9 3h6l1 3 3 1 2 5-2 5-3 1-1 3H9l-1-3-3-1-2-5 2-5 3-1 1-3 M15 12a3 3 0 1 1-6 0 3 3 0 0 1 6 0',
    contact: 'M21 11.5a8.4 8.4 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.4 8.4 0 0 1-3.8-.9L3 21l1.9-5.7a8.4 8.4 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.4 8.4 0 0 1 3.8-.9h.5a8.5 8.5 0 0 1 8 8v.5',
    plus: 'M12 5v14 M5 12h14', close: 'M6 6l12 12 M6 18L18 6',
    arrow: 'M5 12h14 M13 6l6 6-6 6', check: 'M5 12l4 4L19 6',
    eye: 'M2 12s3-7 10-7 10 7 10 7-3 7-10 7-10-7-10-7 M15 12a3 3 0 1 1-6 0 3 3 0 0 1 6 0',
    lock: 'M5 11h14v10H5z M8 11V7a4 4 0 0 1 8 0v4',
    external: 'M15 3h6v6 M10 14L21 3 M21 14v7H3V3h7',
    github: 'M9 19c-4 1-4-2-6-2 M15 22v-4a3.5 3.5 0 0 0-1-3c3 0 6-1.5 6-5a4 4 0 0 0-1-3 4 4 0 0 0 0-3s-1 0-3 1a11 11 0 0 0-6 0C8 4 7 4 7 4a4 4 0 0 0 0 3 4 4 0 0 0-1 3c0 3.5 3 5 6 5a3.5 3.5 0 0 0-1 3v4',
    video: 'M9 8l7 4-7 4V8 M3 5h18v14H3V5', copy: 'M9 9h12v12H9z M15 9V3H3v12h6'
  };
</script>
<svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.7" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d={paths[name] || paths.chart}/></svg>
