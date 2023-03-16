const go = new Go();

const loadWasm = async (path) => {
    const result = await WebAssembly.instantiateStreaming(fetch(path), go.importObject)
    go.run(result.instance)
    return result
}

    ;
(async () => {
    await loadWasm("/lines.wasm")
    const plot = createLinePlot(document.getElementById('canvas_lines'))
    console.log({ plot })
    const rand = n => Math.random() * (n || 1)
    const data = Array(600).fill(0).map(() => rand())
    console.log({ data })
    plot.changeBgColor(1, 1, 1)
    plot.changeLineColor(0.3, 0.02, 0.1)
    plot.setData(...data)
    setInterval(() => {
        // plot.changeBgColor(Math.random(), Math.random(), Math.random())
        plot.flowData(rand())
        plot.render()
    }, 1/144)
})();

// ;
// (async () => {
//     await loadWasm("/hellowasm.wasm")
//     hellowasm("oi")
//     console.log(hellovalue())
// })();

// ;
// (async () => {
//     await loadWasm("/triangle.wasm")
//     const triangleElement = document.getElementById('canvas_triangle')
//     const triangle = createTriangle(triangleElement)
//     console.log({triangle})
//     setInterval(() => triangle.changeBgColor(Math.random(), Math.random(), Math.random()), 300)
// })();
