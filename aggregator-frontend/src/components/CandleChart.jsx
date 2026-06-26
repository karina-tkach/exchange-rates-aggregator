import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    TimeScale,
    Tooltip,
    Legend
} from "chart.js"

import {
    CandlestickController,
    CandlestickElement
} from "chartjs-chart-financial"

import "chartjs-adapter-date-fns"
import {Chart} from "react-chartjs-2";

ChartJS.register(
    CategoryScale,
    LinearScale,
    TimeScale,
    Tooltip,
    Legend,
    CandlestickController,
    CandlestickElement
)

export default function CandleChart({ data }) {
    const chartData = {
        datasets: [
            {
                label: "OHLC",
                data: data.map(c => ({
                    x: new Date(c.time).getTime(),
                    o: Number(c.open),
                    h: Number(c.high),
                    l: Number(c.low),
                    c: Number(c.close),
                })),
                borderColor: "#14a3c7",
                color: {
                    up: "#14a3c7",
                    down: "#ef4444",
                    unchanged: "#DAFCE6",
                },
            },
        ],
    }

    return (
        <div className="h-full w-full rounded-2xl bg-[#020617] border border-[#22267C] p-4">
            <Chart
                type="candlestick"
                data={chartData}
                options={{
                    responsive: true,
                    maintainAspectRatio: false,
                    parsing: false,
                    plugins: {
                        legend: { display: false },
                        tooltip: {
                            backgroundColor: "#00002B",
                            titleColor: "#DAFCE6",
                            bodyColor: "#DAFCE6",
                            borderColor: "#22267C",
                            borderWidth: 1,
                        },
                    },
                    scales: {
                        x: {
                            type: "time",
                            time: {
                                unit: "minute",
                            },
                            ticks: { color: "#DAFCE6" },
                            grid: { color: "#22267C" },
                        },
                        y: {
                            ticks: { color: "#DAFCE6" },
                            grid: { color: "#22267C" },
                        },
                    },
                }}
            />
        </div>
    )
}