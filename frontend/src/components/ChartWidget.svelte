<script lang="ts">
  import { onMount } from 'svelte';
  import { Chart as ChartJS, registerables } from 'chart.js';
  
  export let type: 'bar' | 'line' | 'pie' | 'doughnut' | 'radar' = 'bar';
  export let data: any = {};
  export let options: any = {};
  export let title: string = '';
  export let height: string = '300px';
  export let width: string = '100%';
  
  let canvas: HTMLCanvasElement;
  let chart: ChartJS;
  
  onMount(() => {
    ChartJS.register(...registerables);
    
    const ctx = canvas.getContext('2d');
    chart = new ChartJS(ctx!, {
      type,
      data,
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            position: 'bottom',
          },
          title: {
            display: !!title,
            text: title,
          },
          ...options.plugins,
        },
        ...options,
      },
    });
    
    return () => {
      chart?.destroy();
    };
  });

  export function updateChart(newData: any) {
    if (chart) {
      chart.data = newData;
      chart.update();
    }
  }
</script>

<div class="chart-widget" style="height: {height}; width: {width}">
  <canvas bind:this={canvas}></canvas>
</div>

<style>
  .chart-widget {
    background: white;
    border-radius: 8px;
    padding: 16px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    position: relative;
  }

  canvas {
    max-height: 100%;
  }
</style>
