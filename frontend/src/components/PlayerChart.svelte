<script>
    import { onMount, onDestroy } from 'svelte';
    import { players, selectedStat } from '../store';
    import * as echarts from 'echarts/core';
    import { LineChart } from 'echarts/charts';
    import { GridComponent, TooltipComponent, LegendComponent } from 'echarts/components';
    import { CanvasRenderer } from 'echarts/renderers';
    import { use } from 'echarts/core';
    import { get } from 'svelte/store';

    use([LineChart, GridComponent, TooltipComponent, LegendComponent, CanvasRenderer]);

    let chartDiv;
    let chart;
    let unsubscribePlayers;
    let unsubscribeSelectedStat;

    // Get data from players store
    let allPlayers = [];
    players.subscribe(value => {
        allPlayers = value;
        updateChart();
    });

    onMount(() => {
        // Initialize chart
        chart = echarts.init(chartDiv, 'dark');
        updateChart();

/*         // Subscribe to players store and update chart when it changes
        unsubscribePlayers = players.subscribe(() => {
            updateChart();
        }); */

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

        // Check if allPlayers array is empty
        if (allPlayers.length === 0) {
            // If allPlayers array is empty, clear chart and return
            chart.clear();
            return;
        }

        // Get selected stat from selectedStat store
        const currentSelectedStat = get(selectedStat);

        // Set up series data for each player
        const series = allPlayers.map(player => ({
            name: player.username,
            type: 'line',
            data: player.stats.map(stat => [stat.timestamp, stat[currentSelectedStat]])
        }));

        // Set chart options
        const option = {
            tooltip: {
                trigger: 'axis'
            },
            legend: {},
            xAxis: {
                type: 'time',
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

<div bind:this={chartDiv} style="width: 100%; height: 500px;"></div>
<div>
    Selected Stat: {$selectedStat}
</div>