let chart;
document.addEventListener("DOMContentLoaded", () => {

const options = {
chart: {
  height: "65%",
  type: "area",
  fontFamily: "Inter, sans-serif",
  toolbar: {
    show: false,
  },
},
tooltip: {
  enabled: true,
},
fill: {
  type: "gradient",
  gradient: {
    opacityFrom: 0.55,
    opacityTo: 0,
  },
},
dataLabels: {
  enabled: false,
},
stroke: {
  curve: "smooth",
  width: 6,
},
grid: {
  show: false,
},
series: [
  {
    name: "Temperature (°F)",
    data: [],
    color: "#1A56DB",
  },
],
xaxis: {
  categories: [],
  labels: {
    style: {
      colors: "#6B7280",
    },
    offsetX: 1,
  },
},
yaxis: {
  show: false,
},
};

chart = new ApexCharts(document.getElementById("temperatureChart"), options);
chart.render();

fetchData();
});