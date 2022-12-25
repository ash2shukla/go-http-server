window.onload = () => {
    const p = document.createElement("p")
    p.innerHTML = "created from js loaded from server"

    document.body.appendChild(p)
}