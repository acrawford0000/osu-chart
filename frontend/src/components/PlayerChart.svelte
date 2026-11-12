<script>
  import { onMount } from 'svelte';
  import { players, selectedStat } from '../store';
  import { findNearestPoint, preparePoints } from '../lib/chartData';
  import { getStat } from '../lib/metrics';
  import * as echarts from 'echarts/core';
  import { LineChart } from 'echarts/charts';
  import { AxisPointerComponent, GridComponent, TooltipComponent, LegendComponent, DataZoomInsideComponent } from 'echarts/components';
  import { CanvasRenderer } from 'echarts/renderers';

  echarts.use([LineChart, AxisPointerComponent, GridComponent, TooltipComponent, LegendComponent, CanvasRenderer, DataZoomInsideComponent]);

  let chartDom;
  let chart;
  let histories = [];
  let legendSelection = {};
  let hoverPoint = null;
  let hoverFrame = 0;
  const hoverDate = new Intl.DateTimeFormat(undefined, { year: 'numeric', month: 'short', day: 'numeric' });
  const savedDate = new Intl.DateTimeFormat(undefined, {
    year: 'numeric', month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit'
  });
  let statNumber = new Intl.NumberFormat(undefined, { maximumFractionDigits: 0 });
  let statSuffix = '';

  function hideTooltip() {
    cancelAnimationFrame(hoverFrame);
    hoverFrame = 0;
    hoverPoint = null;
    chart?.dispatchAction({ type: 'hideTip' });
  }

  function updateTooltip() {
    hoverFrame = 0;
    if (!chart || !hoverPoint) return;
    if (!chart.containPixel({ gridIndex: 0 }, hoverPoint)) {
      hideTooltip();
      return;
    }
    // Use the actual cursor date, not the date of ECharts' globally nearest series.
    const timestamp = chart.convertFromPixel({ xAxisIndex: 0 }, hoverPoint[0]);
    if (!Number.isFinite(timestamp)) return;
    const rows = histories.flatMap(history => {
      if (legendSelection[history.name] === false) return [];
      const point = findNearestPoint(history.points, timestamp);
      return point ? [{ ...history, point }] : [];
    });
    if (!rows.length) {
      hideTooltip();
      return;
    }

    // DOM text keeps player names as text, and ECharts still positions/confines the tooltip.
    const content = document.createElement('div');
    content.className = 'nearest-point-tooltip';
    const heading = document.createElement('div');
    heading.className = 'nearest-point-tooltip-heading';
    heading.textContent = hoverDate.format(timestamp);
    content.append(heading);
    for (const { name, color, point } of rows) {
      const row = document.createElement('div');
      row.className = 'nearest-point-tooltip-row';
      const marker = document.createElement('span');
      marker.className = 'nearest-point-tooltip-marker';
      marker.style.backgroundColor = color;
      const label = document.createElement('span');
      label.className = 'nearest-point-tooltip-name';
      label.textContent = name;
      const value = document.createElement('strong');
      value.className = 'nearest-point-tooltip-value';
      value.textContent = statNumber.format(point[1]) + statSuffix;
      const date = document.createElement('time');
      date.dateTime = new Date(point[0]).toISOString();
      date.textContent = savedDate.format(point[0]);
      row.append(marker, label, value, date);
      content.append(row);
    }
    chart.dispatchAction({
      type: 'showTip',
      x: hoverPoint[0],
      y: hoverPoint[1],
      tooltip: { formatter: () => content }
    });
  }

  function scheduleTooltip() {
    if (hoverPoint && !hoverFrame) hoverFrame = requestAnimationFrame(updateTooltip);
  }

  function handleMouseMove(event) {
    hoverPoint = [event.offsetX, event.offsetY];
    scheduleTooltip();
  }

  function updateChartData(value, metric) {
    if (!chart) return;
    hideTooltip();
    const definition = getStat(metric);
    statNumber = new Intl.NumberFormat(undefined, { maximumFractionDigits: definition.decimals });
    statSuffix = definition.suffix || '';
    const palette = chart.getOption().color;
    histories = value.map((player, index) => ({
      id: `${player.username}:${player.selectedMode || ''}:${index}`,
      name: player.username,
      color: palette[index % palette.length],
      points: preparePoints(player.stats, metric)
    }));
    const series = histories.map(history => ({
      id: history.id,
      name: history.name,
      type: 'line',
      smooth: true,
      sampling: 'lttb',
      symbolSize: 5,
      itemStyle: { color: history.color },
      data: history.points,
      emphasis: { focus: 'series' }
    }));
    // Replace the old series so removed players cannot linger in the plot or tooltip.
    chart.setOption({ series, yAxis: { inverse: metric === 'pp_rank' } }, { replaceMerge: ['series'] });
  }

  onMount(() => {
    chart = echarts.init(chartDom, 'dark');
    chart.setOption({
      backgroundColor: 'transparent',
      color: ['#a89bff', '#63d9b4', '#f3c575', '#76baff', '#f29db5', '#a9d576'],
      textStyle: { fontFamily: 'Chart UI, sans-serif' },
      tooltip: {
        // Native axis tooltips discard series at different timestamps. Supply our
        // own nearest-per-player content while retaining ECharts' tooltip layout.
        trigger: 'none',
        triggerOn: 'none',
        confine: true,
        backgroundColor: '#20232f',
        borderColor: '#3a3e52',
        padding: [13, 16],
        textStyle: { color: '#f2f3f5' },
        transitionDuration: 0,
        hideDelay: 0
      },
      axisPointer: { triggerOn: 'mousemove|click' },
      legend: { type: 'scroll', left: 24, right: 24, top: 20, icon: 'circle', itemWidth: 8, itemHeight: 8, itemGap: 22, textStyle: { color: '#b8bdd0', fontSize: 11 }, pageTextStyle: { color: '#959bb1' } },
      grid: { top: 65, left: 24, right: 26, bottom: 24, containLabel: true },
      xAxis: {
        type: 'time',
        axisPointer: { show: true, snap: false, triggerTooltip: false, animation: false, lineStyle: { color: '#74708e' }, label: { backgroundColor: '#343448', color: '#eeeff7' } },
        splitLine: { show: false },
        axisLabel: { formatter: '{MMM} {yyyy}', color: '#868fa8', fontSize: 10, hideOverlap: true, margin: 14 },
        axisLine: { lineStyle: { color: '#36394b' } },
        axisTick: { show: false }
      },
      yAxis: { type: 'value', axisLabel: { color: '#868fa8', fontSize: 10, formatter: value => new Intl.NumberFormat(undefined, { notation: 'compact', maximumFractionDigits: 1 }).format(value) }, splitLine: { lineStyle: { color: '#ffffff09' } } },
      dataZoom: [{ type: 'inside', xAxisIndex: [0], start: 0, end: 100 }]
    });

    const renderer = chart.getZr();
    renderer.on('mousemove', handleMouseMove);
    renderer.on('globalout', hideTooltip);
    const handleLegendChange = (event) => {
      legendSelection = event.selected;
      scheduleTooltip();
    };
    chart.on('legendselectchanged', handleLegendChange);
    chart.on('datazoom', scheduleTooltip);
    const resizeObserver = new ResizeObserver(() => {
      hideTooltip();
      chart.resize();
    });
    resizeObserver.observe(chartDom);
    document.fonts.ready.then(() => chart?.resize());

    return () => {
      cancelAnimationFrame(hoverFrame);
      resizeObserver.disconnect();
      renderer.off('mousemove', handleMouseMove);
      renderer.off('globalout', hideTooltip);
      chart.off('legendselectchanged', handleLegendChange);
      chart.off('datazoom', scheduleTooltip);
      chart.dispose();
      chart = null;
    };
  });

  $: if (chart) updateChartData($players, $selectedStat);
</script>

<div class="dark-mode" bind:this={chartDom}></div>

<style>
  .dark-mode { width: 100%; height: 440px; }
  @media(max-width:600px) { .dark-mode { height: 380px; } }
  :global(.nearest-point-tooltip-heading) {
    margin-bottom: 8px;
  }
  :global(.nearest-point-tooltip-row) {
    display: grid;
    grid-template-columns: 10px minmax(90px, 1fr) auto;
    column-gap: 6px;
    align-items: center;
    margin-top: 6px;
  }
  :global(.nearest-point-tooltip-marker) {
    width: 10px;
    height: 10px;
    border-radius: 50%;
  }
  :global(.nearest-point-tooltip-value) {
    margin-left: 12px;
    text-align: right;
  }
  :global(.nearest-point-tooltip-row time) {
    grid-column: 2 / 4;
    font-size: 11px;
    opacity: 0.7;
  }
</style>
