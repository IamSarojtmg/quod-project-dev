async function fetchData() {
  try {
    const response = await fetch("/api/time-series", { method: "GET" });
    const data = await response.json();

    const labels = data.map((point) => point.Time);
    const values = data.map((point) => point.Value);

    const startTemperature = values[0];
    const endTemperature = values[values.length - 1];
    const deltaTemperature = endTemperature - startTemperature;
    const currentTemperature = endTemperature;

    document.getElementById("startTemp").textContent = startTemperature + "°F";
    document.getElementById("endTemp").textContent = endTemperature + "°F";
    document.getElementById("deltaTemp").textContent =
      "Delta: " + deltaTemperature;
    document.getElementById("currentTemp").textContent =
      currentTemperature + "°F";

    deltaTemperature >= 0
      ? (deltaTemp.style.color = "green")
      : (deltaTemp.style.color = "#739BD0");

    const getBackgroundColor = (temp) => {
      if (temp < 50) return "#008000";
      if (temp < 60) return "#FFFF00";
      if (temp < 71) return "#FFA500";
      return "#B36262";
    };

    startTemp.style.backgroundColor = getBackgroundColor(startTemperature);
    endTemp.style.backgroundColor = getBackgroundColor(endTemperature);
    currentTemp.style.backgroundColor = getBackgroundColor(currentTemperature);

    const setColorOnTemp = (temp) => {
      if (temp < 50) return "#32CD32";
      if (temp < 60) return "#FFFF66";
      if (temp < 71) return "#FFB347";
      return "#D98888";
    };

    const tempDetail = document.querySelectorAll(".tempDetail");
    tempDetail.forEach((temp, index) => {
      if (index === 0) {
        temp.style.backgroundColor = setColorOnTemp(startTemperature);
      } else if (index === 1) {
        temp.style.backgroundColor = setColorOnTemp(endTemperature);
      } else if (index === 2) {
        temp.style.backgroundColor = setColorOnTemp(currentTemperature);
      }
    });
    console.log(chart);

    chart.updateOptions({
      series: [{ name: "Temperature (°F)", data: values }],
      xaxis: { categories: labels },
    });
  } catch (error) {
    console.error("Error fetching or rendering data:", error);
  }
}
