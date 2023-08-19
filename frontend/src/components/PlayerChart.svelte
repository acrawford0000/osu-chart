<script>
    import { onMount, onDestroy } from 'svelte';
    import { players, selectedStat } from '../store';
    import * as echarts from 'echarts/core';
    import { LineChart } from 'echarts/charts';
    import { GridComponent, TooltipComponent, LegendComponent } from 'echarts/components';
    import { CanvasRenderer } from 'echarts/renderers';
    import { use } from 'echarts/core';
    import * as dayjs from 'dayjs';

    use([LineChart, GridComponent, TooltipComponent, LegendComponent, CanvasRenderer]);

    let chartDiv;
    let chart;
    let unsubscribePlayers;
    let unsubscribeSelectedStat;

    onMount(() => {
        // Initialize chart
        chart = echarts.init(chartDiv, 'dark');
        updateChart();

        // Subscribe to players store and update chart when it changes
        unsubscribePlayers = players.subscribe(() => {
            updateChart();
        });

        // Subscribe to selectedStat store and update chart when it changes
        unsubscribeSelectedStat = selectedStat.subscribe(() => {
            updateChart();
        });

        // Add resize event listener to window to resize chart when window is resized
        window.addEventListener('resize', () => {
            chart.resize();
        });
    });

    function updateChart() {
        // Get data from players store
        let allPlayers = [];
        const unsubscribe = players.subscribe(currentPlayers => {
            allPlayers = currentPlayers;
        });
        unsubscribe();

        // Get selected stat from selectedStat store
        let currentSelectedStat = '';
        const unsubscribe2 = selectedStat.subscribe(currentSelectedStat => {
            currentSelectedStat = currentSelectedStat;
        });
        unsubscribe2();

        // Map data to format needed for chart
        const series = allPlayers.map(player => ({
            name: player.username,
            type: 'line',
            data: player.stats.map(stat => stat[currentSelectedStat])
        }));

        // Set chart options
        const option = {
            tooltip: {
                trigger: 'axis'
            },
            legend: {
                data: allPlayers.map(player => player.username)
            },
            xAxis: {
                type: 'category',
                data: allPlayers[0]?.stats.map(stat => dayjs(stat.timestamp).format('YYYY-MM-DD HH:mm:ss'))
            },
            yAxis: {
                type: 'value'
            },
            series
        };
        
        // Set option and resize chart
        chart.setOption(option);
        chart.resize();
    }

    // Clean up subscriptions when component is destroyed
    onDestroy(() => {
        unsubscribePlayers();
        unsubscribeSelectedStat();
    });
</script>

<div bind:this={chartDiv} style="width: 100%; height: 100%;"></div>
