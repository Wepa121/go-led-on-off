document.addEventListener("DOMContentLoaded", () => {
  const onDiv = document.querySelector("#on")
  const offDiv = document.querySelector("#off")

  onDiv.addEventListener("click", async() => {
    const res = await fetch("http://192.168.58.100:3333/api/ac")
    const data = await res.json()
    console.log(data)
  })

  offDiv.addEventListener("click", async() => {
    const res = await fetch("http://192.168.58.100:3333/api/yap")
    const data = await res.json()
    console.log(data)
  })

})