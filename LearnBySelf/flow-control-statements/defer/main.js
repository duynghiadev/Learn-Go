async function fetchData() {
  console.log("Start of async function");

  let data = await new Promise((resolve, reject) => {
    setTimeout(() => {
      return resolve("Data fetched");
    }, 1000);
  });

  console.log(data);
  console.log("End of async function");
}

fetchData();
