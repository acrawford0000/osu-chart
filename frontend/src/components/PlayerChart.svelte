<script>
    import { onMount, onDestroy } from 'svelte';
    import { players, selectedStat } from '../store';
    import { get } from 'svelte/store';
    import * as echarts from 'echarts/core';
    import { LineChart } from 'echarts/charts';
    import { GridComponent, TooltipComponent, LegendComponent } from 'echarts/components';
    import { CanvasRenderer } from 'echarts/renderers';
  
    echarts.use([LineChart, GridComponent, TooltipComponent, LegendComponent, CanvasRenderer]);
  
    let chartDom;
    let chart;
    let unsubscribe;
    let currentStat = $selectedStat;
  
    
    function updateChartData(value) {
        if (!chart) {
            // chart has not been initialized yet
            return;
        }

        // clear existing data
        chart.setOption({
            series: [],
        });
    
        // add new data
        const series = value.map((player, index) => ({
            name: player.username,
            type: 'line',
            smooth:true,
            sampling: 'lttb',
            data: player.stats.map((stat) => [stat.timestamp, stat[currentStat]]),
      }));
  
      chart.setOption({ series });
      console.log("Update qeued!");
    }
  
    onMount(() => {
      chart = echarts.init(chartDom, 'dark');
      chart.setOption({
        tooltip: {
          trigger: 'axis',
          axispointer: {
            snap: true,
          },
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
          snap: true,
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
  
      unsubscribe = players.subscribe(updateChartData);
    });
  
    onDestroy(() => {
      unsubscribe();
    });

    $: {
      currentStat = $selectedStat; 
      updateChartData(get(players));
    }
      </script>
  
  <div class="dark-mode" bind:this={chartDom} style="width: 70vw; height: 500px;"></div>
  