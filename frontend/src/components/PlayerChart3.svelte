<script lang="ts">
    import StatSelectorV2 from './StatSelectorV2.svelte';
    import { onMount, onDestroy } from 'svelte';
    import { players, selectedStat } from '../store';
    import { get } from 'svelte/store';
    import 'chartjs-adapter-date-fns';
    import zoomPlugin from 'chartjs-plugin-zoom';
    import { Chart, Legend, LineController, LineElement, PointElement, LinearScale, CategoryScale, TimeScale, Tooltip } from 'chart.js';
    Chart.register(Legend, LineController, LineElement, PointElement, LinearScale, CategoryScale, TimeScale, Tooltip);

    let canvas;
    let chart;
    let unsubscribe;
    let showAll = false;



    function updateChartData(value) {
        // clear existing data
        chart.data.labels = [];
        chart.data.datasets = [];

        // add new data
        value.forEach((player, index) => {
            let data = player.stats.map(stat => ({
                x: stat.timestamp,
                y: stat[get(selectedStat)]
            }));

            if (chart.data.datasets[index]) {
                // update existing dataset
                chart.data.datasets[index].data = data;
            } else {
                // add new dataset
                chart.data.datasets.push({
                    label: player.username,
                    data,
                    hitRadius: 10,
                    borderColor: `hsl(${index * 60}, 100%, 50%)`,
                    backgroundColor: `hsla(${index * 60}, 100%, 50%, 0.5)`
                });
            }
        });

        // remove any extra datasets
        if (chart.data.datasets.length > value.length) {
            chart.data.datasets.splice(value.length);
        }

        chart.update();
    }

    onMount(() => {
        if (!chart) {
        chart = new Chart(canvas.getContext('2d'), {
            type: 'line',
            data: {
                labels: [],
                datasets: [],
            },
            options: {
                events: ['mousemove'],
                interaction: {
                    intersect: false,
                    mode: 'nearest',
                    axis: 'xy',
                },
                maintainAspectRatio: true,
                datasets: {
                    line: {
                        pointRadius: 0,
                        normalized: true,
                    }
                },
                plugins: {
                    legend: {
                        position: "top",
                        
                    },
                    tooltip: {
                        enabled: true,
                        mode: 'nearest',
                        position: 'nearest',
                        padding: 10,
                        backgroundColor: 'rgba(212, 30, 109, 1)',
                        borderColor: 'rgba(92, 10, 46, 1)',
                        borderWidth: 2,
                        caretSize: 10,
                        cornerRadius: 5,
                        displayColors: false,
                        callbacks: {
                            labelTextColor: function() {
                                return 'white';
                            }
                        },
                    },
                },
                scales: {
                    x: {
                        type: 'time',
                        time: {
                            unit: "year",
                        }
                    }
                }
            }
        });

        unsubscribe = players.subscribe(updateChartData);
    }});

    onDestroy(() => {
        unsubscribe();
    });
</script>

<canvas bind:this={canvas}></canvas>
<div>
    {$selectedStat}
</div>

<style>
    canvas {
        height: 75%;
        width: 80%;
    }
</style>