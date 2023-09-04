<script>
  import { afterUpdate, onDestroy } from 'svelte';
  import * as echarts from 'echarts/core';
  import { LineChart } from 'echarts/charts';
  import { GridComponent, TooltipComponent, LegendComponent } from 'echarts/components';
  import { CanvasRenderer } from 'echarts/renderers';

  echarts.use([LineChart, GridComponent, TooltipComponent, LegendComponent, CanvasRenderer]);

  export let data = [];

  let chartDom;
  let chart;

  function createChart(value) {
      // If a chart already exists, dispose of it before creating a new one
      if (chart) {
          chart.dispose();
      }

      // Initialize the chart
      chart = echarts.init(chartDom, 'dark');
      chart.setOption({
          tooltip: {
              trigger: 'axis',
          },
          legend: {
              textStyle: {
                  color: '#fff',
              },
          },
          grid: {
              left: '3%',
              right: '4%',
              bottom: '3%',
              containLabel: true,
          },
          xAxis: {
              type: 'time',
              splitLine: {
                  show: false,
              },
              axisLabel: {
                  formatter: '{MMM}-{yyyy}',
                  textStyle: {
                      color: '#fff',
                  },
              },
          },
          yAxis: {
              type: 'value',
              axisLabel: {
                  textStyle: {
                      color: '#fff',
                  },
              },
          },
      });

      // Add data to the chart
      const series = value.map((player, index) => ({
          name: player.username,
          type: 'line',
          data: player.stats,
    }));

      chart.setOption({ series });
  }

  afterUpdate(() => {
      createChart(data);
  });

  onDestroy(() => {
      if (chart) chart.dispose();
  });
</script>

<div class="dark-mode" bind:this={chartDom} style="width: 70vw; height: 500px;"></div>
