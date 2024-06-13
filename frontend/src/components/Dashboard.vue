<template>
  <div>
    <canvas ref="chartCanvas"></canvas>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useQuery } from '@tanstack/vue-query'
import { useStore } from '../store'
import Chart from 'chart.js/auto'

const chartCanvas = ref<HTMLCanvasElement | null>(null)
let chartInstance: Chart | null = null

const store = useStore()

// Simulación de la consulta a la API
const fetchData = async () => {
  const response = await fetch('/api/data')
  return await response.json()
}

// Configuración de TanStack Query
const query = useQuery({
  queryKey: ['chartData'],
  queryFn: fetchData,
  refetchInterval: 60000 // Refetch cada 1 minuto
})

// Observa los datos y actualiza el gráfico cuando cambian
watch(query.data, (newData) => {
  if (newData) {
    store.setChartData(newData)
    updateChart()
  }
})

const updateChart = () => {
  if (chartInstance) {
    chartInstance.destroy()
  }

  if (chartCanvas.value) {
    chartInstance = new Chart(chartCanvas.value, {
      type: 'line',
      data: {
        labels: store.chartData.map(d => d.timestamp),
        datasets: [{
          label: 'Tiempo de Envío de Datos',
          data: store.chartData.map(d => d.value),
          backgroundColor: 'rgba(75, 192, 192, 0.2)',
          borderColor: 'rgba(75, 192, 192, 1)',
          borderWidth: 1
        }]
      },
      options: {
        scales: {
          y: {
            beginAtZero: true
          },
          x: {
            type: 'time',
            time: {
              unit: 'minute',
              tooltipFormat: 'll HH:mm'
            }
          }
        },
        responsive: true,
        maintainAspectRatio: false
      }
    })
  }
}

onMounted(() => {
  updateChart()
})
</script>

<style scoped>
canvas {
  width: 100%;
  height: 400px;
}
</style>