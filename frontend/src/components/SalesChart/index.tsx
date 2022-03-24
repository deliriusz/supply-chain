import {
   Chart as ChartJS,
   CategoryScale,
   LinearScale,
   PointElement,
   LineElement,
   Title,
   Tooltip,
   Legend,
} from 'chart.js'
import { Chart } from 'react-chartjs-2'

import React, { useState } from "react"

ChartJS.register(
   CategoryScale,
   LinearScale,
   PointElement,
   LineElement,
   Title,
   Tooltip,
   Legend
)

const labels = [
   'January',
   'February',
   'March',
   'April',
   'May',
   'June',
];

const data = {
   labels: labels,
   datasets: [{
      label: 'My First dataset',
      backgroundColor: 'rgb(255, 99, 132)',
      borderColor: 'rgb(255, 99, 132)',
      lineTension: 0.3,
      data: [0, 10, 5, 2, 20, 30, 45],
   }]
};

const SalesChart = () => {

   return (
      <div>
         <Chart type='line' data={data}  />
      </div>
   )
}

export default SalesChart