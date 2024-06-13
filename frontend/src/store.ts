// src/store.ts
import { defineStore } from 'pinia'

export const useStore = defineStore('main', {
  state: () => ({
    chartData: [] as Array<{ timestamp: string; value: number }>
  }),
  actions: {
    setChartData(data: Array<{ timestamp: string; value: number }>) {
      this.chartData = data
    }
  }
})
