// src/components/ChartComponent.tsx
import { createSignal, onCleanup, onMount } from "solid-js";
import { createQuery } from "@tanstack/solid-query";
import Chart from "chart.js/auto";

// Define the type for the data fetched from the API
interface ChartData {
  labels: string[];
  values: number[];
}

const fetchData = async (): Promise<ChartData> => {
  const response = await fetch("https://api.example.com/data");
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  return response.json();
};

const ChartComponent = () => {
  let canvas: HTMLCanvasElement | undefined;
  const [chart, setChart] = createSignal<Chart | null>(null);

  const query = createQuery<ChartData, Error>({
    queryKey: ["dataKey"],
    queryFn: fetchData,
  });

  onMount(() => {
    if (query.data) {
      const newChart = new Chart(canvas!, {
        type: "line", // Or any other chart type you want
        data: {
          labels: query.data?.labels,
          datasets: [
            {
              label: "Dataset",
              data: query.data.values,
              backgroundColor: "rgba(75, 192, 192, 0.2)",
              borderColor: "rgba(75, 192, 192, 1)",
              borderWidth: 1,
            },
          ],
        },
        options: {
          responsive: true,
          maintainAspectRatio: false,
        },
      });
      setChart(newChart);
    }
  });

  onCleanup(() => {
    chart()?.destroy();
  });

  return (
    <div>
      {query.isLoading && <p>Loading...</p>}
      {query.error && <p>Error: {query.error.message}</p>}
      {!query.isLoading && !query.error && <canvas ref={canvas}></canvas>}
    </div>
  );
};

export default ChartComponent;
